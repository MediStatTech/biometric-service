package options

import (
	"github.com/MediStatTech/biometric-service/internal/app"
	"github.com/MediStatTech/biometric-service/pkg"
)

type Options struct {
	App *app.Facade
	PKG *pkg.Facade
}
