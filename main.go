package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/exenale/blink_led_when_meeting/calendar"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "calendar data"}`))
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if val, ok := pathParams["deviceID"]; ok {
		token, _ := authDevice(val)
		if token {
			eventStatus := calendar.GetBusyStatus()
			w.Write([]byte(fmt.Sprintf(`{"deviceID": "%v", "eventStatus": "%v"}`, val, eventStatus)))
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("unauthorized")))
		return
	}
	// probably never be hit
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "need a device id"}`))

}

func authDevice(deviceID string) (bool, error) {
	if deviceID == "1234" {
		return true, nil
	}
	return false, nil
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)

	api.HandleFunc("/checkEvents/{deviceID}", getEvent).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
