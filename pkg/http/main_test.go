package http

import (
	"net/http"
	"testing"
	"time"
)

func Test_RunHttpServer(t *testing.T) {
	go RunHttpServer()
	time.Sleep(time.Second)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/healthz", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Expected status code to be 200")
	}
}
