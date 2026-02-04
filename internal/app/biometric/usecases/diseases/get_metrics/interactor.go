package get_metrics

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
)

type Interactor struct {
	logger contracts.Logger
}

func New(logger contracts.Logger) *Interactor {
	return &Interactor{
		logger: logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	metricNames := []string{
		domain.DiseaseMetricEnumNameTemperature.String(),
		domain.DiseaseMetricEnumNameHeartRate.String(),
		domain.DiseaseMetricEnumNameBloodPressure.String(),
		domain.DiseaseMetricEnumNameBloodOxygen.String(),
	}

	return &Response{
		MetricNames: metricNames,
	}, nil
}
