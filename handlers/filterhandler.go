package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	f "github.com/filtermatching/filter"
)

/*
HandleFilter ...
*/
func HandleFilter(w http.ResponseWriter, r *http.Request) {
	filter, err := BuildFilterFromQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Println(filter)
	fmt.Fprintf(w, `{"display_name": "test name"}`)
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
