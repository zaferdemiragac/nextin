package main

type Appointment struct {
	ID      int    `json:"id"`
	Patient string `json:"patient"`
	Phone   string `json:"phone"`
	Date    string `json:"date"`
	Status  string `json:"status"`
}
