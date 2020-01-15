package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	ConstHeader         = "www.xx.com"
	ConstHeaderLength   = 10
	ConstSaveDataLength = 4
)

func Packet(messgae []byte) []byte {

	return append(append([]byte(ConstHeader), IntoToBytes(len(messgae))...), messgae...)
}

func UnPack(buffer []byte, readerChannel chan []byte) []byte {
	fmt.Println("unpack...")
	fmt.Println("unpack data...", string(buffer))
	length := len(buffer)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			if length < i+ConstSaveDataLength+ConstHeaderLength+messageLength {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			fmt.Println(string(data))
			readerChannel <- data
			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	fmt.Println("i:", i)
	return buffer[i:]
}

func IntoToBytes(n int) []byte {

	x := int32(n)
	bytesBUffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBUffer, binary.BigEndian, x)
	return bytesBUffer.Bytes()
}

func BytesToInt(b []byte) int {

	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}
