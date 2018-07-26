package main

import (
	"log"

	"github.com/moustik/grafana-simplejson-mongo/api"
)

func main() {
	conf := api.Config{
		Port:      8080,
		MongoHost: "192.168.1.88",
	}
	errs := make(chan error, 2)
	api.StartHTTPServer(conf, errs)
	log.Println("start")
	for {
		err := <-errs
		log.Println(err)
	}
}
