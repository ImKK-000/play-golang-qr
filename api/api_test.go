package api_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	. "play-golang/api"
	"testing"
)

func Test_APIHandler_Should_Be_QrCode(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody, _ := ioutil.ReadFile("../qrcode.png")
	request := httptest.NewRequest("GET", "/api/generateqrcode", nil)
	response := httptest.NewRecorder()

	APIHandler(response, request)
	actualStatusCode := response.Code
	actualBody, _ := ioutil.ReadAll(response.Body)

	if expectedStatusCode != actualStatusCode {
		t.Errorf("expected %d but it got %d", expectedStatusCode, actualStatusCode)
	}
	if len(expectedBody) != len(actualBody) {
		t.Errorf("expected %d but it got %d", len(expectedBody), len(actualBody))
	}
}
