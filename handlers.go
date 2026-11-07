package main

import(
	"net/http"
	"encoding/json"
	"strconv"
)






func listAppointments(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)

}

func createAppointment(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

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

	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)

	if err != nil {

		http.Error(w, "geçersiz id", http.StatusBadRequest)
		return
	}

	for _, a := range appointments {
		if a.ID == id {

			json.NewEncoder(w).Encode(a)
			return

		}
	}

	http.Error(w, "randevu bulunamadı", http.StatusNotFound)

}

func deleteAppointment(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		http.Error(w, "geçersiz id", http.StatusBadRequest)
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