package usecases

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/create"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/get"
	patient_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/create"
	patient_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/get"
	patient_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/retrieve"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/uc_options"
)

type Facade struct {
	SensorCreate          *create.Interactor
	SensorGet             *get.Interactor
	SensorPatientCreate   *patient_create.Interactor
	SensorPatientGet      *patient_get.Interactor
	SensorPatientRetrieve *patient_retrieve.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		SensorCreate:          create.New(o.SensorsRepo, o.Committer, o.Logger),
		SensorGet:             get.New(o.SensorsRepo, o.Logger),
		SensorPatientCreate:   patient_create.New(o.SensorsRepo, o.SensorPatientsRepo, o.Committer, o.Logger),
		SensorPatientGet:      patient_get.New(o.SensorPatientsRepo, o.Logger),
		SensorPatientRetrieve: patient_retrieve.New(o.SensorPatientsRepo, o.Logger),
	}
}
