package main

import (
	"log"
	"time"

	"github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile(time.Now().String(), qrcode.Medium, 256, "qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
}
