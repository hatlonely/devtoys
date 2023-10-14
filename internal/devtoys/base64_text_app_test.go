package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBase64TextApp_Base64Encode(t *testing.T) {
	app := NewBase64TextApp()
	Convey("TestBase64TextApp_Base64Encode", t, func() {
		Convey("test std", func() {
			val, err := app.Base64Encode("hello world", "std", true)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "aGVsbG8gd29ybGQ=")
		})

		Convey("test url", func() {
			val, err := app.Base64Encode("hello world", "url", true)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "aGVsbG8gd29ybGQ=")
		})

		Convey("test raw std", func() {
			val, err := app.Base64Encode("hello world", "std", false)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "aGVsbG8gd29ybGQ")
		})

		Convey("test raw url", func() {
			val, err := app.Base64Encode("hello world", "url", false)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "aGVsbG8gd29ybGQ")
		})
	})
}

func TestBase64TextApp_Base64Decode(t *testing.T) {
	app := NewBase64TextApp()
	Convey("TestBase64TextApp_Base64Decode", t, func() {
		Convey("test std", func() {
			val, err := app.Base64Decode("aGVsbG8gd29ybGQ=", "std", true)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "hello world")
		})

		Convey("test url", func() {
			val, err := app.Base64Decode("aGVsbG8gd29ybGQ=", "url", true)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "hello world")
		})

		Convey("test raw std", func() {
			val, err := app.Base64Decode("aGVsbG8gd29ybGQ", "std", false)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "hello world")
		})

		Convey("test raw url", func() {
			val, err := app.Base64Decode("aGVsbG8gd29ybGQ", "url", false)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "hello world")
		})
	})
}
