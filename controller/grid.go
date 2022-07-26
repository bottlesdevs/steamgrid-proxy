package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"usebottles.com/steamgrid-proxy/proxy"
)

func Search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchTerm := vars["gameName"]
	
	if searchTerm == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing gameName."))
		return
	}

	res, err := proxy.Search(searchTerm)

	if err != nil {
		if res == "404" {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		w.Write([]byte(res))
		return
	}

	json, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Write(json)
}
