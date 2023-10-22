package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrToHex(t *testing.T) {
	Convey("TestStrToHex", t, func() {
		So(textToHex("hello"), ShouldEqual, "68 65 6c 6c 6f")
		So(textToHex("world"), ShouldEqual, "77 6f 72 6c 64")
		So(textToHex("你好"), ShouldEqual, "4f60 597d")
	})
}

func TestStrToBin(t *testing.T) {
	Convey("TestStrToBin", t, func() {
		So(textToBin("hello"), ShouldEqual, "01101000 01100101 01101100 01101100 01101111")
		So(textToBin("world"), ShouldEqual, "01110111 01101111 01110010 01101100 01100100")
		So(textToBin("你好"), ShouldEqual, "0100111101100000 0101100101111101")
	})
}

func TestStrToDec(t *testing.T) {
	Convey("TestStrToDec", t, func() {
		So(textToDec("hello"), ShouldEqual, "104 101 108 108 111")
		So(textToDec("world"), ShouldEqual, "119 111 114 108 100")
		So(textToDec("你好"), ShouldEqual, "20320 22909")
	})
}

func TestConvertTextBinaryConversion(t *testing.T) {
	app := NewTextBinaryConversionApp()

	Convey("TestConvertTextBinaryConversion", t, func() {
		Convey("bin", func() {
			req := &TextBinaryConversionReq{
				In:     "hello",
				ToType: "bin",
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.To, ShouldEqual, "01101000 01100101 01101100 01101100 01101111")
		})

		Convey("hex", func() {
			req := &TextBinaryConversionReq{
				In:     "hello",
				ToType: "hex",
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.To, ShouldEqual, "68 65 6C 6C 6F")
		})

		Convey("hex lower case", func() {
			req := &TextBinaryConversionReq{
				In:        "hello",
				ToType:    "hex",
				LowerCase: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.To, ShouldEqual, "68 65 6c 6c 6f")
		})
	})
}
