package metric_type

import "errors"

var (
	errRequestNil            = errors.New("request is nil")
	errInvalidMetricTypeData = errors.New("invalid metric type data")
)
