package api

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	pb "github.com/iegomez/smart-ac/api"
	"github.com/iegomez/smart-ac/internal/api/helpers"

	"github.com/iegomez/smart-ac/internal/storage"
)

// DeviceAPI exports the Node related functions.
type DeviceAPI struct {
}

// NewDeviceAPI creates a new NodeAPI.
func NewDeviceAPI() *DeviceAPI {
	return &DeviceAPI{}
}

// Create creates the given device.
func (a *DeviceAPI) Create(ctx context.Context, req *pb.CreateDeviceRequest) (*empty.Empty, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	d := storage.Device{
		SerialNumber:    req.SerialNumber,
		FirmwareVersion: req.FirmwareVersion,
	}

	err = storage.CreateDevice(storage.DB(), &d)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return &empty.Empty{}, nil
}

// Get retrieves a device given an id.
func (a *DeviceAPI) Get(ctx context.Context, req *pb.DeviceRequest) (*pb.GetDeviceResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	d, err := storage.GetDevice(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	device := &pb.GetDeviceResponse{
		Id:              d.ID,
		SerialNumber:    d.SerialNumber,
		FirmwareVersion: d.FirmwareVersion,
	}

	device.RegisteredAt, err = ptypes.TimestampProto(d.RegisteredAt)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return device, nil

}

// GetAPIKey retrieves a device's API key'.
func (a *DeviceAPI) GetAPIKey(ctx context.Context, req *pb.DeviceRequest) (*pb.GetDeviceKeyResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	d, err := storage.GetDevice(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	key := &pb.GetDeviceKeyResponse{
		ApiKey: d.APIKey,
	}

	return key, nil

}

// GetBySerialNumber retrieves a device given a serial number.
func (a *DeviceAPI) GetBySerialNumber(ctx context.Context, req *pb.DeviceBySerialNumberRequest) (*pb.GetDeviceResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	d, err := storage.GetDeviceBySerialNumber(storage.DB(), req.SerialNumber)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	device := &pb.GetDeviceResponse{
		Id:              d.ID,
		SerialNumber:    d.SerialNumber,
		FirmwareVersion: d.FirmwareVersion,
	}

	device.RegisteredAt, err = ptypes.TimestampProto(d.RegisteredAt)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	return device, nil

}

// List retrieves all devices given a limit and an offset.
func (a *DeviceAPI) List(ctx context.Context, req *pb.ListDeviceRequest) (*pb.ListDeviceResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	devices, err := storage.ListDevices(storage.DB(), req.Limit, req.Offset)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	count, err := storage.GetDeviceCount(storage.DB())
	if err != nil {
		return nil, err
	}

	resp := &pb.ListDeviceResponse{
		TotalCount: count,
		Result:     make([]*pb.GetDeviceResponse, len(devices)),
	}

	for i, d := range devices {
		device := &pb.GetDeviceResponse{
			Id:              d.ID,
			SerialNumber:    d.SerialNumber,
			FirmwareVersion: d.FirmwareVersion,
		}
		device.RegisteredAt, err = ptypes.TimestampProto(d.RegisteredAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		resp.Result[i] = device
	}

	return resp, nil

}

// ListAll retrieves all devices.
func (a *DeviceAPI) ListAll(ctx context.Context, req *empty.Empty) (*pb.ListDeviceResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	count, err := storage.GetDeviceCount(storage.DB())
	if err != nil {
		return nil, err
	}

	devices, err := storage.ListDevices(storage.DB(), count, 0)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	resp := &pb.ListDeviceResponse{
		TotalCount: count,
		Result:     make([]*pb.GetDeviceResponse, len(devices)),
	}

	for i, d := range devices {
		device := &pb.GetDeviceResponse{
			Id:              d.ID,
			SerialNumber:    d.SerialNumber,
			FirmwareVersion: d.FirmwareVersion,
		}
		device.RegisteredAt, err = ptypes.TimestampProto(d.RegisteredAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		resp.Result[i] = device
	}

	return resp, nil

}

//Update updates the given device.
func (a *DeviceAPI) Update(ctx context.Context, req *pb.UpdateDeviceRequest) (*empty.Empty, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	device := &storage.Device{
		ID:              req.Id,
		SerialNumber:    req.SerialNumber,
		FirmwareVersion: req.FirmwareVersion,
	}
	err = storage.UpdateDevice(storage.DB(), device)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}
	return &empty.Empty{}, nil
}

//UpdateAPIKey updates the given device.
func (a *DeviceAPI) UpdateAPIKey(ctx context.Context, req *pb.DeviceRequest) (*pb.GetDeviceKeyResponse, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	key, err := storage.UpdateDeviceKey(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}
	return &pb.GetDeviceKeyResponse{
		ApiKey: key,
	}, nil
}

//Delete deletes a device given an id.
func (a *DeviceAPI) Delete(ctx context.Context, req *pb.DeviceRequest) (*empty.Empty, error) {

	isUser, err := GetIsUser(ctx)
	if err != nil || !isUser {
		return nil, helpers.ErrToRPCError(err)
	}

	err = storage.DeleteDevice(storage.DB(), req.Id)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}
	return &empty.Empty{}, nil
}
