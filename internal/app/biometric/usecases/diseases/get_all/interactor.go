package get_all

import (
	"context"

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
	diseases, err := it.diseasesRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetDiseases.SetInternal(err)
	}

	return &Response{
		Diseases: diseases,
	}, nil
}
