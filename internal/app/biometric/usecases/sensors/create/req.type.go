package create

type Request struct {
	Name   string
	Code   string
	Symbol string
}

type Response struct {
	SensorID string
}
