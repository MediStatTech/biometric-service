package diseas

import (
	"context"

	disease_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/create"
	disease_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/get"
	disease_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/retrieve"
	disease_sensor_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/create"
	disease_sensor_retrieve "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/retrieve"
	disease_sensor_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/diseases/sensor/get"
	pb_services "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
)

func (h *Handler) DiseasCreate(
	ctx context.Context,
	req *pb_services.DiseasCreateRequest,
) (*pb_services.DiseasCreateReply, error) {
	if req == nil || req.Diseas == nil {
		return nil, errRequestNil
	}

	diseasData := req.Diseas
	if diseasData.Name == "" || diseasData.Code == "" {
		return nil, errInvalidDiseaseData
	}

	resp, err := h.commands.DiseaseCreate.Execute(ctx, disease_create.Request{
		Name: diseasData.Name,
		Code: diseasData.Code,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to create disease: %v", err)
		return nil, err
	}

	retrieveResp, err := h.queries.DiseaseRetrieve.Execute(ctx, disease_retrieve.Request{
		DiseaseID: resp.DiseaseID,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to retrieve created disease: %v", err)
		return nil, err
	}

	return &pb_services.DiseasCreateReply{
		Diseas: diseasePropsToPb(retrieveResp.Disease),
	}, nil
}

func (h *Handler) DiseasGet(
	ctx context.Context,
	req *pb_services.DiseasGetRequest,
) (*pb_services.DiseasGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.DiseaseGet.Execute(ctx, disease_get.Request{})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get diseases: %v", err)
		return nil, err
	}

	if len(resp.Diseases) == 0 {
		return &pb_services.DiseasGetReply{
			Diseases: []*pb_models.Diseas_Read{},
		}, nil
	}

	diseases := make([]*pb_models.Diseas_Read, 0, len(resp.Diseases))
	for _, disease := range resp.Diseases {
		diseases = append(diseases, diseasePropsToPb(disease))
	}

	return &pb_services.DiseasGetReply{
		Diseases: diseases,
	}, nil
}

func (h *Handler) DiseasSensorCreate(
	ctx context.Context,
	req *pb_services.DiseasSensorCreateRequest,
) (*pb_services.DiseasSensorCreateReply, error) {
	if req == nil || req.DiseasSensor == nil {
		return nil, errRequestNil
	}

	diseasSensorData := req.DiseasSensor
	if diseasSensorData.DiseasId == "" || diseasSensorData.SensorId == "" {
		return nil, errInvalidDiseaseData
	}

	resp, err := h.commands.DiseaseSensorCreate.Execute(ctx, disease_sensor_create.Request{
		DiseaseID: diseasSensorData.DiseasId,
		SensorID:  diseasSensorData.SensorId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to create disease sensor: %v", err)
		return nil, err
	}

	retrieveResp, err := h.queries.DiseaseSensorRetrieve.Execute(ctx, disease_sensor_retrieve.Request{
		DiseaseID: resp.DiseaseID,
		SensorID:  resp.SensorID,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to retrieve created disease sensor: %v", err)
		return nil, err
	}

	return &pb_services.DiseasSensorCreateReply{
		DiseasSensor: diseaseSensorPropsToPb(retrieveResp.DiseaseSensor),
	}, nil
}

func (h *Handler) DiseasSensorGet(
	ctx context.Context,
	req *pb_services.DiseasSensorGetRequest,
) (*pb_services.DiseasSensorGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.DiseaseSensorGet.Execute(ctx, disease_sensor_get.Request{
		DiseaseID: req.DiseasId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get disease sensors: %v", err)
		return nil, err
	}

	if len(resp.DiseaseSensors) == 0 {
		return &pb_services.DiseasSensorGetReply{
			DiseasSensors: []*pb_models.DiseasSensor_Read{},
		}, nil
	}

	diseaseSensors := make([]*pb_models.DiseasSensor_Read, 0, len(resp.DiseaseSensors))
	for _, diseaseSensor := range resp.DiseaseSensors {
		diseaseSensors = append(diseaseSensors, diseaseSensorPropsToPb(diseaseSensor))
	}

	return &pb_services.DiseasSensorGetReply{
		DiseasSensors: diseaseSensors,
	}, nil
}
