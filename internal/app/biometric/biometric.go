package biometric

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/uc_options"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {

	useCasesInstance := usecases.New(&uc_options.Options{
		Committer: pkg.Committer,
		Logger:    pkg.Logger,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
