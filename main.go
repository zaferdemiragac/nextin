package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type AppointmentRouter struct {
}

func main() {

	if err := godotenv.Load(); err != nil { // ← .env'i oku, içindekileri ortam değişkeni yap
		log.Println(".env okunamadi, ortam degiskenleri kullanilacak")
	}

	connectDB()
	defer db.Close()

	repo := NewPostgresRepository(db)

	service := NewAppointmentService(repo)

	handler := NewAppointmentHandler(service)

	http.HandleFunc("DELETE /appointments/{id}", handler.Delete)

	http.HandleFunc("GET /appointments", handler.List)
	http.HandleFunc("POST /appointments", handler.Create)

	http.HandleFunc("GET /appointments/{id}", handler.GetByID)

	http.HandleFunc("GET /health", healthHandler)
	log.Println("sunucu :8080 portunda dinliyor")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) //log.println + os.exit(1)
}
