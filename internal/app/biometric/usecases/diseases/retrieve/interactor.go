package retrieve

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	diseasesRepo contracts.DiseasesRepo
	logger       contracts.Logger
}

func New(
	diseasesRepo contracts.DiseasesRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseasesRepo: diseasesRepo,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	disease, err := it.diseasesRepo.FindByID(ctx, req.DiseaseID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errDiseaseNotFound
		}
		return nil, errFailedToGetDisease.SetInternal(err)
	}

	return &Response{
		Disease: disease,
	}, nil
}
