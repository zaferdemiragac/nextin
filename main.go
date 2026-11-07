package main

import (
	"log"
	"net/http"
)

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
