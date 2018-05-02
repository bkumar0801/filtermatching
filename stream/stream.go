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
	Close()
}

/*
Subscribe ...
*/
func Subscribe(matcher Matcher) Subscription {
	s := &sub{
		matcher: matcher,
		updates: make(chan db.Profile),
	}
	go s.filter()
	return s
}

type sub struct {
	matcher Matcher
	updates chan db.Profile
}

func (s *sub) Updates() <-chan db.Profile {
	return s.updates
}

func (s *sub) Close() {
	close(s.updates)
}

func (s *sub) filter() {
	s.updates <- s.matcher.Apply()
}
