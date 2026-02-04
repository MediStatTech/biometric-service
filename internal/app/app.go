package app

import (
	"fmt"

	"github.com/MediStatTech/biometric-service/internal/app/biometric"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Facade struct {
	Biometric *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	biometric, err := biometric.New(pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to create biometric: %w", err)
	}

	return &Facade{
		Biometric: biometric.UseCases,
	}, nil
}
