package uc_options

import "github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"

type Options struct {
	Committer                contracts.Committer
	Logger                   contracts.Logger
	// Repos
	SensorsRepo              contracts.SensorsRepo
	SensorPatientsRepo       contracts.SensorPatientsRepo
	SensorPatientMetricsRepo contracts.SensorPatientMetricsRepo

	DiseasesRepo              contracts.DiseasesRepo
	DiseaseSensorsRepo        contracts.DiseaseSensorsRepo
}
