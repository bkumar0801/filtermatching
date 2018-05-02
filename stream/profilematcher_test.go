package stream

import (
	"reflect"
	"testing"

	"github.com/filtermatching/db"
	"github.com/filtermatching/filter"
)

func TestApplyFilter(t *testing.T) {
	cases := []struct {
		description    string
		profileMatcher ProfileMatcher
		want           db.Profile
	}{
		{
			description: "When filter match profile",
			profileMatcher: ProfileMatcher{
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
			},
			want: db.Profile{
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
		},
		{
			description: "When profile age is more than filter age",
			profileMatcher: ProfileMatcher{
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
					Age:      45,
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
			},
			want: db.Profile{},
		},
		{
			description: "When profile height is higher than filter height",
			profileMatcher: ProfileMatcher{
				Filter: filter.Filter{
					HasPhoto:           true,
					InContact:          true,
					Favouraite:         true,
					CompatibilityScore: 0.87,
					Age:                40,
					Height:             140,
					Distance:           40,
				},
				Profile: db.Profile{
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
				},
			},
			want: db.Profile{},
		},
		{
			description: "When profile compatibility score is higher than filter score",
			profileMatcher: ProfileMatcher{
				Filter: filter.Filter{
					HasPhoto:           true,
					InContact:          true,
					Favouraite:         true,
					CompatibilityScore: 0.87,
					Age:                40,
					Height:             140,
					Distance:           40,
				},
				Profile: db.Profile{
					Display:  "Caroline",
					Age:      40,
					JobTitle: "Corporate Lawyer",
					Height:   133,
					Location: db.City{
						Name:      "Leeds",
						Latitude:  53.801277,
						Longitude: -1.548567,
					},
					Photo:     "http://thecatapi.com/api/images/get?format=src&type=gif",
					Score:     0.89,
					Contacts:  2,
					Favourite: true,
					Religion:  "Atheist",
				},
			},
			want: db.Profile{},
		},
		{
			description: "When there is no profile photo",
			profileMatcher: ProfileMatcher{
				Filter: filter.Filter{
					HasPhoto:           true,
					InContact:          true,
					Favouraite:         true,
					CompatibilityScore: 0.87,
					Age:                40,
					Height:             140,
					Distance:           40,
				},
				Profile: db.Profile{
					Display:  "Caroline",
					Age:      40,
					JobTitle: "Corporate Lawyer",
					Height:   133,
					Location: db.City{
						Name:      "Leeds",
						Latitude:  53.801277,
						Longitude: -1.548567,
					},
					Score:     0.80,
					Contacts:  2,
					Favourite: true,
					Religion:  "Atheist",
				},
			},
			want: db.Profile{},
		},
		{
			description: "When person is not in contact",
			profileMatcher: ProfileMatcher{
				Filter: filter.Filter{
					HasPhoto:           true,
					InContact:          true,
					Favouraite:         true,
					CompatibilityScore: 0.87,
					Age:                40,
					Height:             140,
					Distance:           40,
				},
				Profile: db.Profile{
					Display:  "Caroline",
					Age:      40,
					JobTitle: "Corporate Lawyer",
					Height:   133,
					Location: db.City{
						Name:      "Leeds",
						Latitude:  53.801277,
						Longitude: -1.548567,
					},
					Photo:     "http://thecatapi.com/api/images/get?format=src&type=gif",
					Score:     0.80,
					Contacts:  0,
					Favourite: true,
					Religion:  "Atheist",
				},
			},
			want: db.Profile{},
		},
		{
			description: "When the profile favouraite is false",
			profileMatcher: ProfileMatcher{
				Filter: filter.Filter{
					HasPhoto:           true,
					InContact:          true,
					Favouraite:         true,
					CompatibilityScore: 0.87,
					Age:                40,
					Height:             140,
					Distance:           40,
				},
				Profile: db.Profile{
					Display:  "Caroline",
					Age:      40,
					JobTitle: "Corporate Lawyer",
					Height:   133,
					Location: db.City{
						Name:      "Leeds",
						Latitude:  53.801277,
						Longitude: -1.548567,
					},
					Photo:     "http://thecatapi.com/api/images/get?format=src&type=gif",
					Score:     0.80,
					Contacts:  1,
					Favourite: false,
					Religion:  "Atheist",
				},
			},
			want: db.Profile{},
		},
	}

	for _, c := range cases {
		got := c.profileMatcher.Apply()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nProfile mismatch:\n\texpected:%v\n\tgot:%v", c.want, got)
		}
	}
}
