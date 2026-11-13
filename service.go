package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var ErrGecersizGirdi = errors.New("geçersiz girdi")

func createAppointmentService(ctx context.Context, a Appointment) (Appointment, error) {

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

	return createAppointmentInDB(ctx, a)

}

func listAppointmentsService(ctx context.Context) ([]Appointment, error) {
	return listAppointmentsFromDB(ctx)
}

func getAppointmentService(ctx context.Context, id int) (Appointment, error) {
	return getAppointmentFromDB(ctx, id)
}

func deleteAppointmentService(ctx context.Context, id int) (int64, error) {
	return deleteAppointmentFromDB(ctx, id)
}
