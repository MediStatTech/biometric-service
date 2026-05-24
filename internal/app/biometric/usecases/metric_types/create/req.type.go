package create

type Request struct {
	SensorID string
	Code     string
	Name     string
	Symbol   string
	MinValue float64
	MaxValue float64
}

type Response struct {
	MetricTypeID string
}
