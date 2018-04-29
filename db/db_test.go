package db

import (
	"reflect"
	"testing"
)

func TestLoadProfiles(t *testing.T) {
	filename := "./testProfile.json"
	got, err := LoadProfiles(filename)
	if err != nil {
		t.Error("Error occurred while profile loading: ", err.Error())
	}

	want := Profiles{
		Matches: []Profile{
			{
				Display:  "Caroline",
				Age:      41,
				JobTitle: "Corporate Lawyer",
				Height:   153,
				Location: City{
					Name:      "Leeds",
					Latitude:  53.801277,
					Longitude: -1.548567,
				},
				Photo:     "http://thecatapi.com/api/images/get?format=src&type=gif",
				Score:     0.76,
				Contacts:  2,
				Favourite: true,
				Religion:  "Atheist",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nFetched profile mismatched: \n\t\t expected: %v, \n\t\t actual: %v", want, got)
	}
}
