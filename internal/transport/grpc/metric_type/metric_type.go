package metric_type

import (
	mt_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/metric_types/create"
	mt_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/metric_types/get"
	s_options "github.com/MediStatTech/biometric-service/internal/app/options"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	MetricTypeCreate *mt_create.Interactor
}

type Queries struct {
	MetricTypeGet *mt_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			MetricTypeCreate: opts.App.Biometric.MetricTypeCreate,
		},
		queries: &Queries{
			MetricTypeGet: opts.App.Biometric.MetricTypeGet,
		},
	}
}
