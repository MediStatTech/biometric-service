package create

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	sensorsRepo contracts.SensorsRepo
	committer   contracts.Committer
	logger      contracts.Logger
}

func New(
	sensorsRepo contracts.SensorsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorsRepo: sensorsRepo,
		committer:   committer,
		logger:      logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	now := time.Now().UTC()
	sensor := domain.NewSensor(
		req.Name,
		req.Code,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.sensorsRepo.CreateMut(sensor))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateSensor.SetInternal(err)
	}

	return &Response{
		SensorID: sensor.SensorID(),
	}, nil
}
