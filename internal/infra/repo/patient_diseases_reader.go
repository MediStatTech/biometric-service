package repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const listPatientDiseasesByPatientID = `
SELECT diseas_id
FROM patient_diseases
WHERE patient_id = $1
`

type PatientDiseasesReader struct {
	db *sql.DB
}

func NewPatientDiseasesReader(db *sql.DB) *PatientDiseasesReader {
	return &PatientDiseasesReader{db: db}
}

func (r *PatientDiseasesReader) FindByPatientID(ctx context.Context, patientID string) ([]string, error) {
	pid, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, listPatientDiseasesByPatientID, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []string{}
	for rows.Next() {
		var did uuid.UUID
		if err := rows.Scan(&did); err != nil {
			return nil, err
		}
		result = append(result, did.String())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
