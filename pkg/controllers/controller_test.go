package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yashptel/test-api/pkg/controllers"
)

var srv *httptest.Server

func RunTestServer() {
	router := controllers.NewRouter()
	srv = httptest.NewServer(router)
}

func TestMain(m *testing.M) {
	RunTestServer()
	defer srv.Close()
	code := m.Run()
	os.Exit(code)
}

func Test_NewRouter(t *testing.T) {
	router := controllers.NewRouter()
	if router == nil {
		t.Error("Expected router to be not nil")
	}

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/healthz", nil)
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

	req, err = http.NewRequest(http.MethodGet, srv.URL+"/api/ofdjk", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Error("Expected status code to be 404")
	}
}
