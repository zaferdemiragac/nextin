package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func listAppointments(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(r.Context(), `SELECT id, patient, phone, appointment_date::text, status, clinic_id
		 FROM appointments
		 ORDER BY id`) //r mesela kullanıcı sekmeyi kapattı bize bu bilgiyi verir

	if err != nil {
		log.Print("sorgu hatası:", err)
		http.Error(w, "sunucu hatası", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	list := []Appointment{}

	for rows.Next() { //ya false döner ya true
		var a Appointment

		if err := rows.Scan(&a.ID, &a.Patient, &a.Phone, &a.Date, &a.Status, &a.ClinicID); err != nil { //& bunu verince o sütunu bu değişkene yazıyor
			log.Println("satir okunamadi: ", err)
			http.Error(w, "sunucu hatasi", http.StatusInternalServerError)
			return //burdaki hata kontrölü mevcut satır okunamazsa
		}

		list = append(list, a)
	}

	if err := rows.Err(); err != nil { //burdaki hata
		log.Println("satirlar okunurken hata: ", err)
		http.Error(w, "sunucu hatasi", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(list)

}

func getAppointment(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {

		http.Error(w, "geçersiz id", http.StatusBadRequest)
		return
	}

	var a Appointment

	err = db.QueryRow(r.Context(), // ← tek satır bekliyoruz
		`SELECT id, patient, phone, appointment_date::text, status, clinic_id
		 FROM appointments
		 WHERE id = $1`, id). // ← $1 yer tutucu, değeri ayrı gidiyor
		Scan(&a.ID, &a.Patient, &a.Phone, &a.Date, &a.Status, &a.ClinicID)

	if errors.Is(err, pgx.ErrNoRows) { // ← o id'de kayıt yok
		http.Error(w, "randevu bulunamadi", http.StatusNotFound)
		return
	}

	if err != nil { // ← başka bir hata
		log.Println("sorgu hatasi:", err)
		http.Error(w, "sunucu hatasi", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)

}

func createAppointment(w http.ResponseWriter, r *http.Request) {

	var a Appointment

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, "geçersiz JSON", http.StatusBadRequest)
		return
	}

	err = db.QueryRow(r.Context(), `
	INSERT INTO appointments (patient, phone, appointment_date, status, clinic_id)
	VALUES ($1, $2, $3::date, $4, $5)
	RETURNING id`,
		a.Patient, a.Phone, a.Date, a.Status, a.ClinicID).Scan(&a.ID)

	if err != nil {
		log.Print("kayıt hatası:", err)
		http.Error(w, "sunucu hatası", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)

}

func deleteAppointment(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "geçersiz id", http.StatusBadRequest) //girilen id int deilse
		return

	}

	tag, err := db.Exec(r.Context(), // ← satır dönmeyen sorgu
		`DELETE FROM appointments WHERE id = $1`, id)
	if err != nil {
		log.Println("silme hatasi:", err)
		http.Error(w, "sunucu hatasi", http.StatusInternalServerError)
		return
	}

	if tag.RowsAffected() == 0 { //sql de silinecek o satır yoksa hata dönmez etkilenen satır sayısı bize durumu söyler.

		http.Error(w, "randevu bulunamadı", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)

}
