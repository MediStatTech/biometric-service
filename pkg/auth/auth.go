package auth

import "context"

type Auth struct {
	StaffID    string
	PositionID string
}

// ===== context keys (unexported) =====

type staffIDKeyType struct{}
type positionIDKeyType struct{}

var (
	staffIDKey    = staffIDKeyType{}
	positionIDKey = positionIDKeyType{}
)

// ===== setters =====

func WithStaffID(ctx context.Context, staffID string) context.Context {
	return context.WithValue(ctx, staffIDKey, staffID)
}

func WithPositionID(ctx context.Context, positionID string) context.Context {
	return context.WithValue(ctx, positionIDKey, positionID)
}

func WithAuth(ctx context.Context, auth Auth) context.Context {
	ctx = WithStaffID(ctx, auth.StaffID)
	ctx = WithPositionID(ctx, auth.PositionID)
	return ctx
}

// ===== getters =====

func GetAuth(ctx context.Context) Auth {
	staffID := ctx.Value(staffIDKey).(string)
	positionID := ctx.Value(positionIDKey).(string)

	return Auth{
		StaffID:    staffID,
		PositionID: positionID,
	}
}

func GetStaffID(ctx context.Context) (string, bool) {
	staffID, ok := ctx.Value(staffIDKey).(string)
	return staffID, ok
}

func GetPositionID(ctx context.Context) (string, bool) {
	positionID, ok := ctx.Value(positionIDKey).(string)
	return positionID, ok
}
