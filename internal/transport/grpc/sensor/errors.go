package sensor

import (
	"errors"
)

var (
	errRequestNil          = errors.New("request is nil")
	errFailedToCreateSensor = errors.New("failed to create sensor")
	errFailedToGetSensors   = errors.New("failed to get sensors")
	errInvalidSensorData    = errors.New("invalid sensor data")
)
