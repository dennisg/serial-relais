package main


import (
	"flag"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"io"
	"github.com/dennisg/serial-relais/rly08"
	"time"
)

var device string
var port io.ReadWriteCloser

func init() {
	flag.StringVar(&device, "port", "/dev/ttyACM0", "the device to use")
	flag.Parse()
}

func main() {

	options := serial.OpenOptions{
		PortName: device,
		BaudRate: 19200,
		DataBits: 8,
		StopBits: 1,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}


	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/cmd/{command}", CommandsHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	port.Write([]byte{rly08.GetVersion})
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CommandsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Command: %v\n", vars["command"])

	cmd := rly08.GetCommand(vars["command"])
	_, err := port.Write(cmd)

	if (err != nil) {
		w.WriteHeader(http.StatusPreconditionFailed)
	} else {
		w.WriteHeader(http.StatusOK)

		if rly08.IsGetter(vars["command"]) {
			var data [2]byte
			n, _ := port.Read(data[:])
			if n > 0 {
				fmt.Printf("bytes read: %d Bytes: [% x]\n", n, data)
			}
			fmt.Fprintf(w, "Read back: [% x]\n", data)
		}

	}

}