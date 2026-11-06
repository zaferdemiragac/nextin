package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Appointment struct {
	ID      int    `json:"id"`
	Patient string `json:"patient"`
	Phone   string `json:"phone"`
	Date    string `json:"date"`
	Status  string `json:"status"`
}

var appointments = []Appointment{}

var nextID = 1

func listAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)

}

func createAppointment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newappointment Appointment

	json.NewDecoder(r.Body).Decode(&newappointment)
	newappointment.ID = nextID

	nextID += 1

	appointments = append(appointments, newappointment)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newappointment)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)

}

func main() {

	http.HandleFunc("GET /appointments", listAppointments)
	http.HandleFunc("POST /appointments", createAppointment)

	http.HandleFunc("GET /health", healthHandler)
	log.Println("sunucu :8080 portunda dinliyor")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) //log.println + os.exit(1)
}
