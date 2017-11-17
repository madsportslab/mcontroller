package main;

import (
	"log"
	"net/http"

)

func updateAPIHandler(w http.ResponseWriter, r *http.Request) {
	
		switch r.Method {
		case http.MethodPost:
			log.Println("update requested")

			// parse hash id
		case http.MethodGet:
		case http.MethodDelete:
		case http.MethodPut:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	
	} // updateAPIHandler
	