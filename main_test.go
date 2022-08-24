package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	testClient = http.Client{}
	testURL = "http://" + getEnvValue("HOST") + ":" + getEnvValue("PORT")
}
func TestServerPing(t *testing.T) {
	res, err := testClient.Get(testURL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	b := string(body)
	if !strings.Contains(b, "QR Code") {
		t.Fatal()
	}
}
