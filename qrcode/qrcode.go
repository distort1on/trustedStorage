package qrcode

import (
	"github.com/skip2/go-qrcode"
	"strconv"
)

func WriteQrToFile(index uint) {
	//test qr code to get block by id

	err := qrcode.WriteFile("localhost:8080/blockchain/"+strconv.Itoa(int(index)), qrcode.Medium, 256, "qr.png")
	if err != nil {
		panic(err)
	}
}
