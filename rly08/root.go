package rly08

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const GetVersion    = 0x5A 	//	Get software version - returns 2 bytes, the first being the Module ID which is 8, followed by the software version
const GetRelayState = 0x5B  //	Get relay states - sends a single byte back to the controller, bit high meaning the corresponding relay is powered

const SetRelayState = 0x5C	//	Set relay states - the next single byte will set all relays states, All on = 255 (11111111) All off = 0
const AllOn = 0x64			//	All relays on
const Relais1On = 0x65		//	Turn relay 1 on
const Relais2On = 0x66		//	Turn relay 2 on
const Relais3On = 0x67		//	Turn relay 3 on
const Relais4On = 0x68		//	Turn relay 4 on
const Relais5On = 0x69		//	Turn relay 5 on
const Relais6On = 0x6A		//	Turn relay 6 on
const Relais7On = 0x6B		//	Turn relay 7 on
const Relais8On = 0x6C		//	Turn relay 8 on

const AllOff	= 0x6E		//	All relays off
const Relais1Off = 0x6F		//	Turn relay 1 off
const Relais2Off = 0x70		//	Turn relay 2 off
const Relais3Off = 0x71		//	Turn relay 3 off
const Relais4Off = 0x72		//	Turn relay 4 off
const Relais5Off = 0x73		//	Turn relay 5 off
const Relais6Off = 0x74		//	Turn relay 6 off
const Relais7Off = 0x75		//	Turn relay 7 off
const Relais8Off = 0x76		//	Turn relay 8 off

var mapping = make(map[string][]byte)

var lowerCase  = cases.Lower(language.AmericanEnglish)

func init()  {
	mapping["getversion"] = []byte{GetVersion}
	mapping["getrelaystate"] = []byte{GetRelayState}

	mapping["setrelaystate"] = []byte{SetRelayState}
	mapping["allon"] = []byte{AllOn}
	mapping["alloff"] = []byte{AllOff}

	mapping["relais1on"] = []byte{Relais1On}
	mapping["relais2on"] = []byte{Relais2On}
	mapping["relais3on"] = []byte{Relais3On}
	mapping["relais4on"] = []byte{Relais4On}
	mapping["relais5on"] = []byte{Relais5On}
	mapping["relais6on"] = []byte{Relais6On}
	mapping["relais7on"] = []byte{Relais7On}
	mapping["relais8on"] = []byte{Relais8On}

	mapping["relais1off"] = []byte{Relais1Off}
	mapping["relais2off"] = []byte{Relais2Off}
	mapping["relais3off"] = []byte{Relais3Off}
	mapping["relais4off"] = []byte{Relais4Off}
	mapping["relais5off"] = []byte{Relais5Off}
	mapping["relais6off"] = []byte{Relais6Off}
	mapping["relais7off"] = []byte{Relais7Off}
	mapping["relais8off"] = []byte{Relais8Off}

}

func GetCommand(cmd string) []byte {
	return mapping[lowerCase.String(cmd)]
}

func IsGetter(cmd string) bool {
	switch GetCommand(cmd)[0] {
		case GetRelayState:
		case GetVersion:
			return true
	}

	return false
}
