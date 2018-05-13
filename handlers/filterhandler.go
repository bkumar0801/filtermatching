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
		(w).Header().Set("Access-Control-Allow-Origin", "*")
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

		var subscriptions []stream.Subscription
		for _, profile := range profiles.Matches {
			matcher := stream.Apply(profile, *filter)
			subscription := stream.Subscribe(matcher)
			subscriptions = append(subscriptions, subscription)
		}
		var matchedProfiles db.Profiles
		for _, subscription := range subscriptions {
			matched := <-subscription.Updates()
			if !reflect.DeepEqual(matched, db.Profile{}) {
				matchedProfiles.Matches = append(matchedProfiles.Matches, matched)
			}
			subscription.Close()
		}
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
	maxAge, _ := strconv.ParseInt(r.FormValue("age"), 10, 32)
	maxHeight, _ := strconv.ParseInt(r.FormValue("height"), 10, 32)
	upperDistance, _ := strconv.ParseInt(r.FormValue("distance"), 10, 32)
	return f.NewFilter(photo, inContacts, favouraite, float32(score), int32(maxAge), int32(maxHeight), int32(upperDistance)), nil
}

func checkQueryParams(r *http.Request) bool {
	keys := []string{"photo", "in_contacts", "favouraite", "age", "height", "distance"}
	for _, key := range keys {
		if r.URL.Query()[key] == nil {
			log.Println("Missing query params : ", key)
			return false
		}
	}
	return true
}
