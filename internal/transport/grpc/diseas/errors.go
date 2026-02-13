package diseas

import (
	"errors"
)

var (
	errRequestNil              = errors.New("request is nil")
	errFailedToCreateDisease   = errors.New("failed to create disease")
	errFailedToGetDiseases     = errors.New("failed to get diseases")
	errInvalidDiseaseData      = errors.New("invalid disease data")
)
