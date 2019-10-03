package storage

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/iegomez/smart-ac/internal/test"
)

func TestDevice(t *testing.T) {
	conf := test.GetConfig()
	if err := Setup(conf); err != nil {
		t.Fatal(err)
	}

	Convey("Given a clean database", t, func() {
		test.MustResetDB(DB())

		key, err := GenerateRandomKey()
		So(err, ShouldBeNil)
		now := time.Now()

		Convey("Creating a device without serial number or firmware version should give an error", func() {

			device := Device{
				RegisteredAt:    now,
				FirmwareVersion: "0.1.0",
				APIKey:          key,
			}

			So(CreateDevice(DB(), &device), ShouldNotBeNil)

			device.SerialNumber = "0001"
			device.FirmwareVersion = ""

			So(CreateDevice(DB(), &device), ShouldNotBeNil)

			Convey("Giving valid values should work", func() {
				device.FirmwareVersion = "0.1.0"
				So(CreateDevice(DB(), &device), ShouldBeNil)

				Convey("Listing them should give 1", func() {
					count, err := GetDeviceCount(DB())
					So(count, ShouldEqual, 1)
					devices, err := ListDevices(DB(), 10, 0)
					So(err, ShouldBeNil)
					So(len(devices), ShouldEqual, 1)
					So(devices[0].ID, ShouldEqual, device.ID)

					Convey("Updating it should work", func() {
						updateDevice := device
						updateDevice.FirmwareVersion = "0.2.0"
						So(UpdateDevice(DB(), &updateDevice), ShouldBeNil)
						So(device.FirmwareVersion, ShouldEqual, device.FirmwareVersion)
						Convey("And getting it bys erial number should render the same device", func() {
							getDevice, err := GetDeviceBySerialNumber(DB(), device.SerialNumber)
							So(err, ShouldBeNil)
							So(getDevice.ID, ShouldEqual, device.ID)

							Convey("And deleting it should leave us with 0 devices", func() {

								So(DeleteDevice(DB(), device.ID), ShouldBeNil)
								count, err := GetDeviceCount(DB())
								So(err, ShouldBeNil)
								So(count, ShouldEqual, 0)
							})

						})
					})

				})

			})

		})

	})
}
