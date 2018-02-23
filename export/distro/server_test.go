//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	validNofificationUpdate string = `{"name": "somename", "operation":"add"}`
)

func TestServerPing(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/unused", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(replyPing)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `pong`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestServerNotifyRegistrations(t *testing.T) {
	var tests = []struct {
		name     string
		postData string
		status   int
	}{
		{"empty", "", 400},
		{"valid", validNofificationUpdate, 200},
		{"valid", `{"name": "somename", "operation":"delete"}`, 200},
		{"valid", `{"name": "somename", "operation":"update"}`, 200},
		{"invalid", `{"name": "somename", "operation":"upd"}`, 400},
		{"invalid", `{"operation":"update"}`, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// RefreshRegistrations is sending a NotifyUpdate to a channel each
			// time is called. We need to empty the channel between calls
			for len(registrationChanges) > 0 {
				<-registrationChanges
			}
			// Create a request to pass to our handler.
			req, err := http.NewRequest("GET", "/unused", strings.NewReader(tt.postData))
			if err != nil {
				t.Fatal(err)
			}

			handler := http.HandlerFunc(replyNotifyRegistrations)

			// We create a ResponseRecorder
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
		})
	}
}

func TestServerNotifyRegistrationsUnavailable(t *testing.T) {
	// RefreshRegistrations is sending a NotifyUpdate to a channel each
	// time is called. We need to empty the channel at the start
	for len(registrationChanges) > 0 {
		<-registrationChanges
	}
	for i := 0; i < notifyUpdateSize; i++ {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("GET", "/unused",
			strings.NewReader(validNofificationUpdate))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(replyNotifyRegistrations)

		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %d want %d",
				rr.Code, http.StatusOK)
		}
	}

	req, err := http.NewRequest("GET", "/unused",
		strings.NewReader(validNofificationUpdate))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(replyNotifyRegistrations)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %d want %d",
			rr.Code, http.StatusServiceUnavailable)
	}
}
