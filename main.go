package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

const (
	BLOBS				= "blobs"
	MBOARD      = "mboard"
	REPOSITORY 	= "repository"
	UPDATE  		= "update"
	UPDATES   	= "updates"
	VERSION 		= "version"
)

var (
  port     = flag.String("port", "8900", "Host address")
  conf     = flag.String("conf", "config.json", "Configuration file")
)

var config map[string]string

func initEndpoints() {

	config[BLOBS]   = fmt.Sprintf("http://%s", config[REPOSITORY])
	config[UPDATE] 	= fmt.Sprintf(
		"http://%s/api/upgrades/mboard", config[REPOSITORY])
  config[VERSION] = fmt.Sprintf("http://%s/api/version", config[MBOARD])
	
} // initEndpoints

func initConfig() {

	buf, err := ioutil.ReadFile(*conf)

	if err != nil {
		log.Fatal(err)
	} else {

		err := json.Unmarshal(buf, &config)
		
		if err != nil {
			log.Fatal(err)
		}

		initEndpoints()

	}

} // initConfig

func initRouter() *mux.Router {

	router := mux.NewRouter()
	
	router.HandleFunc("/api/update",
		updateAPIHandler)

	return router

} // initRouter

func main() {

	flag.Parse()

	initConfig()

	log.Printf("Listening on port %s...", *port)

	addr := fmt.Sprintf(":%s", *port)

	http.ListenAndServe(addr, initRouter())

}
