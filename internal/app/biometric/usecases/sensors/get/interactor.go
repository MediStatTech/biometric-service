package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
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
	sensors, err := it.sensorsRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetSensors.SetInternal(err)
	}
	if len(sensors) == 0 {
		return &Response{Sensors: nil}, nil
	}

	metricTypes, err := it.metricTypesRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetSensors.SetInternal(err)
	}

	bySensor := make(map[string][]domain.MetricTypeProps, len(sensors))
	for _, mt := range metricTypes {
		bySensor[mt.SensorID] = append(bySensor[mt.SensorID], mt)
	}

	result := make([]SensorWithMetricTypes, 0, len(sensors))
	for _, s := range sensors {
		result = append(result, SensorWithMetricTypes{
			Sensor:      s,
			MetricTypes: bySensor[s.SensorID],
		})
	}

	return &Response{
		Sensors: result,
	}, nil
}
