package uc_options

import "github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"

type Options struct {
	Committer                contracts.Committer
	Logger                   contracts.Logger
	SensorsRepo              contracts.SensorsRepo
	SensorPatientsRepo       contracts.SensorPatientsRepo
	SensorPatientMetricsRepo contracts.SensorPatientMetricsRepo
}
