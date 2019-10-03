package storage

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/iegomez/smart-ac/internal/test"
)

func TestDatum(t *testing.T) {
	conf := test.GetConfig()
	if err := Setup(conf); err != nil {
		t.Fatal(err)
	}

	Convey("Given a clean database", t, func() {
		test.MustResetDB(DB())

		key, err := GenerateRandomKey()
		So(err, ShouldBeNil)
		now := time.Now()

		device := Device{
			SerialNumber:    "0001",
			RegisteredAt:    now,
			FirmwareVersion: "0.1.0",
			APIKey:          key,
		}

		So(CreateDevice(DB(), &device), ShouldBeNil)

		Convey("When creating wrong and correct data", func() {

			wrongDatum1 := Datum{
				DeviceID:   device.ID,
				SensorType: AirHumidity,
				Val:        2.0,
				CreatedAt:  now.Add(-5 * time.Minute),
			}

			tooLongStr := "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"

			wrongDatum2 := Datum{
				DeviceID:   device.ID,
				SensorType: HealthStatus,
				StrVal:     tooLongStr,
				CreatedAt:  now.Add(-4 * time.Minute),
			}

			correctDatum1 := Datum{
				DeviceID:   device.ID,
				SensorType: AirHumidity,
				Val:        0.7,
				CreatedAt:  now.Add(-3 * time.Minute),
			}

			correctDatum2 := Datum{
				DeviceID:   device.ID,
				SensorType: Temperature,
				Val:        -5.3,
				CreatedAt:  now.Add(-2 * time.Minute),
			}

			data := []Datum{wrongDatum1, wrongDatum2, correctDatum1, correctDatum2}

			err = CreateData(DB(), data, device.ID)

			Convey("Then only the correct ones should be stored but the creation shouldn't error", func() {
				So(err, ShouldBeNil)
				count, err := GetDatumCount(DB(), now.Add(-10*time.Minute), now.Add(10*time.Minute), []string{"temperature", "air_humidity", "carbon_moxide", "health_status"})
				So(err, ShouldBeNil)
				So(count, ShouldEqual, 2)
				getData, err := ListData(DB(), now.Add(-10*time.Minute), now.Add(10*time.Minute), 10, 0, []string{"temperature", "air_humidity", "carbon_moxide", "health_status"})
				So(err, ShouldBeNil)
				So(len(getData), ShouldEqual, 2)

				Convey("Then listing by device should give the same result", func() {
					count, err := GetDatumCountForDevice(DB(), device.ID, now.Add(-10*time.Minute), now.Add(10*time.Minute), []string{"temperature", "air_humidity", "carbon_moxide", "health_status"})
					So(err, ShouldBeNil)
					So(count, ShouldEqual, 2)
					deviceData, err := ListDataForDevice(DB(), device.ID, now.Add(-10*time.Minute), now.Add(10*time.Minute), 10, 0, []string{"temperature", "air_humidity", "carbon_moxide", "health_status"})
					So(err, ShouldBeNil)
					So(len(deviceData), ShouldEqual, 2)
					So(deviceData[0].ID, ShouldEqual, getData[0].ID)
					So(deviceData[1].ID, ShouldResemble, getData[1].ID)

					Convey("Filtering by one type should give only those records", func() {
						getData, err := ListData(DB(), now.Add(-10*time.Minute), now.Add(10*time.Minute), 10, 0, []string{"temperature"})
						So(err, ShouldBeNil)
						So(len(getData), ShouldEqual, 1)
						deviceData, err := ListDataForDevice(DB(), device.ID, now.Add(-10*time.Minute), now.Add(10*time.Minute), 10, 0, []string{"temperature"})
						So(err, ShouldBeNil)
						So(len(deviceData), ShouldEqual, 1)
					})

				})

			})

		})

	})
}
