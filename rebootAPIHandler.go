package main;

import (
	"log"
	"net/http"
	"os/exec"

)


func rebootAPIHandler(w http.ResponseWriter, r *http.Request) {
	
		switch r.Method {
		case http.MethodPost:

			err := exec.Command("reboot")

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

		case http.MethodGet:
		case http.MethodDelete:
		case http.MethodPut:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	
	} // rebootAPIHandler
	
