package main

import (
	"context"
)

func listAppointmentsFromDB(ctx context.Context) ([]Appointment, error) {

	rows, err := db.Query(ctx, //rows db satırları üzerinde hazır func ları olan bir tutamak

		`SELECT id, patient, phone, appointment_date::text, status, clinic_id
		 FROM appointments
		 ORDER BY id`)

	if err != nil {

		return nil, err

	}

	defer rows.Close()

	list := []Appointment{}

	for rows.Next() {
		var a Appointment

		if err := rows.Scan(&a.ID, &a.Patient, &a.Phone, &a.Date, &a.Status, &a.ClinicID); err != nil {
			return nil, err
		}

		list = append(list, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil

}

func getAppointmentFromDB(ctx context.Context, id int) (Appointment, error) { //QueryRow tek bir satır okur.

	var a Appointment

	err := db.QueryRow(ctx, `SELECT id, patient, phone, appointment_date::text, status, clinic_id
		 FROM appointments
		 WHERE id = $1`, id).Scan(&a.ID, &a.Patient, &a.Phone, &a.Date, &a.Status, &a.ClinicID)

	if err != nil {
		return Appointment{}, err
	}

	return a, nil

}

func createAppointmentInDB(ctx context.Context, a Appointment) (Appointment, error) {

	err := db.QueryRow(ctx, `
	INSERT INTO appointments (patient, phone, appointment_date, status, clinic_id)
	VALUES ($1, $2, $3::date, $4, $5)
	RETURNING id`,
		a.Patient, a.Phone, a.Date, a.Status, a.ClinicID).Scan(&a.ID)

	if err != nil {
		return Appointment{}, err
	}

	return a, nil

}

func deleteAppointmentFromDB(ctx context.Context, id int) (int64, error) {

	tag, err := db.Exec(ctx, `DELETE FROM appointments WHERE id = $1`, id)

	if err != nil {
		return 0, err
	}

	return tag.RowsAffected(), nil

}
