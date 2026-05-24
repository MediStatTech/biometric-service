package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
)

type Interactor struct {
	metricTypesRepo contracts.MetricTypesRepo
	logger          contracts.Logger
}

func New(
	metricTypesRepo contracts.MetricTypesRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		metricTypesRepo: metricTypesRepo,
		logger:          logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	var (
		metricTypes []domain.MetricTypeProps
		err         error
	)

	if req.SensorID != "" {
		metricTypes, err = it.metricTypesRepo.FindBySensorID(ctx, req.SensorID)
	} else {
		metricTypes, err = it.metricTypesRepo.FindAll(ctx)
	}
	if err != nil {
		return nil, errFailedToGetMetricTypes.SetInternal(err)
	}

	return &Response{
		MetricTypes: metricTypes,
	}, nil
}
