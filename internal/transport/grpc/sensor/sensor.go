package sensor

import (
	"context"
	"time"

	sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/create"
	sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/get"
	patient_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/create"
	patient_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/get"
	metric_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/get"
	metric_history "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/history"
	patient_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/retrieve"
	sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/retrieve"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
	pb_services "github.com/MediStatTech/biometric-client/pb/go/services/v1"
)

func (h *Handler) SensorCreate(
	ctx context.Context,
	req *pb_services.SensorCreateRequest,
) (*pb_services.SensorCreateReply, error) {
	if req == nil || req.Sensor == nil {
		return nil, errRequestNil
	}
	if req.Sensor.Name == "" || req.Sensor.Code == "" {
		return nil, errInvalidSensorData
	}

	resp, err := h.commands.SensorCreate.Execute(ctx, sensor_create.Request{
		Name:   req.Sensor.Name,
		Code:   req.Sensor.Code,
		Symbol: req.Sensor.Symbol,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to create sensor: %v", err)
		return nil, err
	}

	retrieveResp, err := h.queries.SensorRetrieve.Execute(ctx, sensor_retrieve.Request{
		SensorID: resp.SensorID,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to retrieve created sensor: %v", err)
		return nil, err
	}

	return &pb_services.SensorCreateReply{
		Sensor: sensorPropsToPb(retrieveResp.Sensor, retrieveResp.MetricTypes),
	}, nil
}

func (h *Handler) SensorGet(
	ctx context.Context,
	req *pb_services.SensorGetRequest,
) (*pb_services.SensorGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.SensorGet.Execute(ctx, sensor_get.Request{})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get sensors: %v", err)
		return nil, err
	}

	if len(resp.Sensors) == 0 {
		return &pb_services.SensorGetReply{
			Sensors: []*pb_models.Sensor_Read{},
		}, nil
	}

	sensors := make([]*pb_models.Sensor_Read, 0, len(resp.Sensors))
	for _, s := range resp.Sensors {
		sensors = append(sensors, sensorPropsToPb(s.Sensor, s.MetricTypes))
	}

	return &pb_services.SensorGetReply{
		Sensors: sensors,
	}, nil
}

func (h *Handler) SensorPatientCreate(
	ctx context.Context,
	req *pb_services.SensorPatientCreateRequest,
) (*pb_services.SensorPatientCreateReply, error) {
	if req == nil {
		return nil, errRequestNil
	}
	if req.SensorId == "" || req.PatientId == "" {
		return nil, errInvalidSensorData
	}

	resp, err := h.commands.SensorPatientCreate.Execute(ctx, patient_create.Request{
		SensorID:  req.SensorId,
		PatientID: req.PatientId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to create sensor patient: %v", err)
		return nil, err
	}

	retrieveResp, err := h.queries.SensorPatientRetrieve.Execute(ctx, patient_retrieve.Request{
		SensorID:  resp.SensorID,
		PatientID: resp.PatientID,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to retrieve created sensor patient: %v", err)
		return nil, err
	}

	return &pb_services.SensorPatientCreateReply{
		SensorPatient: sensorPatientPropsToPb(retrieveResp.SensorPatient),
	}, nil
}

func (h *Handler) SensorPatientGet(
	ctx context.Context,
	req *pb_services.SensorPatientGetRequest,
) (*pb_services.SensorPatientGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}
	if req.SensorId == "" {
		return nil, errInvalidSensorData
	}

	resp, err := h.queries.SensorPatientGet.Execute(ctx, patient_get.Request{
		SensorID: req.SensorId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get sensor patients: %v", err)
		return nil, err
	}

	if len(resp.SensorPatients) == 0 {
		return &pb_services.SensorPatientGetReply{
			SensorPatients: []*pb_models.SensorPatient_Read{},
		}, nil
	}

	sensorPatients := make([]*pb_models.SensorPatient_Read, 0, len(resp.SensorPatients))
	for _, sp := range resp.SensorPatients {
		sensorPatients = append(sensorPatients, sensorPatientPropsToPb(sp))
	}

	return &pb_services.SensorPatientGetReply{
		SensorPatients: sensorPatients,
	}, nil
}

func (h *Handler) SensorPatientRetrieve(
	ctx context.Context,
	req *pb_services.SensorPatientRetrieveRequest,
) (*pb_services.SensorPatientRetrieveReply, error) {
	if req == nil {
		return nil, errRequestNil
	}
	if req.SensorId == "" || req.PatientId == "" {
		return nil, errInvalidSensorData
	}

	resp, err := h.queries.SensorPatientRetrieve.Execute(ctx, patient_retrieve.Request{
		SensorID:  req.SensorId,
		PatientID: req.PatientId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to retrieve sensor patient: %v", err)
		return nil, err
	}

	return &pb_services.SensorPatientRetrieveReply{
		SensorPatient: sensorPatientPropsToPb(resp.SensorPatient),
	}, nil
}

func (h *Handler) SensorPatientMetricGet(
	ctx context.Context,
	req *pb_services.SensorPatientMetricGetRequest,
) (*pb_services.SensorPatientMetricGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}
	if req.SensorId == "" || req.PatientId == "" {
		return nil, errInvalidSensorData
	}

	resp, err := h.queries.SensorPatientMetricGet.Execute(ctx, metric_get.Request{
		SensorID:  req.SensorId,
		PatientID: req.PatientId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get sensor patient metrics: %v", err)
		return nil, err
	}

	if len(resp.Measurements) == 0 {
		return &pb_services.SensorPatientMetricGetReply{
			Measurements: []*pb_models.SensorPatientMetric_Read{},
		}, nil
	}

	measurements := make([]*pb_models.SensorPatientMetric_Read, 0, len(resp.Measurements))
	for _, m := range resp.Measurements {
		measurements = append(measurements, measurementToPb(m))
	}

	return &pb_services.SensorPatientMetricGetReply{
		Measurements: measurements,
	}, nil
}

func (h *Handler) SensorPatientMetricHistoryGet(
	ctx context.Context,
	req *pb_services.SensorPatientMetricHistoryGetRequest,
) (*pb_services.SensorPatientMetricHistoryGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}
	if req.SensorId == "" || req.PatientId == "" {
		return nil, errInvalidSensorData
	}

	start := req.GetStartTime().AsTime()
	end := req.GetEndTime().AsTime()
	if end.IsZero() {
		end = time.Now().UTC()
	}

	resp, err := h.queries.SensorPatientMetricHistoryGet.Execute(ctx, metric_history.Request{
		SensorID:  req.SensorId,
		PatientID: req.PatientId,
		StartTime: start,
		EndTime:   end,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get sensor patient metric history: %v", err)
		return nil, err
	}

	if len(resp.Measurements) == 0 {
		return &pb_services.SensorPatientMetricHistoryGetReply{
			Measurements: []*pb_models.SensorPatientMetric_Read{},
			Total:        int32(resp.Total),
		}, nil
	}

	measurements := make([]*pb_models.SensorPatientMetric_Read, 0, len(resp.Measurements))
	for _, m := range resp.Measurements {
		measurements = append(measurements, measurementToPb(m))
	}

	return &pb_services.SensorPatientMetricHistoryGetReply{
		Measurements: measurements,
		Total:        int32(resp.Total),
	}, nil
}
