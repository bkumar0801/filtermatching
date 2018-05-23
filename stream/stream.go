package stream

import (
	"github.com/filtermatching/db"
	f "github.com/filtermatching/filter"
)

/*
Matcher ...
*/
type Matcher interface {
	Apply() db.Profile
}

/*
Apply ...
*/
func Apply(profile db.Profile, filter f.Filter) Matcher {
	return &ProfileMatcher{
		Profile: profile,
		Filter:  filter,
	}
}

/*
Subscription ...
*/
type Subscription interface {
	Updates() <-chan db.Profile
}

/*
Subscribe ...
*/
func Subscribe(in chan db.Profile, matcher Matcher) Subscription {
	s := &sub{
		matcher: matcher,
		profile: in,
	}
	go s.filter()
	return s
}

type sub struct {
	matcher Matcher
	profile chan db.Profile
}

func (s *sub) filter() {
	s.profile <- s.matcher.Apply()
}

func (s *sub) Updates() <-chan db.Profile {
	return s.profile
}
