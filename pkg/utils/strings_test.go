package utils_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/yashptel/test-api/pkg/utils"
)

func Test_GetIntFromURL(t *testing.T) {
	num := utils.GetIntFromURL(&http.Request{URL: &url.URL{RawQuery: "id=1"}}, "id", 10)
	if num != 1 {
		t.Error("Expected 1, got ", num)
	}

	num = utils.GetIntFromURL(&http.Request{URL: &url.URL{RawQuery: ""}}, "id", 10)
	if num != 10 {
		t.Error("Expected 10, got ", num)
	}

	num = utils.GetIntFromURL(&http.Request{URL: &url.URL{RawQuery: "id=abc"}}, "id", 10)
	if num != 10 {
		t.Error("Expected 10, got ", num)
	}

	num = utils.GetIntFromURL(&http.Request{URL: &url.URL{RawQuery: "id=-1.1"}}, "id", 10)
	if num != 10 {
		t.Error("Expected 10, got ", num)
	}
}
