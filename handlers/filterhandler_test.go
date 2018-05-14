package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/filtermatching/filter"
)

func TestHandlerFilter(t *testing.T) {
	cases := []struct {
		uri    string
		status int
	}{
		{
			uri:    "/filter?photo=true&in_contacts=true&favouraite=true&compatibility_score=0.80&age=70&height=210&distance=300",
			status: http.StatusOK,
		},
		{
			uri:    "/filter?photo=true&in_contacts=true",
			status: http.StatusUnprocessableEntity,
		},
	}

	for _, c := range cases {
		r, err := http.NewRequest("GET", c.uri, nil)
		if err != nil {
			t.Error("NewRequest Error: ", err.Error())
		}
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(HandleFilter("../db/matches.json"))
		handler.ServeHTTP(w, r)
		if c.status != w.Code {
			t.Errorf("\nHandler returned wrong status code:\n\texpected:%d\n\tgot:%d", c.status, w.Code)
		}
	}
}

func TestBuildFilterFromQuery(t *testing.T) {
	cases := []struct {
		description string
		uri         string
		want        *filter.Filter
		wantErr     error
	}{
		{
			description: "When all query parameters are given",
			uri:         "/filter?photo=true&in_contacts=true&favouraite=true&compatibility_score=0.80&age=70&height=210&distance=300",
			want: &filter.Filter{
				HasPhoto:           true,
				InContact:          true,
				Favouraite:         true,
				CompatibilityScore: 0.80,
				Age:                70,
				Height:             210,
				Distance:           300,
			},
			wantErr: nil,
		},
		{
			description: "When compatibility score is less than and equal to zero",
			uri:         "/filter?photo=true&in_contacts=true&favouraite=true&compatibility_score=0&age=70&height=210&distance=300",
			want: &filter.Filter{
				HasPhoto:           true,
				InContact:          true,
				Favouraite:         true,
				CompatibilityScore: 0.01,
				Age:                70,
				Height:             210,
				Distance:           300,
			},
			wantErr: nil,
		},
		{
			description: "When there is missing query params",
			uri:         "/filter?photo=true&in_contacts=true&favouraite=true",
			want:        nil,
			wantErr:     errors.New("missing query params"),
		},
	}

	for _, c := range cases {
		r, err := http.NewRequest("GET", c.uri, nil)
		if err != nil {
			t.Error("NewRequest Error: ", err.Error())
		}
		got, gotErr := BuildFilterFromQuery(r)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nUnexpected output: \n\t\t expected: %v, \n\t\t actual: %v", c.want, got)
		}

		if !reflect.DeepEqual(gotErr, c.wantErr) {
			t.Errorf("\nError mismatch: \n\t\t expected: %s, \n\t\t actual: %s", c.wantErr.Error(), gotErr.Error())
		}
	}
}
