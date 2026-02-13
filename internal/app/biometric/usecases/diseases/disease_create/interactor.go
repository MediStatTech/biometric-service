package disease_create

import (
	"context"
	"database/sql"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	diseasesRepo contracts.DiseasesRepo
	committer    contracts.Committer
	logger       contracts.Logger
}

func New(
	diseasesRepo contracts.DiseasesRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseasesRepo: diseasesRepo,
		committer:    committer,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.Name == "" || req.Code == "" {
		return nil, errInvalidRequest
	}

	_, err := it.diseasesRepo.FindByCode(ctx, req.Code)
	if err == nil {
		return nil, errDiseaseCodeExists
	}
	if err != sql.ErrNoRows {
		return nil, errFailedToCreateDisease.SetInternal(err)
	}

	now := time.Now().UTC()
	disease := domain.NewDisease(
		req.Name,
		req.Code,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.diseasesRepo.CreateMut(disease))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateDisease.SetInternal(err)
	}

	return &Response{
		DiseaseID: disease.DiseaseID(),
	}, nil
}
