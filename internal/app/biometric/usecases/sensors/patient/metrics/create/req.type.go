package create

type Request struct {
	SensorID  string
	PatientID string
	MetricID  string
	Value     float64
}

type Response struct {
	SensorID  string
	PatientID string
	MetricID  string
	Value     float64
	CreatedAt string
}
