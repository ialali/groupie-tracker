package handlers

import (
	"encoding/json"
	global "groupie-tracker/globals"
)

// var Client *http.Client

func GetJson(url string, target interface{}) error {
	response, err := global.Client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}
