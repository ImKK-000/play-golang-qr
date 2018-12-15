package api

import (
	"net/http"

	qrcodeEncode "github.com/skip2/go-qrcode"
)

func APIHandler(responseWriter http.ResponseWriter, request *http.Request) {
	qrCode, _ := qrcodeEncode.Encode("https://google.com", qrcodeEncode.Medium, 256)
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(qrCode)
}
