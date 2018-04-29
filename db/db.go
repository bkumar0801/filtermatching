package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/*
City ...
*/
type City struct {
	Name      string  `json:"name"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

/*
Profile ...
*/
type Profile struct {
	Display   string  `json:"display_name"`
	Age       int32   `json:"age"`
	JobTitle  string  `json:"job_title"`
	Height    int32   `json:"height_in_cm"`
	Location  City    `json:"city"`
	Photo     string  `json:"main_photo"`
	Score     float32 `json:"compatibility_score"`
	Contacts  int32   `json:"contacts_exchanged"`
	Favourite bool    `json:"favourite"`
	Religion  string  `json:"religion"`
}

/*
Profiles ...
*/
type Profiles struct {
	Matches []Profile `json:"matches"`
}

/*
LoadProfiles ...
*/
func LoadProfiles(filename string) (Profiles, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}

	profiles := Profiles{}
	err = json.Unmarshal(raw, &profiles)
	if err != nil {
		log.Println(err.Error())
	}
	return profiles, err
}
