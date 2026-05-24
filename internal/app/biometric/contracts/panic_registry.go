package contracts

import "time"

type PanicRegistry interface {
	Trigger(patientID string, duration time.Duration) time.Time
	IsPanicking(patientID string) bool
}
