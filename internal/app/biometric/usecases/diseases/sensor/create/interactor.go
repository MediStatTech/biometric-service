package create

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	diseaseSensorsRepo contracts.DiseaseSensorsRepo
	committer          contracts.Committer
	logger             contracts.Logger
}

func New(
	diseaseSensorsRepo contracts.DiseaseSensorsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseaseSensorsRepo: diseaseSensorsRepo,
		committer:          committer,
		logger:             logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.DiseaseID == "" || req.SensorID == "" {
		return nil, errInvalidRequest
	}

	now := time.Now().UTC()
	diseaseSensor := domain.NewDiseaseSensor(
		req.DiseaseID,
		req.SensorID,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.diseaseSensorsRepo.CreateMut(diseaseSensor))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateDiseaseSensor.SetInternal(err)
	}

	return &Response{
		DiseaseID: diseaseSensor.DiseaseID(),
		SensorID:  diseaseSensor.SensorID(),
	}, nil
}
