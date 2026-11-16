package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type FakeRepository struct {
	data   []Appointment
	nextID int
}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{data: []Appointment{}, nextID: 1}
}

func (f *FakeRepository) List(ctx context.Context) ([]Appointment, error) {
	return f.data, nil // ← SQL yok, slice'ı döndür
}

func (f *FakeRepository) GetByID(ctx context.Context, id int) (Appointment, error) {
	for _, a := range f.data {
		if a.ID == id {
			return a, nil
		}
	}
	return Appointment{}, pgx.ErrNoRows // ← "bulunamadı" sinyali aynı kalmalı
}

func (f *FakeRepository) Create(ctx context.Context, a Appointment) (Appointment, error) {

	a.ID = f.nextID
	f.nextID += 1
	f.data = append(f.data, a)
	return a, nil
}

func (f *FakeRepository) Delete(ctx context.Context, id int) (int64, error) {

	for i, a := range f.data {

		if a.ID == id {

			f.data = append(f.data[:i], f.data[i+1:]...)
			return 1, nil

		}

	}

	return 0, nil

}
