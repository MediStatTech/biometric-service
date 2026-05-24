package retrieve

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	sensorsRepo     contracts.SensorsRepo
	metricTypesRepo contracts.MetricTypesRepo
	logger          contracts.Logger
}

func New(
	sensorsRepo contracts.SensorsRepo,
	metricTypesRepo contracts.MetricTypesRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorsRepo:     sensorsRepo,
		metricTypesRepo: metricTypesRepo,
		logger:          logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	sensor, err := it.sensorsRepo.FindByID(ctx, req.SensorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errSensorNotFound
		}
		return nil, errFailedToGetSensor.SetInternal(err)
	}

	metricTypes, err := it.metricTypesRepo.FindBySensorID(ctx, req.SensorID)
	if err != nil {
		return nil, errFailedToGetSensor.SetInternal(err)
	}

	return &Response{
		Sensor:      sensor,
		MetricTypes: metricTypes,
	}, nil
}
