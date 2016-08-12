// Test utilities for http functionality

package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type HTTPRoundTripper struct {
	RoundTripF func(r *http.Request) (*http.Response, error)
}

func (rt *HTTPRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if rt.RoundTripF != nil {
		return rt.RoundTripF(r)
	}

	return nil, nil
}

type HTTPTestResponse struct {
	Status int
	Body   string
}

type HTTPTestServer struct {
	count     int
	responses []HTTPTestResponse
}

func (s *HTTPTestServer) Start(t *testing.T) (*url.URL, func()) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { s.count += 1 }()
		if len(s.responses) == 1 {
			r := s.responses[0]
			w.WriteHeader(r.Status)
			fmt.Fprintln(w, r.Body)
			return
		}

		for i, r := range s.responses {
			if s.count == i {
				w.WriteHeader(r.Status)
				fmt.Fprintln(w, r.Body)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}))

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
	}

	return u, ts.Close
}

// Creates a new http test server, allows for multiple different responses based on the
// request count and the numberof http responses provided, if only 1 is provided then
// it will always return this response, else more than one will return the response for that
// requtest number, for example the second request would write the second response
func NewHTTPTestServer(responses []HTTPTestResponse) *HTTPTestServer {
	return &HTTPTestServer{
		responses: responses,
		count:     0,
	}
}
