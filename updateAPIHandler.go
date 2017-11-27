package main;

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/madsportslab/glbs"
)

func getCurrentFirmware() *string {

	current, err := http.Get(config[VERSION])
	
	if err != nil {
		log.Println(err)
		return nil
	} else {

		data, _ := ioutil.ReadAll(current.Body)
		
		j := map[string]string{}

		json.Unmarshal(data, &j)

		if j[VERSION] != "" {
			str := j[VERSION]
			return &str
		} else {
			return nil
		}

	}

} // getCurrentFirmware

func getUpdateFirmware(version string) map[string]string {

	url := fmt.Sprintf("%s/%s", config[UPDATE], version)
	
	update, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return nil
	} else {

		js, _ := ioutil.ReadAll(update.Body)

		r := map[string]string{}

		json.Unmarshal(js, &r)

		return r
		
	}

} // getUpdateFirmware

func downloadUpdate() *string {

	current := getCurrentFirmware()

	if current != nil {

		update := getUpdateFirmware(*current)

		glbs.SetNamespace(BLOBS)

		url := fmt.Sprintf("%s/%s", config[BLOBS],
			*(glbs.GetPath(update["hashId"])))

		res, err := http.Get(url)

		if err != nil {
			log.Println(err)
			return nil
		} else {

			if res.StatusCode == 200 {

				glbs.SetNamespace(UPDATES)

				k := glbs.Put(res.Body)
				
				if k != nil {
					return glbs.GetPath(*k)
				} else {
					return nil
				}

			} else {
				return nil
			}

		}
		
	} else {
		return nil
	}

			
} // downloadUpdate

func installUpdate(file string) bool {
	
	glbs.SetNamespace(UPDATES)

	deb := fmt.Sprintf("sudo dpkg -i %s", file)
	
	out, err := exec.Command(deb).Output()

	if err != nil {
		log.Printf("installUpdate(): %s", err)
		return false
	} else {
		log.Println(string(out))
		return true			
	}
	
} // installUpdate

func updateAPIHandler(w http.ResponseWriter, r *http.Request) {
	
		switch r.Method {
		case http.MethodPost:

			id := downloadUpdate()

			if id != nil {
				
				if !installUpdate(*id) {
					w.WriteHeader(http.StatusInternalServerError)
				}

			} else {
				w.WriteHeader(http.StatusNotFound)
			}

		case http.MethodGet:
		case http.MethodDelete:
		case http.MethodPut:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	
	} // updateAPIHandler
	