package process

import (
	"context"
	"math/rand"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
	"github.com/google/uuid"
)

// MetricRange defines the min/max range for a disease's metric values.
// TODO: fill in real disease_id -> range mappings
var diseaseMetricRanges = map[string]MetricRange{
	// Example: "disease-uuid-here": {Min: 60, Max: 100},
	"ce6d289f-73de-441c-b2af-8f781ab2d162": {Min: 3.9, Max: 10},
	"cb09a1b3-c33f-4090-9361-11567bde2443": {Min: 60, Max: 90},
	"d06fc6dd-976d-4a91-88e6-b60c37740a45": {Min: 95, Max: 100},

}

type MetricRange struct {
	Min float64
	Max float64
}

type Interactor struct {
	sensorsRepo              contracts.SensorsRepo
	sensorPatientsRepo       contracts.SensorPatientsRepo
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo
	diseaseSensorsRepo       contracts.DiseaseSensorsRepo
	committer                contracts.Committer
	logger                   contracts.Logger
}

func New(
	sensorsRepo contracts.SensorsRepo,
	sensorPatientsRepo contracts.SensorPatientsRepo,
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo,
	diseaseSensorsRepo contracts.DiseaseSensorsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorsRepo:              sensorsRepo,
		sensorPatientsRepo:       sensorPatientsRepo,
		sensorPatientMetricsRepo: sensorPatientMetricsRepo,
		diseaseSensorsRepo:       diseaseSensorsRepo,
		committer:                committer,
		logger:                   logger,
	}
}

func (it *Interactor) Execute(ctx context.Context) error {
	// 1. Get all sensors -> map[sensorID]SensorProps
	sensors, err := it.sensorsRepo.FindAll(ctx)
	if err != nil {
		it.logger.Error("[MetricGenerator] Failed to get sensors", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	sensorsMap := make(map[string]domain.SensorProps, len(sensors))
	for _, s := range sensors {
		sensorsMap[s.SensorID] = s
	}

	// 2. Get all disease_sensors -> map[sensorID][]diseaseID
	diseaseSensors, err := it.diseaseSensorsRepo.FindAll(ctx)
	if err != nil {
		it.logger.Error("[MetricGenerator] Failed to get disease_sensors", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	sensorDiseasesMap := make(map[string][]string)
	for _, ds := range diseaseSensors {
		sensorDiseasesMap[ds.SensorID] = append(sensorDiseasesMap[ds.SensorID], ds.DiseaseID)
	}

	// 3. Get all active sensor_patients
	activePatients, err := it.sensorPatientsRepo.FindByStatus(ctx, domain.SensorPatientStatusActive)
	if err != nil {
		it.logger.Error("[MetricGenerator] Failed to get active sensor_patients", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	if len(activePatients) == 0 {
		it.logger.Debug("[MetricGenerator] No active sensor_patients found", map[string]any{})
		return nil
	}

	// 4. Generate metrics for each active patient
	now := time.Now()
	var metrics []*domain.SensorPatientMetric

	for _, sp := range activePatients {
		sensor, ok := sensorsMap[sp.SensorID]
		if !ok {
			continue
		}

		value := generateMetricValue(sp.SensorID, sensorDiseasesMap)

		metric := domain.NewSensorPatientMetric(
			sp.SensorID,
			sp.PatientID,
			uuid.NewString(),
			value,
			sensor.Symbol,
			now,
		)
		metrics = append(metrics, metric)
	}

	if len(metrics) == 0 {
		return nil
	}

	// 5. Batch commit
	mutations := it.sensorPatientMetricsRepo.CreateBatchMut(metrics)
	plan := commitplan.NewPlan()
	plan.AddMuts(mutations...)

	if err := it.committer.Apply(ctx, plan); err != nil {
		it.logger.Error("[MetricGenerator] Failed to commit metrics", map[string]any{
			"error": err.Error(),
			"count": len(metrics),
		})
		return err
	}

	it.logger.Debug("[MetricGenerator] Generated metrics", map[string]any{
		"count": len(metrics),
	})

	return nil
}

func generateMetricValue(sensorID string, sensorDiseasesMap map[string][]string) float64 {
	diseaseIDs, ok := sensorDiseasesMap[sensorID]
	if ok {
		for _, diseaseID := range diseaseIDs {
			if r, exists := diseaseMetricRanges[diseaseID]; exists {
				return r.Min + rand.Float64()*(r.Max-r.Min)
			}
		}
	}

	// Default range if no disease mapping found
	return 60.0 + rand.Float64()*40.0
}
