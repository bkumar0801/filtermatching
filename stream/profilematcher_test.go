package stream

import (
	"reflect"
	"testing"

	"github.com/filtermatching/db"
	"github.com/filtermatching/filter"
)

func TestApplyFilter(t *testing.T) {
	profileMatcher := ProfileMatcher{
		Filter: filter.Filter{
			HasPhoto:           true,
			InContact:          true,
			Favouraite:         true,
			CompatibilityScore: 0.87,
			Age:                40,
			Height:             150,
			Distance:           40,
		},
		Profile: db.Profile{
			Display:  "Caroline",
			Age:      38,
			JobTitle: "Corporate Lawyer",
			Height:   143,
			Location: db.City{
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
	}

	profile := profileMatcher.Apply()
	if !reflect.DeepEqual(profile, profileMatcher.Profile) {
		t.Errorf("\nProfile mismatch:\n\texpected:%v\n\tgot:%v", profileMatcher.Profile, profile)
	}
}
