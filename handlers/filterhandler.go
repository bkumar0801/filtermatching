package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/filtermatching/stream"

	"github.com/filtermatching/db"
	f "github.com/filtermatching/filter"
)

/*
HandleFilter ...
*/
func HandleFilter(filename string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Content-Type", "application/json")
		filter, err := BuildFilterFromQuery(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		profiles, err := db.LoadProfiles(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var matchedProfiles db.Profiles
		profileOutChan := make(chan db.Profile)
		for _, profile := range profiles.Matches {
			matcher := stream.Apply(profile, *filter)
			subscription := stream.Subscribe(profileOutChan, matcher)
			matched := <-subscription.Updates()
			if !reflect.DeepEqual(matched, db.Profile{}) {
				matchedProfiles.Matches = append(matchedProfiles.Matches, matched)
			}
		}
		close(profileOutChan)
		json.NewEncoder(w).Encode(matchedProfiles)
	}
}

/*
BuildFilterFromQuery ...
*/
func BuildFilterFromQuery(r *http.Request) (*f.Filter, error) {
	if !checkQueryParams(r) {
		return nil, errors.New("missing query params")
	}
	photo, _ := strconv.ParseBool(r.FormValue("photo"))
	inContacts, _ := strconv.ParseBool(r.FormValue("in_contacts"))
	favouraite, _ := strconv.ParseBool(r.FormValue("favouraite"))
	score, _ := strconv.ParseFloat(r.FormValue("compatibility_score"), 32)
	if score <= 0 {
		score = 0.01
	}
	minAge, _ := strconv.ParseInt(r.FormValue("minAge"), 10, 32)
	maxAge, _ := strconv.ParseInt(r.FormValue("maxAge"), 10, 32)
	maxHeight, _ := strconv.ParseInt(r.FormValue("height"), 10, 32)
	upperDistance, _ := strconv.ParseInt(r.FormValue("distance"), 10, 32)
	return f.NewFilter(photo, inContacts, favouraite, float32(score), int32(minAge), int32(maxAge), int32(maxHeight), int32(upperDistance)), nil
}

func checkQueryParams(r *http.Request) bool {
	keys := []string{"photo", "in_contacts", "favouraite", "minAge", "maxAge", "height", "distance"}
	for _, key := range keys {
		if r.URL.Query()[key] == nil {
			log.Println("Missing query params : ", key)
			return false
		}
	}
	return true
}
