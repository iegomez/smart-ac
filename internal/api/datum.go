package api

import (
	"regexp"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	pb "github.com/iegomez/smart-ac/api"
	"github.com/iegomez/smart-ac/internal/api/helpers"

	"github.com/iegomez/smart-ac/internal/storage"
	log "github.com/sirupsen/logrus"
)

// DatumAPI exports the Node related functions.
type DatumAPI struct {
}

// NewDatumAPI creates a new NodeAPI.
func NewDatumAPI() *DatumAPI {
	return &DatumAPI{}
}

var validApiKeyAuthorizationRegexp = regexp.MustCompile(`(?i)^Apikey (.*)$`)

//getKeyFromContext tries to get the api key sent in the Authorization header when receiving data from a device.
func getKeyFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata in context")
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", errors.New("no authorization-data in metadata")
	}

	match := validApiKeyAuthorizationRegexp.FindStringSubmatch(token[0])

	// authorization header should respect RFC1945
	if len(match) == 0 {
		log.Warning("RFC1945 format expected : Authorization: <type> <credentials>")
		return token[0], nil
	}

	return match[1], nil
}

// Create creates the given datum.
func (a *DatumAPI) Create(ctx context.Context, req *pb.CreateDataRequest) (*empty.Empty, error) {

	//We first need to check for an api key.
	key, err := getKeyFromContext(ctx)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	//Now we need to get the device by it's serial number.
	device, err := storage.GetDeviceBySerialNumber(storage.DB(), req.SerialNumber)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	//Now check that the given key is correct.
	if device.APIKey != key {
		return nil, helpers.ErrToRPCError(errors.New("wrong api key"))
	}

	data := make([]storage.Datum, len(req.Data))
	for i, d := range req.Data {
		data[i] = storage.Datum{
			DeviceID:       device.ID,
			Temperature:    d.Temperature,
			CarbonMonoxide: d.CarbonMonoxide,
			AirHumidity:    d.AirHumidity,
			HealthStatus:   d.HealthStatus,
		}
	}

	err = storage.CreateData(storage.DB(), data, device.ID)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}

// Get retrieves a datum given an id.
func (a *DatumAPI) Get(ctx context.Context, req *pb.DatumRequest) (*pb.GetDatumResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	d, err := storage.GetDatum(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	datum := &pb.GetDatumResponse{
		Id:             d.ID,
		DeviceId:       d.DeviceID,
		Temperature:    d.Temperature,
		CarbonMonoxide: d.CarbonMonoxide,
		AirHumidity:    d.AirHumidity,
		HealthStatus:   d.HealthStatus,
	}

	datum.CreatedAt, err = ptypes.TimestampProto(d.CreatedAt)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return datum, nil

}

// List retrieves all data between a start and an end date, given a limit and an offset.
func (a *DatumAPI) List(ctx context.Context, req *pb.ListDataRequest) (*pb.ListDataResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	startDate, err := ptypes.Timestamp(req.StartDate)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	endDate, err := ptypes.Timestamp(req.EndDate)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	data, err := storage.ListData(storage.DB(), startDate, endDate, req.Limit, req.Offset)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	count, err := storage.GetDatumCount(storage.DB(), startDate, endDate)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListDataResponse{
		TotalCount: count,
		Result:     make([]*pb.GetDatumResponse, len(data)),
	}

	for i, d := range data {
		resp.Result[i] = &pb.GetDatumResponse{
			Id:             d.ID,
			DeviceId:       d.DeviceID,
			Temperature:    d.Temperature,
			CarbonMonoxide: d.CarbonMonoxide,
			AirHumidity:    d.AirHumidity,
			HealthStatus:   d.HealthStatus,
			SerialNumber:   d.SerialNumber,
		}
		resp.Result[i].CreatedAt, err = ptypes.TimestampProto(d.CreatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
	}

	return resp, nil
}

// ListForDevice retrieves all data for a given device between a start and an end date, given a limit and an offset.
func (a *DatumAPI) ListForDevice(ctx context.Context, req *pb.ListDataForDeviceRequest) (*pb.ListDataResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	startDate, err := ptypes.Timestamp(req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := ptypes.Timestamp(req.EndDate)
	if err != nil {
		return nil, err
	}

	data, err := storage.ListDataForDevice(storage.DB(), req.DeviceId, startDate, endDate, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	count, err := storage.GetDatumCountForDevice(storage.DB(), req.DeviceId, startDate, endDate)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListDataResponse{
		TotalCount: count,
		Result:     make([]*pb.GetDatumResponse, len(data)),
	}

	for i, d := range data {
		resp.Result[i] = &pb.GetDatumResponse{
			Id:             d.ID,
			DeviceId:       d.DeviceID,
			Temperature:    d.Temperature,
			CarbonMonoxide: d.CarbonMonoxide,
			AirHumidity:    d.AirHumidity,
			HealthStatus:   d.HealthStatus,
		}
		resp.Result[i].CreatedAt, err = ptypes.TimestampProto(d.CreatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
	}

	return resp, nil

}

//Delete deletes a datum given an id.
func (a *DatumAPI) Delete(ctx context.Context, req *pb.DatumRequest) (*empty.Empty, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	err = storage.DeleteDatum(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}
	return &empty.Empty{}, nil
}
