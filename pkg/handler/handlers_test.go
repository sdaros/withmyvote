package withmyvote

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(homePageHandler))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		t.Errorf("Erfror not expected: %s\n", err)
	}

	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}
	indexFile, _ := ioutil.ReadFile("index.html")
	if string(result[:]) != string(indexFile[:]) {
		t.Errorf("Expected result: %s\n got: %s\n", "Hello, client", result)
	}
}
