package uc_options

import "github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"

type Options struct {
	Committer          contracts.Committer
	Logger             contracts.Logger
	DiseasesRepo       contracts.DiseasesRepo
	DiseaseMetricsRepo contracts.DiseaseMetricsRepo
	SensorsRepo        contracts.SensorsRepo
}
