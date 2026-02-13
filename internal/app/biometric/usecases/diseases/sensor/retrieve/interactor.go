package retrieve

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	diseaseSensorsRepo contracts.DiseaseSensorsRepo
	logger             contracts.Logger
}

func New(
	diseaseSensorsRepo contracts.DiseaseSensorsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseaseSensorsRepo: diseaseSensorsRepo,
		logger:             logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	diseaseSensor, err := it.diseaseSensorsRepo.FindByDiseaseAndSensor(ctx, req.DiseaseID, req.SensorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errDiseaseSensorNotFound
		}
		return nil, errFailedToGetDiseaseSensor.SetInternal(err)
	}

	return &Response{
		DiseaseSensor: diseaseSensor,
	}, nil
}
