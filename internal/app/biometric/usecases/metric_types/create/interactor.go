package create

import (
	"context"
	"database/sql"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	metricTypesRepo contracts.MetricTypesRepo
	committer       contracts.Committer
	logger          contracts.Logger
}

func New(
	metricTypesRepo contracts.MetricTypesRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		metricTypesRepo: metricTypesRepo,
		committer:       committer,
		logger:          logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.SensorID == "" || req.Code == "" || req.Name == "" {
		return nil, errInvalidRequest
	}
	if req.MinValue >= req.MaxValue {
		return nil, errInvalidRange
	}

	if _, err := it.metricTypesRepo.FindBySensorAndCode(ctx, req.SensorID, req.Code); err == nil {
		return nil, errMetricTypeCodeExists
	} else if err != sql.ErrNoRows {
		return nil, errFailedToCreateMetricType.SetInternal(err)
	}

	now := time.Now().UTC()
	mt := domain.NewMetricType(
		req.SensorID,
		req.Code,
		req.Name,
		req.Symbol,
		req.MinValue,
		req.MaxValue,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.metricTypesRepo.CreateMut(mt))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateMetricType.SetInternal(err)
	}

	return &Response{
		MetricTypeID: mt.MetricTypeID(),
	}, nil
}
