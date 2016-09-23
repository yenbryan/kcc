package main

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"net/http"
)

func Start() {
	m := martini.Classic()
	martini.Env = martini.Prod
	logger := log.New(ioutil.Discard, "", 0)
	m.Map(logger)

	m.Get("/", mainHandler)

	m.Use(martini.Static("www"))
	m.RunOnAddr(":8081")
}

var	c = make(chan bool)

func main() {
	go Start()
	<- c
}

func mainHandler(res http.ResponseWriter, req *http.Request) (int, string) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	data, _ := ioutil.ReadFile("www/index.html")
	return 200, string(data)
}
