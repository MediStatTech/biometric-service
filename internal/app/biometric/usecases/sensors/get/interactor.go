package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	sensorsRepo contracts.SensorsRepo
	logger      contracts.Logger
}

func New(
	sensorsRepo contracts.SensorsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorsRepo: sensorsRepo,
		logger:      logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	sensors, err := it.sensorsRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetSensors.SetInternal(err)
	}

	return &Response{
		Sensors: sensors,
	}, nil
}
