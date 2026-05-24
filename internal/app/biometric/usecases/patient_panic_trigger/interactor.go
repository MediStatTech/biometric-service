package patient_panic_trigger

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

const (
	defaultDuration = 60 * time.Second
	maxDuration     = 600 * time.Second
)

type Interactor struct {
	panicRegistry contracts.PanicRegistry
	logger        contracts.Logger
}

func New(
	panicRegistry contracts.PanicRegistry,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		panicRegistry: panicRegistry,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" {
		return nil, errInvalidRequest
	}

	d := time.Duration(req.DurationSeconds) * time.Second
	if d <= 0 {
		d = defaultDuration
	}
	if d > maxDuration {
		d = maxDuration
	}

	until := it.panicRegistry.Trigger(req.PatientID, d)

	it.logger.Info("[PanicTrigger] patient panic mode enabled", map[string]any{
		"patient_id":  req.PatientID,
		"duration_s":  d.Seconds(),
		"panic_until": until.Format(time.RFC3339),
	})

	return &Response{PanicUntil: until}, nil
}
