package storage

import (
	"database/sql"
	"time"
	"unicode/utf8"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// TODO
/*
	create table datum (
  id bigserial primary key,
  device_id bigint references datum on delete cascade,
	temperature double precision not null default 0.0,
	carbon_monoxide double precision not null default 0.0,
  air_humidity double precision not null default 0.0,
  health_status text not null default ''
);
*/

//Datum holds a message from a device.
type Datum struct {
	ID             int64     `db:"id"`
	DeviceID       int64     `db:"device_id"`
	Temperature    float64   `db:"temperature"`
	CarbonMonoxide float64   `db:"carbon_monoxide"`
	AirHumidity    float64   `db:"air_humidity"`
	HealthStatus   string    `db:"health_status"`
	CreatedAt      time.Time `db:"created_at"`
}

//Validate checks that a datum is associated to a datum, air humidity is a float in the [0.0, 1.0] range and the health status length is less than 150.
func (d Datum) Validate() error {
	if d.DeviceID == 0 {
		return errors.New("datum must be associated to a datum")
	}
	if d.AirHumidity < 0.0 || d.AirHumidity > 1.0 {
		return errors.New("air humidity must be in [0.0, 1.0]")
	}
	if utf8.RuneCountInString(d.HealthStatus) >= 150 {
		return errors.New("health status must be shorter than 150 characters")
	}
	return nil
}

//CreateData adds all given data records to the DB.
//Since data may be one message or several, we use a common approach of accepting a data slice and bulk insert them, logging any error and continuing to insert the remaining data.
func CreateData(db *sqlx.DB, data []Datum) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("device_id", "temperature", "carbon_monoxide", "air_humidity", "health_status", "created_at"))
	if err != nil {
		return err
	}

	for _, d := range data {
		if err := d.Validate(); err != nil {
			log.Errorf("couldn't insert datum for device %d: %#v\n", d)
			continue
		}
		_, err = stmt.Exec(d.DeviceID, d.Temperature, d.CarbonMonoxide, d.AirHumidity, d.HealthStatus, time.Now())
		if err != nil {
			log.Errorf("couldn't insert datum for device %d: %#v\n", d)
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
func GetDatumCount(db *sqlx.DB) (int64, error) {
	var count int64
	err := db.Get(&count, `select count(id) from datum`)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//ListData retrieves all data.
func ListData(db *sqlx.DB, limit, offset int64) ([]Datum, error) {
	var data []Datum
	err := db.Select(&data, `select * from datum order by created_at desc limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, handlePSQLError(Select, err, "select error")
	}
	return data, nil
}

//GetDatumCountForDevice returns the count of all data for a given device id.
func GetDatumCountForDevice(db *sqlx.DB, deviceID int64) (int64, error) {
	var count int64
	err := db.Get(&count, `select count(id) from datum where device_id = $1`, deviceID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//ListDataForDevice retrieves all data for a given device id.
func ListDataForDevice(db *sqlx.DB, deviceID, limit, offset int64) ([]Datum, error) {
	var data []Datum
	err := db.Select(&data, `select * from datum where device_id = $1 order by created_at desc limit $2 offset $3`, deviceID, limit, offset)
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
