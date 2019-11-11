package main

import (
	"fmt"
	"log"
	"net/http"

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

		deviceID := val
		w.Write([]byte(fmt.Sprintf(`{"deviceID": %v }`, deviceID)))
		return
	}
	// probably never be hit
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "need a device id"}`))

}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)

	api.HandleFunc("/checkEvents/{deviceID}", getEvent).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
