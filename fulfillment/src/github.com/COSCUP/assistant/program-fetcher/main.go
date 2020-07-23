package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// var PROGRAM_URL = "https://api2019.coscup.org/programs.json" // 2019
var PROGRAM_URL = "https://coscup.org/2020/json/session.json"

var cache *ProgramsResponedPayload

func RefreshCache() error {
	c, err := FetchProgramsResponedPayload()
	if err == nil {
		cache = c
	}
	return err
}

func GetPrograms() (*ProgramsResponedPayload, error) {
	return cache, nil
}

func FetchProgramsResponedPayload() (*ProgramsResponedPayload, error) {

	resp, err := http.Get(PROGRAM_URL)
	if err != nil {
		// handle err
		return nil, err
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle err
		return nil, err
	}
	ret := &ProgramsResponedPayload{}
	json.Unmarshal(payload, ret)
	return ret, nil

}
