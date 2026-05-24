package process

import (
	"context"
	"math/rand"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	sensorPatientsRepo         contracts.SensorPatientsRepo
	sensorPatientMetricsRepo   contracts.SensorPatientMetricsRepo
	metricTypesRepo            contracts.MetricTypesRepo
	diseaseMetricOverridesRepo contracts.DiseaseMetricOverridesRepo
	patientDiseasesReader      contracts.PatientDiseasesReader
	panicRegistry              contracts.PanicRegistry
	committer                  contracts.Committer
	logger                     contracts.Logger
}

func New(
	sensorPatientsRepo contracts.SensorPatientsRepo,
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo,
	metricTypesRepo contracts.MetricTypesRepo,
	diseaseMetricOverridesRepo contracts.DiseaseMetricOverridesRepo,
	patientDiseasesReader contracts.PatientDiseasesReader,
	panicRegistry contracts.PanicRegistry,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorPatientsRepo:         sensorPatientsRepo,
		sensorPatientMetricsRepo:   sensorPatientMetricsRepo,
		metricTypesRepo:            metricTypesRepo,
		diseaseMetricOverridesRepo: diseaseMetricOverridesRepo,
		patientDiseasesReader:      patientDiseasesReader,
		panicRegistry:              panicRegistry,
		committer:                  committer,
		logger:                     logger,
	}
}

func (it *Interactor) Execute(ctx context.Context) error {
	activePatients, err := it.sensorPatientsRepo.FindByStatus(ctx, domain.SensorPatientStatusActive)
	if err != nil {
		it.logger.Error("[MetricGenerator] Failed to get active sensor_patients", map[string]any{
			"error": err.Error(),
		})
		return err
	}
	if len(activePatients) == 0 {
		return nil
	}

	metricTypes, err := it.metricTypesRepo.FindAll(ctx)
	if err != nil {
		it.logger.Error("[MetricGenerator] Failed to get metric_types", map[string]any{
			"error": err.Error(),
		})
		return err
	}
	if len(metricTypes) == 0 {
		return nil
	}

	overrides, err := it.diseaseMetricOverridesRepo.FindAll(ctx)
	if err != nil {
		it.logger.Error("[MetricGenerator] Failed to get disease_metric_overrides", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	typesBySensor := make(map[string][]domain.MetricTypeProps, len(metricTypes))
	for _, mt := range metricTypes {
		typesBySensor[mt.SensorID] = append(typesBySensor[mt.SensorID], mt)
	}

	overridesByDiseaseMT := make(map[string]map[string]domain.DiseaseMetricOverrideProps)
	for _, o := range overrides {
		m, ok := overridesByDiseaseMT[o.DiseaseID]
		if !ok {
			m = make(map[string]domain.DiseaseMetricOverrideProps)
			overridesByDiseaseMT[o.DiseaseID] = m
		}
		m[o.MetricTypeID] = o
	}

	patientDiseases := make(map[string][]string)
	for _, sp := range activePatients {
		if _, cached := patientDiseases[sp.PatientID]; cached {
			continue
		}
		diseases, err := it.patientDiseasesReader.FindByPatientID(ctx, sp.PatientID)
		if err != nil {
			it.logger.Warn("[MetricGenerator] Failed to get patient_diseases", map[string]any{
				"patient_id": sp.PatientID,
				"error":      err.Error(),
			})
			diseases = nil
		}
		patientDiseases[sp.PatientID] = diseases
	}

	now := time.Now().UTC()
	var metrics []*domain.SensorPatientMetric

	for _, sp := range activePatients {
		types, ok := typesBySensor[sp.SensorID]
		if !ok || len(types) == 0 {
			continue
		}
		diseases := patientDiseases[sp.PatientID]
		inPanic := it.panicRegistry.IsPanicking(sp.PatientID)

		for _, mt := range types {
			minVal, maxVal := mt.MinValue, mt.MaxValue
			for _, did := range diseases {
				if mts, ok := overridesByDiseaseMT[did]; ok {
					if o, ok := mts[mt.MetricTypeID]; ok {
						minVal, maxVal = o.MinValue, o.MaxValue
						break
					}
				}
			}

			var value float64
			if inPanic {
				value = generateCriticalValue(minVal, maxVal)
			} else {
				value = minVal + rand.Float64()*(maxVal-minVal)
			}

			metrics = append(metrics, domain.NewSensorPatientMetric(
				sp.SensorID,
				sp.PatientID,
				mt.MetricTypeID,
				value,
				mt.Symbol,
				now,
			))
		}
	}

	if len(metrics) == 0 {
		return nil
	}

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

func generateCriticalValue(minVal, maxVal float64) float64 {
	rangeSize := maxVal - minVal
	if rangeSize <= 0 {
		return maxVal
	}
	low := maxVal + rangeSize*0.3
	high := maxVal + rangeSize*0.6
	return low + rand.Float64()*(high-low)
}
