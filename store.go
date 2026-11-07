package main

import "sync"

var mu sync.Mutex
var appointments = []Appointment{}

var nextID = 1
