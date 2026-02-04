package create

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	diseaseMetricsRepo contracts.DiseaseMetricsRepo
	sensorsRepo        contracts.SensorsRepo
	committer          contracts.Committer
	logger             contracts.Logger
}

func New(
	diseaseMetricsRepo contracts.DiseaseMetricsRepo,
	sensorsRepo contracts.SensorsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseaseMetricsRepo: diseaseMetricsRepo,
		sensorsRepo:        sensorsRepo,
		committer:          committer,
		logger:             logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	metrics, err := it.diseaseMetricsRepo.FindByDiseaseID(ctx, req.DiseaseID)
	if err != nil {
		return nil, errFailedToGetDiseaseMetrics.SetInternal(err)
	}

	if len(metrics) == 0 {
		return nil, errNoDiseaseMetricsFound
	}

	now := time.Now().UTC()
	sensors := make([]*domain.Sensor, 0, len(metrics))
	sensorIDs := make([]string, 0, len(metrics))

	for _, metric := range metrics {
		sensor := domain.NewSensor(
			metric.Name,    
			metric.EnumName, 
			now,
		)
		sensors = append(sensors, sensor)
		sensorIDs = append(sensorIDs, sensor.SensorID())
	}

	plan := commitplan.NewPlan()
	for _, sensor := range sensors {
		plan.AddMut(it.sensorsRepo.CreateMut(sensor))
	}

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateSensors.SetInternal(err)
	}

	return &Response{
		SensorIDs: sensorIDs,
	}, nil
}
