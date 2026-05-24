package contracts

import "context"

type PatientDiseasesReader interface {
	FindByPatientID(ctx context.Context, patientID string) ([]string, error)
}
