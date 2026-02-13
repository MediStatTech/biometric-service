package sensor

import (
	sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/create"
	sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/get"
	sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/retrieve"
	patient_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/create"
	patient_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/get"
	patient_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/retrieve"
	s_options "github.com/MediStatTech/biometric-service/internal/app/options"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	SensorCreate        *sensor_create.Interactor
	SensorPatientCreate *patient_create.Interactor
}

type Queries struct {
	SensorGet             *sensor_get.Interactor
	SensorRetrieve        *sensor_retrieve.Interactor
	SensorPatientGet      *patient_get.Interactor
	SensorPatientRetrieve *patient_retrieve.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			SensorCreate:        opts.App.Biometric.SensorCreate,
			SensorPatientCreate: opts.App.Biometric.SensorPatientCreate,
		},
		queries: &Queries{
			SensorGet:             opts.App.Biometric.SensorGet,
			SensorRetrieve:        opts.App.Biometric.SensorRetrieve,
			SensorPatientGet:      opts.App.Biometric.SensorPatientGet,
			SensorPatientRetrieve: opts.App.Biometric.SensorPatientRetrieve,
		},
	}
}
