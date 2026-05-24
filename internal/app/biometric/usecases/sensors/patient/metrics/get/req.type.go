package get

import "time"

type Request struct {
	SensorID  string
	PatientID string
}

type Component struct {
	MetricTypeID string
	Code         string
	Name         string
	Value        float64
	Symbol       string
}

type Measurement struct {
	SensorID   string
	PatientID  string
	CreatedAt  time.Time
	Components []Component
}

type Response struct {
	Measurements []Measurement
}
