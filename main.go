package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)

}

func main() {

	http.HandleFunc("GET /health", healthHandler)
	log.Println("sunucu :8080 portunda dinliyor")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) //log.println + os.exit(1)
}
