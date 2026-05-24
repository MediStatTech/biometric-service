package panic_registry

import (
	"sync"
	"time"
)

type Registry struct {
	mu    sync.RWMutex
	until map[string]time.Time
}

func New() *Registry {
	return &Registry{
		until: make(map[string]time.Time),
	}
}

func (r *Registry) Trigger(patientID string, d time.Duration) time.Time {
	r.mu.Lock()
	defer r.mu.Unlock()
	u := time.Now().Add(d)
	r.until[patientID] = u
	return u
}

func (r *Registry) IsPanicking(patientID string) bool {
	r.mu.RLock()
	u, ok := r.until[patientID]
	r.mu.RUnlock()
	if !ok {
		return false
	}
	if time.Now().Before(u) {
		return true
	}
	r.mu.Lock()
	delete(r.until, patientID)
	r.mu.Unlock()
	return false
}
