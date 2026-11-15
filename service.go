package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var ErrGecersizGirdi = errors.New("geçersiz girdi")

type AppointmentService struct {
	repo AppointmentRepository //Service’te interface kullandık çünkü repository’yi testlerde değiştirmek istiyoruz
}

func NewAppointmentService(repo AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (s *AppointmentService) Create(ctx context.Context, a Appointment) (Appointment, error) {

	if strings.TrimSpace(a.Patient) == "" {
		return Appointment{}, fmt.Errorf("%w: hasta adi bos olamaz", ErrGecersizGirdi)
	}
	if strings.TrimSpace(a.Phone) == "" {
		return Appointment{}, fmt.Errorf("%w: telefon no  bos olamaz", ErrGecersizGirdi)
	}
	if strings.TrimSpace(a.Date) == "" {
		return Appointment{}, fmt.Errorf("%w: tarih bos olamaz", ErrGecersizGirdi)
	}

	if a.ClinicID <= 0 {
		return Appointment{}, fmt.Errorf("%w: klinik belirtilmeli", ErrGecersizGirdi)
	}

	if strings.TrimSpace(a.Status) == "" {

		a.Status = "beklemede"
	}

	return s.repo.Create(ctx, a)

}

func (s *AppointmentService) List(ctx context.Context) ([]Appointment, error) {
	return s.repo.List(ctx)
}

func (s *AppointmentService) GetByID(ctx context.Context, id int) (Appointment, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AppointmentService) Delete(ctx context.Context, id int) (int64, error) {
	return s.repo.Delete(ctx, id)
}
