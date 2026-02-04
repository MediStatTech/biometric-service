package usecases

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/create"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/get_all"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/get_metrics"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/uc_options"
)

type Facade struct {
	GetAllDiseases       *get_all.Interactor
	GetAllDiseaseMetrics *get_metrics.Interactor
	CreateSensors        *create.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		GetAllDiseases:       get_all.New(o.DiseasesRepo, o.Logger),
		GetAllDiseaseMetrics: get_metrics.New(o.Logger),
		CreateSensors:        create.New(o.DiseaseMetricsRepo, o.SensorsRepo, o.Committer, o.Logger),
	}
}
