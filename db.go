package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool" // ← PostgreSQL bağlantı havuzu
)

var db *pgxpool.Pool

func connectDB() {

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("DATABASE_URL tanımlı değil")
	}

	pool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal("havuz oluşturulmadı")
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("veritabanına bağlanılamadı", err)
	}

	db = pool

	log.Println("veritabanına bağlanıldı")
}
