package main

import (
	"log"

	"github.com/grafana-plugin-mongodb/api"
)

func main() {
	conf := api.Config{Port: 8080, Host: "localhost"}
	errs := make(chan error, 2)
	api.StartHTTPServer(conf, errs)
	log.Println("start")
	for {
		err := <-errs
		log.Println(err)
	}
}