package serialization

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
)

func Serialize(data interface{}) []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(data)
	if err != nil {
		log.Println("encode error:", err)
	}
	return buffer.Bytes()
}

func DeSerialize(object interface{}, data []byte) error {
	buffer := bytes.NewBuffer(data)
	d := gob.NewDecoder(buffer)
	if err := d.Decode(object); err != nil {
		return errors.New("Cant deserialize")
	}
	return nil
}
