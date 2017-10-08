package utils

import (
	"bytes"
	"encoding/binary"
)

func Obj2ByteAarray(obj interface{}) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, obj)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func ByteAarray2Obj(b []byte, obj interface{}) {

	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &obj)
	if err != nil {
		panic(err)
	}

}
