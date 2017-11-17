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

var (
  port     = flag.String("port", "8090", "Host address")
  conf     = flag.String("conf", "config.json", "Configuration file")
)

var config map[string]string

func initRepository() {

	buf, err := ioutil.ReadFile(*conf)

	if err != nil {
		log.Fatal(err)
	} else {

		err := json.Unmarshal(buf, &config)
		
		if err != nil {
			log.Fatal(err)
		}

	}

} // initRepository

func initRouter() *mux.Router {

	router := mux.NewRouter()
	
	router.HandleFunc("/api/update",
		updateAPIHandler)

	return router

} // initRouter

func main() {

	flag.Parse()

	initRepository()

	log.Println(config)

	log.Printf("Listening on port %s...", *port)

	addr := fmt.Sprintf(":%s", *port)

	http.ListenAndServe(addr, initRouter())

}
