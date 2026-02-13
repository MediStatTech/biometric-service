package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	diseaseSensorsRepo contracts.DiseaseSensorsRepo
	logger       contracts.Logger
}

func New(
	diseaseSensorsRepo contracts.DiseaseSensorsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseaseSensorsRepo: diseaseSensorsRepo,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	diseaseSensors, err := it.diseaseSensorsRepo.FindByDiseaseID(ctx, req.DiseaseID)
	if err != nil {
		return nil, errFailedToGetDiseaseSensors.SetInternal(err)
	}

	return &Response{
		DiseaseSensors: diseaseSensors,
	}, nil
}
