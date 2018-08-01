package muxhandler

import (
	"testing"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"io"
	"io/ioutil"
	"log"
)

func TestTopicHandler(t *testing.T){

	router := mux.NewRouter()
	topicHanlder := NewTopicHandler()
	router.HandleFunc("/topic", topicHanlder.HandleAdd).Methods("POST")
	router.HandleFunc("/topic/{id}", topicHanlder.HandleGet).Methods("GET")
	router.HandleFunc("/topic", topicHanlder.HandleModify).Methods("PUT")
	router.HandleFunc("/topic/{id}", topicHanlder.HandleDelete).Methods("DELETE")
	router.HandleFunc("/topic", topicHanlder.HandleGetAll).Methods("GET")
	tt := []struct{
		method string
		path string
		requestBody string
		result      string
	}{
		{"POST", "/topic",`{"id":1,"title":"The Go Standard Library1","content":"It contains many packages1."}`, "HandleAdd Success."},
		{"POST", "/topic",`{"id":2,"title":"The Go Standard Library2","content":"It contains many packages2."}`, "HandleAdd Success."},
		{"GET", "/topic","", "HandleAdd Success."},
		{"GET", "/topic/%s","1", `{"id":1,"title":"The Go Standard Library1","content":"It contains many packages1.","created_at":"0001-01-01T00:00:00Z"}`},
		{"PUT", "/topic",`{"id":1,"title":"The Go Standard Library By Example","content":"It contains many packages, enjoying it."}`, "HandleModify Success."},
		{"DELETE", "/topic/%s","1", "HandleDelete Success."},
	}

	for _, tc := range tt {
		var path string
		var body io.Reader
		switch tc.method {
		case "POST":
			path = tc.path
			body = strings.NewReader(tc.requestBody)
		case "GET":
			if tc.requestBody != ""{
				path = fmt.Sprintf(tc.path, tc.requestBody)
			}else{
				path = tc.path
			}
			body = nil
		case "PUT":
			path = tc.path
			body = strings.NewReader(tc.requestBody)
		case "DELETE":
			path = fmt.Sprintf(tc.path, tc.requestBody)
			body = nil
		}

		req, err := http.NewRequest(tc.method, path, body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router.ServeHTTP(rr, req)

		// In this case, our MetricsHandler returns a non-200 response
		// for a route variable it doesn't know about.
		bodyStr, readErr := ioutil.ReadAll(rr.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		if rr.Code != http.StatusOK || tc.result != string(bodyStr) {
			t.Errorf(" %s, %s: got %v want %v",
				string(bodyStr), tc.result, rr.Code, http.StatusOK)
		}
	}
}
