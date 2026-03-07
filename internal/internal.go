package internal

import (
	"context"
	"fmt"

	"github.com/MediStatTech/biometric-service/internal/app"
	grpc "github.com/MediStatTech/biometric-service/internal/transport/grpc"
	"github.com/MediStatTech/biometric-service/pkg"
)

func New(_ context.Context, p *pkg.Facade, appInstance *app.Facade) (*grpc.Server, error) {
	grpcServer, err := grpc.New(p, appInstance)
	if err != nil {
		return nil, fmt.Errorf("initialize grpc: %w", err)
	}

	return grpcServer, nil
}
