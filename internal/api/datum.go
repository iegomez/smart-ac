package api

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	pb "github.com/iegomez/smart-ac/api"
	"github.com/iegomez/smart-ac/internal/api/helpers"

	"github.com/iegomez/smart-ac/internal/storage"
)

// DatumAPI exports the Node related functions.
type DatumAPI struct {
}

// NewDatumAPI creates a new NodeAPI.
func NewDatumAPI() *DatumAPI {
	return &DatumAPI{}
}

// Create creates the given datum.
func (a *DatumAPI) Create(ctx context.Context, req *pb.CreateDataRequest) (*empty.Empty, error) {

	data := make([]storage.Datum, len(req.Data))
	for i, d := range req.Data {
		data[i] = storage.Datum{
			DeviceID:       d.DeviceId,
			Temperature:    d.Temperature,
			CarbonMonoxide: d.CarbonMonoxide,
			AirHumidity:    d.AirHumidity,
			HealthStatus:   d.HealthStatus,
		}
	}

	err := storage.CreateData(storage.DB(), data)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}

// Get retrieves a datum given an id.
func (a *DatumAPI) Get(ctx context.Context, req *pb.DatumRequest) (*pb.GetDatumResponse, error) {

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

// List retrieves all data given a limit and an offset.
func (a *DatumAPI) List(ctx context.Context, req *pb.ListDataRequest) (*pb.ListDataResponse, error) {

	data, err := storage.ListData(storage.DB(), req.Limit, req.Offset)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	count, err := storage.GetDatumCount(storage.DB())
	if err != nil {
		return nil, err
	}

	resp := &pb.ListDataResponse{
		TotalCount: count,
		Data:       make([]*pb.GetDatumResponse, len(data)),
	}

	for i, d := range data {
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
		resp.Data[i] = datum
	}

	return resp, nil

}

//Delete deletes a datum given an id.
func (a *DatumAPI) Delete(ctx context.Context, req *pb.DatumRequest) (*empty.Empty, error) {
	err := storage.DeleteDatum(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}
	return &empty.Empty{}, nil
}
