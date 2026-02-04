package contracts

import (
	"context"

	"github.com/MediStatTech/commitplan"
	"github.com/MediStatTech/jwt"
)

type Committer interface {
	Apply(ctx context.Context, p commitplan.PlanLike, opts ...commitplan.ApplyOption) error
}

// Logger is the interface for logging
type Logger interface {
	Debug(msg string, fields map[string]any)
	Info(msg string, fields map[string]any)
	Warn(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
	Fatal(msg string, fields map[string]any)
}

type JWT interface {
	Generate(staff_id, position_id string) (string, error)
	Verify(token string) (*jwt.Claims, error)
}
