package serialization

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Serialize(data interface{}) []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(data)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return buffer.Bytes()
}

func DeSerialize(object interface{}, data []byte) bool {
	buffer := bytes.NewBuffer(data)
	d := gob.NewDecoder(buffer)
	if err := d.Decode(object); err != nil {
		panic(err)
	}
	return true
}
