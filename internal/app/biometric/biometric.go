package biometric

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/uc_options"
	"github.com/MediStatTech/biometric-service/internal/infra/repo"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	// Initialize repositories
	sensorsRepo := repo.NewSensorsRepository(pkg.Postgres.DB)
	sensorPatientsRepo := repo.NewSensorPatientsRepository(pkg.Postgres.DB)
	diseasesRepo := repo.NewDiseasesRepository(pkg.Postgres.DB)
	diseaseSensorsRepo := repo.NewDiseaseSensorsRepository(pkg.Postgres.DB)

	useCasesInstance := usecases.New(&uc_options.Options{
		Committer:          pkg.Committer,
		Logger:             pkg.Logger,
		SensorsRepo:        sensorsRepo,
		SensorPatientsRepo: sensorPatientsRepo,
		DiseasesRepo:       diseasesRepo,
		DiseaseSensorsRepo: diseaseSensorsRepo,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
