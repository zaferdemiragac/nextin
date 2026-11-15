package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type AppointmentHandler struct {
	service *AppointmentService
}

func NewAppointmentHandler(service *AppointmentService) *AppointmentHandler {

	return &AppointmentHandler{service: service}

}

func (h *AppointmentHandler) List(w http.ResponseWriter, r *http.Request) {

	list, err := h.service.List(r.Context())

	if err != nil {
		log.Println("listeleme hatası")
		http.Error(w, "sunucu hatası", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) //burada json verisi sunucumuza gelir.

}

func (h *AppointmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {

		http.Error(w, "geçersiz id", http.StatusBadRequest)
		return
	}

	a, err := h.service.GetByID(r.Context(), id)

	if errors.Is(err, pgx.ErrNoRows) {

		http.Error(w, "randevu bulunamadı", http.StatusNotFound) //neden listappointmenta sarmalamıyorduk hatayı?
		return

	}

	if err != nil {
		log.Println("getirme hatası:", err)
		http.Error(w, "sunucu hatası", http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)

}

func (h *AppointmentHandler) Create(w http.ResponseWriter, r *http.Request) {

	var a Appointment

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, "geçersiz JSON", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(r.Context(), a)

	if errors.Is(err, ErrGecersizGirdi) { // ← sarmalın içine bakar
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("kayit hatasi:", err)
		http.Error(w, "sunucu hatasi", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)

}

func (h *AppointmentHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "geçersiz id", http.StatusBadRequest) //girilen id int deilse
		return

	}

	rowsAffected, err := h.service.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, "sunucu hatasi", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "kayıt bulunamadı", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)

}
