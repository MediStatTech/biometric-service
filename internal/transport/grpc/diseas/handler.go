package diseas

import (
	disease_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/create"
	disease_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/get"
	disease_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/retrieve"
	disease_sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/create"
	disease_sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/get"
	disease_sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/retrieve"
	s_options "github.com/MediStatTech/biometric-service/internal/app/options"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	DiseaseCreate       *disease_create.Interactor
	DiseaseSensorCreate *disease_sensor_create.Interactor
}

type Queries struct {
	DiseaseGet            *disease_get.Interactor
	DiseaseRetrieve       *disease_retrieve.Interactor
	DiseaseSensorRetrieve *disease_sensor_retrieve.Interactor
	DiseaseSensorGet      *disease_sensor_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			DiseaseCreate:       opts.App.Biometric.DiseaseCreate,
			DiseaseSensorCreate: opts.App.Biometric.DiseaseSensorCreate,
		},
		queries: &Queries{
			DiseaseGet:            opts.App.Biometric.DiseaseGet,
			DiseaseRetrieve:       opts.App.Biometric.DiseaseRetrieve,
			DiseaseSensorRetrieve: opts.App.Biometric.DiseaseSensorRetrieve,
			DiseaseSensorGet:      opts.App.Biometric.DiseaseSensorGet,
		},
	}
}
