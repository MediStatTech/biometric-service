package patient_status_batch_get

type Request struct {
	PatientIDs []string
}

type PatientStatus struct {
	PatientID string
	Status    string
}

type Response struct {
	Statuses []PatientStatus
}
