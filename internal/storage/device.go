package storage

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//Device holds the representation of a device.
type Device struct {
	ID int64 `db:"id"`
	//UserID          int64     `db:"user_id"`  // We'll keep this out for now.
	SerialNumber    string    `db:"serial_number"`
	RegisteredAt    time.Time `db:"registered_at"`
	FirmwareVersion string    `db:"firmware_version"`
	APIKey          string    `db:"api_key"`
}

// Validate validates the device data.
func (d Device) Validate() error {
	if d.SerialNumber == "" {
		return errors.New("device must have a serial number")
	}
	if d.FirmwareVersion == "" {
		return errors.New("device must have a firmware version")
	}
	return nil
}

//CreateDevice creates a new device.
func CreateDevice(db *sqlx.DB, d *Device) error {

	if err := d.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	//Generate an api key.
	key, err := GenerateRandomKey()
	if err != nil {
		return errors.Wrap(err, "generate random key error")
	}

	err = db.Get(&d.ID,
		`insert into device(serial_number, registered_at, firmware_version, api_key)
		values($1, $2, $3, $4) returning id`,
		d.SerialNumber,
		time.Now(),
		d.FirmwareVersion,
		key,
	)

	if err != nil {
		return handlePSQLError(Insert, err, "insert error")
	}

	log.WithFields(log.Fields{
		"id":            d.ID,
		"serial number": d.SerialNumber,
	}).Info("device created")

	return nil

}

//GetDevice retrieves a device by id.
func GetDevice(db *sqlx.DB, id int64) (Device, error) {
	var device Device
	err := db.Get(&device, "select * from device where id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return device, ErrDoesNotExist
		}
		return device, errors.Wrap(err, "select error")
	}

	return device, nil
}

//GetDeviceBySerialNumber retrieves a device by serial number.
func GetDeviceBySerialNumber(db *sqlx.DB, serialNumber string) (Device, error) {
	var device Device
	err := db.Get(&device, "select * from device where serial_number = $1", serialNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return device, ErrDoesNotExist
		}
		return device, errors.Wrap(err, "select error")
	}

	return device, nil
}

//GetDeviceCount returns the count of all devices.
func GetDeviceCount(db *sqlx.DB) (int64, error) {
	var count int64
	err := db.Get(&count, `select count(id) from device`)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//ListDevices retrieves all devices.
func ListDevices(db *sqlx.DB, limit, offset int64) ([]Device, error) {
	var devices []Device
	err := db.Select(&devices, `select * from device order by registered_at desc limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, handlePSQLError(Select, err, "select error")
	}
	return devices, nil
}

// UpdateDevice updates the given device.
func UpdateDevice(db *sqlx.DB, d *Device) error {

	if err := d.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	res, err := db.Exec(`
      update device
        set serial_number = $2, firmware_version = $3
        where id = $1`,
		d.ID,
		d.SerialNumber,
		d.FirmwareVersion,
	)
	if err != nil {
		return handlePSQLError(Update, err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"id": d.ID,
	}).Info("device updated")

	return nil
}

// UpdateDeviceKey updates the given device's api key.
func UpdateDeviceKey(db *sqlx.DB, id int64) (string, error) {

	//Generate an api key.
	key, err := GenerateRandomKey()
	if err != nil {
		return "", errors.Wrap(err, "generate random key error")
	}

	res, err := db.Exec(`
      update device
        set api_key = $2
        where id = $1`,
		id,
		key,
	)
	if err != nil {
		return "", handlePSQLError(Update, err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return "", errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return "", ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info("device updated")

	return key, nil
}

// DeleteDevice deletes the device matching the given DevEUI.
func DeleteDevice(db *sqlx.DB, id int64) error {
	res, err := db.Exec("delete from device where id = $1", id)
	if err != nil {
		return handlePSQLError(Delete, err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info("device deleted")

	return nil
}

// GenerateRandomKey returns a base64 encoded securely generated random string.
func GenerateRandomKey() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
