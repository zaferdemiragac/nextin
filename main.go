package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	err := json.NewDecoder(r.Body).Decode(&newappointment)
	if err != nil {
		http.Error(w, "geçersiz JSON", http.StatusBadRequest)
		return
	}

	newappointment.ID = nextID

	nextID += 1

	appointments = append(appointments, newappointment)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newappointment)

}

func getAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)

	if err != nil {

		http.Error(w, "geçersiz id", http.StatusBadRequest)
		return
	}

	for _, a := range appointments {
		if a.ID == id {

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(a)
			return

		}
	}

	http.Error(w, "randevu bulunamadı", http.StatusNotFound)

}

func deleteAppointment(w http.ResponseWriter, r *http.Request) {

	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		http.Error(w, "id bulunamadı", http.StatusBadRequest)
		return

	}

	for i, a := range appointments {

		if a.ID == id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "randevu bulunamadı", http.StatusNotFound)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)

}

func main() {

	http.HandleFunc("DELETE /appointments/{id}", deleteAppointment)

	http.HandleFunc("GET /appointments", listAppointments)
	http.HandleFunc("POST /appointments", createAppointment)

	http.HandleFunc("GET /appointments/{id}", getAppointment)

	http.HandleFunc("GET /health", healthHandler)
	log.Println("sunucu :8080 portunda dinliyor")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) //log.println + os.exit(1)
}
