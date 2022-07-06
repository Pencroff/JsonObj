package parsing

import (
	"bufio"
	"github.com/Pencroff/JsonStruct"
)

func ParsePrimitiveValue(rd *bufio.Reader, v JsonStruct.JStructOps) (e error) {
	bt, e := rd.ReadByte()
	switch bt {
	case 'n':
		v.SetNull()
		return
	case 'f':
		v.SetBool(false)
		return
	case 't':
		v.SetBool(true)
		return
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		v.SetInt(int64(bt - '0'))
	case '"':
		v.SetString(string(bt))
	}
	return
}
