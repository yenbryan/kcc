package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	go Start()
	os.Exit(m.Run())
}

func TestConnect(t *testing.T) {
	resp, err := http.Get("http://localhost:8081")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(body), "<!DOCTYPE html>") {
		t.Errorf("Body not html")
	}
}
