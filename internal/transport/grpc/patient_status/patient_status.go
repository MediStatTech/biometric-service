package patient_status

import (
	patient_panic_trigger "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/patient_panic_trigger"
	patient_status_batch_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/patient_status_batch_get"
	s_options "github.com/MediStatTech/biometric-service/internal/app/options"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	PatientPanicTrigger *patient_panic_trigger.Interactor
}

type Queries struct {
	PatientStatusGetBatch *patient_status_batch_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			PatientPanicTrigger: opts.App.Biometric.PatientPanicTrigger,
		},
		queries: &Queries{
			PatientStatusGetBatch: opts.App.Biometric.PatientStatusGetBatch,
		},
	}
}
