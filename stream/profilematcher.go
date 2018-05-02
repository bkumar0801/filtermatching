package stream

import (
	"github.com/filtermatching/constants"
	"github.com/filtermatching/db"
	"github.com/filtermatching/filter"
)

/*
ProfileMatcher ...
*/
type ProfileMatcher struct {
	Filter  filter.Filter
	Profile db.Profile
}

/*
Apply ...
*/
func (pf *ProfileMatcher) Apply() db.Profile {
	var matchedProfiles db.Profile
	if pf.Filter.HasPhoto != (len(pf.Profile.Photo) > 0) {
		return matchedProfiles
	}
	if pf.Filter.InContact != (pf.Profile.Contacts > 0) {
		return matchedProfiles
	}
	if pf.Filter.Favouraite != pf.Profile.Favourite {
		return matchedProfiles
	}
	if pf.Profile.Score > pf.Filter.CompatibilityScore {
		return matchedProfiles
	}
	if (pf.Profile.Age < constants.MinAge) || (pf.Profile.Age > pf.Filter.Age) {
		return matchedProfiles
	}
	if (pf.Profile.Height < constants.MinHeight) || (pf.Profile.Height > pf.Filter.Height) {
		return matchedProfiles
	}
	matchedProfiles = pf.Profile
	return matchedProfiles
}
