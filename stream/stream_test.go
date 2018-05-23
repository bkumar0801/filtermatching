package stream

import (
	"reflect"
	"testing"

	"github.com/filtermatching/db"
	"github.com/filtermatching/filter"
)

func TestMatcherApply(t *testing.T) {
	filter := filter.Filter{
		HasPhoto:           true,
		InContact:          true,
		Favouraite:         true,
		CompatibilityScore: 0.87,
		MinAge:             18,
		MaxAge:             70,
		Height:             140,
		Distance:           40,
	}
	profile := db.Profile{
		Display:  "Caroline",
		Age:      40,
		JobTitle: "Corporate Lawyer",
		Height:   143,
		Location: db.City{
			Name:      "Leeds",
			Latitude:  53.801277,
			Longitude: -1.548567,
		},
		Photo:     "http://thecatapi.com/api/images/get?format=src&type=gif",
		Score:     0.80,
		Contacts:  2,
		Favourite: true,
		Religion:  "Atheist",
	}
	want := &ProfileMatcher{
		Filter:  filter,
		Profile: profile,
	}

	got := Apply(profile, filter)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nProfile mismatch:\n\texpected:%v\n\tgot:%v", want, got)
	}
}

func TestStreamSubscription(t *testing.T) {
	filter := filter.Filter{
		HasPhoto:           true,
		InContact:          true,
		Favouraite:         true,
		CompatibilityScore: 0.87,
		MinAge:             18,
		MaxAge:             70,
		Height:             140,
		Distance:           40,
	}
	profile := db.Profile{
		Display:  "Caroline",
		Age:      40,
		JobTitle: "Corporate Lawyer",
		Height:   135,
		Location: db.City{
			Name:      "Leeds",
			Latitude:  53.801277,
			Longitude: -1.548567,
		},
		Photo:     "http://thecatapi.com/api/images/get?format=src&type=gif",
		Score:     0.80,
		Contacts:  2,
		Favourite: true,
		Religion:  "Atheist",
	}
	matcher := &ProfileMatcher{
		Filter:  filter,
		Profile: profile,
	}

	subs := Subscribe(matcher)
	got := <-subs.Updates()
	if !reflect.DeepEqual(profile, got) {
		t.Errorf("\n Updates mismatch:\n\texpected:%v\n\tgot:%v", profile, got)
	}
	subs.Close()
}
