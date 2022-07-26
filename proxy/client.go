package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"usebottles.com/steamgrid-proxy/config"
)


type ProxySearchResponse struct {
	Success bool `json:"success"`
	Data []GameData `json:"data"`
} 

type GameData struct {
	Types []string `json:"types"`
	Id int `json:"id"`
	Name string `json:"name"`
	Verified bool `json:"verified"`
	ReleaseDate string `json:"release_date,omitempty"`
}

type ProxyGridResponse struct {
	Success bool `json:"success"`
	Data []GridData `json:"data"`
} 

type GridData struct {
	Url string `json:"url"`
}

const BASE_URL = "https://www.steamgriddb.com/api/v2"

func callAPI(e string, t string) (r *http.Response, err error) {
	cnf := *config.Cnf
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL + e + t, nil)

	req.Header.Set("Authorization", "Bearer " + cnf.ApiKey)
	req.Header.Set("User-Agent", "curl/7.79.1")
	r, err = client.Do(req)
	return r, err
}

func Search(t string) (m string, err error) {
	res, err := callAPI("/search/autocomplete/", t)

	if err != nil {
		return "", err
	}

	var searchResponse ProxySearchResponse
	buf, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(buf, &searchResponse)

	if len(searchResponse.Data) == 0 {
		return "404", errors.New("404 Game Not Found")
	}

	msg := searchResponse.Data[0].Name
	res, err = callAPI("/grids/game/", fmt.Sprint(searchResponse.Data[0].Id))

	if err != nil {
		return "", err
	}

	var gridResponse ProxyGridResponse
	buf, _ = ioutil.ReadAll(res.Body)
	json.Unmarshal(buf, &gridResponse)

	if len(gridResponse.Data) == 0 {
		return "404", errors.New("404 Grid Not Found")
	}

	msg = gridResponse.Data[0].Url

	return msg, nil
}