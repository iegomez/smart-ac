package storage

import (
	"database/sql"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//Datum holds a message from a device.
type Datum struct {
	ID         int64     `db:"id"`
	DeviceID   int64     `db:"device_id"`
	SensorType string    `db:"sensor_type"`
	Val        float64   `db:"val"`
	StrVal     string    `db:"str_val"`
	CreatedAt  time.Time `db:"created_at"`
}

var sensorTypes = map[string]bool{
	"temperature":     true,
	"carbon_monoxide": true,
	"air_humidity":    true,
	"health_status":   true,
}

// Sensor types constants.
const (
	Temperature    string = "temperature"
	CarbonMonoxide string = "carbon_monoxide"
	AirHumidity    string = "air_humidity"
	HealthStatus   string = "health_status"
)

//DatumWithSerial holds a datum and the associated device's serial number.
type DatumWithSerial struct {
	Datum
	SerialNumber string `db:"serial_number"`
}

//Validate checks that the sensor type is correct, air humidity is a float in the [0.0, 1.0] range, health status' length is less than 150.
//It also zero values unused vals (e.g., if sensor type is health status, then val is set to 0).
func (d *Datum) Validate() error {
	if _, ok := sensorTypes[d.SensorType]; !ok {
		return errors.Errorf("unkwown sensor type %s", d.SensorType)
	}
	if d.SensorType == HealthStatus {
		d.Val = 0.0
	} else {
		d.StrVal = ""
	}
	if d.SensorType == AirHumidity && (d.Val < 0.0 || d.Val > 1.0) {
		return errors.New("air humidity must be in [0.0, 1.0]")
	}
	if d.SensorType == HealthStatus && utf8.RuneCountInString(d.StrVal) >= 150 {
		return errors.New("health status must be shorter than 150 characters")
	}
	return nil
}

//CreateData adds all given data records to the DB.
//Since data may be one message or several, we use a common approach of accepting a data slice and bulk insert them, logging any error and continuing to insert the remaining data.
func CreateData(db *sqlx.DB, data []Datum, deviceID int64) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("datum", "device_id", "sensor_type", "val", "str_val", "created_at"))
	if err != nil {
		return err
	}

	//Requirements say that data for each sensor may be sent on batches of at most 500 value, so we need to keep count of them while preparing the copy and discard those that are past the limit.
	counts := map[string]int{
		Temperature:    0,
		CarbonMonoxide: 0,
		AirHumidity:    0,
		HealthStatus:   0,
	}

	for _, d := range data {
		if err := d.Validate(); err != nil {
			log.Errorf("couldn't insert datum for device %d: %#v\n", deviceID, d)
			continue
		}
		counts[d.SensorType]++
		if counts[d.SensorType] > 500 {
			continue
		}
		_, err = stmt.Exec(deviceID, d.SensorType, d.Val, d.StrVal, d.CreatedAt)
		if err != nil {
			log.Errorf("couldn't insert datum for device %d: %#v\n", deviceID, d)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"count": len(data),
	}).Info("data created")

	return nil

}

//GetDatum retrieves a datum by id.
func GetDatum(db *sqlx.DB, id int64) (Datum, error) {
	var datum Datum
	err := db.Get(&datum, "select * from datum where id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return datum, ErrDoesNotExist
		}
		return datum, errors.Wrap(err, "select error")
	}

	return datum, nil
}

//GetDatumCount returns the count of all data.
func GetDatumCount(db *sqlx.DB, startDate, endDate time.Time, sensors []string) (int64, error) {
	var count int64
	filter := "{" + strings.Join(sensors, ",") + "}"
	err := db.Get(&count, `select count(id) from datum where sensor_type = any($3::text[]) and created_at >= $1 and created_at <= $2`, startDate, endDate, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//ListData retrieves all data filtered by sensor type.
func ListData(db *sqlx.DB, startDate, endDate time.Time, limit, offset int64, sensors []string) ([]DatumWithSerial, error) {
	var data []DatumWithSerial
	filter := "{" + strings.Join(sensors, ",") + "}"
	err := db.Select(&data, `select dat.*, dev.serial_number from datum as dat, device as dev where dat.sensor_type = any($5::text[]) and dat.device_id = dev.id and dat.created_at >= $1 and dat.created_at <= $2 order by dat.created_at desc limit $3 offset $4`, startDate, endDate, limit, offset, filter)
	if err != nil {
		return nil, handlePSQLError(Select, err, "select error")
	}
	return data, nil
}

//GetDatumCountForDevice returns the count of all data for a given device id.
func GetDatumCountForDevice(db *sqlx.DB, deviceID int64, startDate, endDate time.Time, sensors []string) (int64, error) {
	var count int64
	filter := "{" + strings.Join(sensors, ",") + "}"
	err := db.Get(&count, `select count(id) from datum where sensor_type = any($4::text[]) and device_id = $1 and created_at >= $2 and created_at <= $3`, deviceID, startDate, endDate, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//ListDataForDevice retrieves all data for a given device id.
func ListDataForDevice(db *sqlx.DB, deviceID int64, startDate, endDate time.Time, limit, offset int64, sensors []string) ([]Datum, error) {
	var data []Datum
	filter := "{" + strings.Join(sensors, ",") + "}"
	err := db.Select(&data, `select * from datum where sensor_type = any($6::text[]) and device_id = $1 and created_at >= $2 and created_at <= $3 order by created_at desc limit $4 offset $5`, deviceID, startDate, endDate, limit, offset, filter)
	if err != nil {
		return nil, handlePSQLError(Select, err, "select error")
	}
	return data, nil
}

// DeleteDatum deletes the datum matching the given DevEUI.
func DeleteDatum(db *sqlx.DB, id int64) error {
	res, err := db.Exec("delete from datum where id = $1", id)
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
	}).Info("datum deleted")

	return nil
}
