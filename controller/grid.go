package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"usebottles.com/steamgrid-proxy/config"
	"usebottles.com/steamgrid-proxy/proxy"
)

func Search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchTerm := vars["gameName"]
	searchType := vars["type"]

	if searchTerm == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing gameName."))
		return
	}

	if searchType == "" {
		searchType = "grids"
	}

	if (searchType != "grids") && (searchType != "heroes") && (searchType != "logos") && (searchType != "icons") {
		w.WriteHeader(400)
		w.Write([]byte("Invalid search type"))
		return
	}

	link := getFromCache(searchTerm, searchType)
	if link != "" {
		jsonRes, _ := json.Marshal(link)
		w.WriteHeader(200)
		w.Write(jsonRes)
		return
	}

	res, err := proxy.Search(searchTerm, searchType)

	if err != nil {
		if res == "404" {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		w.Write([]byte(res))
		return
	}

	jsonRes, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Write(jsonRes)
}

func getFromCache(g string, s string) string {
	data, err := os.ReadFile(filepath.Join(config.ProcessPath, "cache", s, g+".txt"))
	if err != nil {
		return ""
	}

	link := string(data)

	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	r, err := client.Do(req)
	if err != nil || r.StatusCode != 200 {
		return ""
	}

	return link
}
