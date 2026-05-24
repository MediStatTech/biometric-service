package usecases

import (
	disease_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/create"
	disease_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/get"
	disease_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/retrieve"
	disease_sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/create"
	disease_sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/get"
	disease_sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/retrieve"
	metric_type_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/metric_types/create"
	metric_type_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/metric_types/get"
	patient_panic_trigger "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/patient_panic_trigger"
	patient_status_batch_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/patient_status_batch_get"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/process"
	sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/create"
	sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/get"
	patient_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/create"
	patient_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/get"
	metric_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/get"
	metric_history "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/history"
	patient_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/retrieve"
	sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/retrieve"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/uc_options"
)

type Facade struct {
	SensorCreate           *sensor_create.Interactor
	SensorGet              *sensor_get.Interactor
	SensorRetrieve         *sensor_retrieve.Interactor
	SensorPatientCreate    *patient_create.Interactor
	SensorPatientGet       *patient_get.Interactor
	SensorPatientRetrieve  *patient_retrieve.Interactor
	SensorPatientMetricGet        *metric_get.Interactor
	SensorPatientMetricHistoryGet *metric_history.Interactor

	MetricTypeCreate *metric_type_create.Interactor
	MetricTypeGet    *metric_type_get.Interactor

	PatientStatusGetBatch *patient_status_batch_get.Interactor
	PatientPanicTrigger   *patient_panic_trigger.Interactor

	Process *process.Interactor

	DiseaseCreate         *disease_create.Interactor
	DiseaseGet            *disease_get.Interactor
	DiseaseRetrieve       *disease_retrieve.Interactor
	DiseaseSensorCreate   *disease_sensor_create.Interactor
	DiseaseSensorGet      *disease_sensor_get.Interactor
	DiseaseSensorRetrieve *disease_sensor_retrieve.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		SensorCreate:           sensor_create.New(o.SensorsRepo, o.Committer, o.Logger),
		SensorGet:              sensor_get.New(o.SensorsRepo, o.MetricTypesRepo, o.Logger),
		SensorRetrieve:         sensor_retrieve.New(o.SensorsRepo, o.MetricTypesRepo, o.Logger),
		SensorPatientCreate:    patient_create.New(o.SensorsRepo, o.SensorPatientsRepo, o.Committer, o.Logger),
		SensorPatientGet:       patient_get.New(o.SensorPatientsRepo, o.Logger),
		SensorPatientRetrieve:  patient_retrieve.New(o.SensorPatientsRepo, o.Logger),
		SensorPatientMetricGet:        metric_get.New(o.SensorPatientMetricsRepo, o.MetricTypesRepo, o.Logger),
		SensorPatientMetricHistoryGet: metric_history.New(o.SensorPatientMetricsRepo, o.MetricTypesRepo, o.Logger),

		MetricTypeCreate: metric_type_create.New(o.MetricTypesRepo, o.Committer, o.Logger),
		MetricTypeGet:    metric_type_get.New(o.MetricTypesRepo, o.Logger),

		PatientStatusGetBatch: patient_status_batch_get.New(o.SensorPatientMetricsRepo, o.MetricTypesRepo, o.Logger),
		PatientPanicTrigger:   patient_panic_trigger.New(o.PanicRegistry, o.Logger),

		Process: process.New(
			o.SensorPatientsRepo,
			o.SensorPatientMetricsRepo,
			o.MetricTypesRepo,
			o.DiseaseMetricOverridesRepo,
			o.PatientDiseasesReader,
			o.PanicRegistry,
			o.Committer,
			o.Logger,
		),

		DiseaseCreate:         disease_create.New(o.DiseasesRepo, o.Committer, o.Logger),
		DiseaseGet:            disease_get.New(o.DiseasesRepo, o.Logger),
		DiseaseRetrieve:       disease_retrieve.New(o.DiseasesRepo, o.Logger),
		DiseaseSensorCreate:   disease_sensor_create.New(o.DiseaseSensorsRepo, o.Committer, o.Logger),
		DiseaseSensorGet:      disease_sensor_get.New(o.DiseaseSensorsRepo, o.Logger),
		DiseaseSensorRetrieve: disease_sensor_retrieve.New(o.DiseaseSensorsRepo, o.Logger),
	}
}
