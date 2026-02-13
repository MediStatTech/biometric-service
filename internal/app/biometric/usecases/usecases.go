package usecases

import (
	disease_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/create"
	disease_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/get"
	disease_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/retrieve"
	disease_sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/create"
	disease_sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/get"
	disease_sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/retrieve"
	sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/create"
	sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/get"
	sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/retrieve"
	patient_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/create"
	patient_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/get"
	patient_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/retrieve"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/uc_options"
)

type Facade struct {
	SensorCreate          *sensor_create.Interactor
	SensorGet             *sensor_get.Interactor
	SensorRetrieve        *sensor_retrieve.Interactor
	SensorPatientCreate   *patient_create.Interactor
	SensorPatientGet      *patient_get.Interactor
	SensorPatientRetrieve *patient_retrieve.Interactor

	// Diseases
	DiseaseCreate          *disease_create.Interactor
	DiseaseGet             *disease_get.Interactor
	DiseaseRetrieve        *disease_retrieve.Interactor
	DiseaseSensorCreate    *disease_sensor_create.Interactor
	DiseaseSensorGet       *disease_sensor_get.Interactor
	DiseaseSensorRetrieve  *disease_sensor_retrieve.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		SensorCreate:          sensor_create.New(o.SensorsRepo, o.Committer, o.Logger),
		SensorGet:             sensor_get.New(o.SensorsRepo, o.Logger),
		SensorRetrieve:        sensor_retrieve.New(o.SensorsRepo, o.Logger),
		SensorPatientCreate:   patient_create.New(o.SensorsRepo, o.SensorPatientsRepo, o.Committer, o.Logger),
		SensorPatientGet:      patient_get.New(o.SensorPatientsRepo, o.Logger),
		SensorPatientRetrieve: patient_retrieve.New(o.SensorPatientsRepo, o.Logger),

		DiseaseCreate:         disease_create.New(o.DiseasesRepo, o.Committer, o.Logger),
		DiseaseGet:            disease_get.New(o.DiseasesRepo, o.Logger),
		DiseaseRetrieve:       disease_retrieve.New(o.DiseasesRepo, o.Logger),
		DiseaseSensorCreate:   disease_sensor_create.New(o.DiseaseSensorsRepo, o.Committer, o.Logger),
		DiseaseSensorGet:      disease_sensor_get.New(o.DiseaseSensorsRepo, o.Logger),
		DiseaseSensorRetrieve: disease_sensor_retrieve.New(o.DiseaseSensorsRepo, o.Logger),
	}
}
