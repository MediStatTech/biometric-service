package create

type Request struct {
	SensorID  string
	PatientID string
}

type Response struct {
	SensorID  string
	PatientID string
	Status    string
}
