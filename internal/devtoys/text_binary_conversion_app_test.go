package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrToHex(t *testing.T) {
	Convey("TestStrToHex", t, func() {
		So(strToHex("hello"), ShouldEqual, "68 65 6c 6c 6f")
		So(strToHex("world"), ShouldEqual, "77 6f 72 6c 64")
		So(strToHex("你好"), ShouldEqual, "4f60 597d")
	})
}

func TestStrToBase64(t *testing.T) {
	Convey("TestStrToBase64", t, func() {
		So(strToBase64("hello"), ShouldEqual, "aGVsbG8=")
		So(strToBase64("world"), ShouldEqual, "d29ybGQ=")
		So(strToBase64("你好"), ShouldEqual, "5L2g5aW9")
	})
}

func TestStrToBin(t *testing.T) {
	Convey("TestStrToBin", t, func() {
		So(strToBin("hello"), ShouldEqual, "01101000 01100101 01101100 01101100 01101111")
		So(strToBin("world"), ShouldEqual, "01110111 01101111 01110010 01101100 01100100")
		So(strToBin("你好"), ShouldEqual, "0100111101100000 0101100101111101")
	})
}

func TestStrToDec(t *testing.T) {
	Convey("TestStrToDec", t, func() {
		So(strToDec("hello"), ShouldEqual, "104 101 108 108 111")
		So(strToDec("world"), ShouldEqual, "119 111 114 108 100")
		So(strToDec("你好"), ShouldEqual, "20320 22909")
	})
}

func TestConvertTextBinaryConversion(t *testing.T) {
	app := NewTextBinaryConversionApp()

	Convey("TestConvertTextBinaryConversion", t, func() {
		Convey("bin", func() {
			req := &TextBinaryConversionReq{
				Text: "hello",
				To:   "bin",
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "01101000 01100101 01101100 01101100 01101111")
		})

		Convey("bin without space", func() {
			req := &TextBinaryConversionReq{
				Text:         "hello",
				To:           "bin",
				WithoutSpace: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "0110100001100101011011000110110001101111")
		})

		Convey("bin without fill zero", func() {
			req := &TextBinaryConversionReq{
				Text:            "hello",
				To:              "bin",
				WithoutFillZero: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "1101000 1100101 1101100 1101100 1101111")
		})

		Convey("bin without space and fill zero", func() {
			req := &TextBinaryConversionReq{
				Text:            "hello",
				To:              "bin",
				WithoutSpace:    true,
				WithoutFillZero: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "11010001100101110110011011001101111")
		})

		Convey("hex", func() {
			req := &TextBinaryConversionReq{
				Text: "hello",
				To:   "hex",
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "68 65 6C 6C 6F")
		})

		Convey("hex without space", func() {
			req := &TextBinaryConversionReq{
				Text:         "hello",
				To:           "hex",
				WithoutSpace: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "68656C6C6F")
		})

		Convey("hex without fill zero", func() {
			req := &TextBinaryConversionReq{
				Text:            "hello",
				To:              "hex",
				WithoutFillZero: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "68 65 6C 6C 6F")
		})

		Convey("hex without space and fill zero", func() {
			req := &TextBinaryConversionReq{
				Text:            "hello",
				To:              "hex",
				WithoutSpace:    true,
				WithoutFillZero: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "68656C6C6F")
		})

		Convey("hex lower case", func() {
			req := &TextBinaryConversionReq{
				Text:      "hello",
				To:        "hex",
				LowerCase: true,
			}
			res, err := app.ConvertTextBinary(req)
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, "68 65 6c 6c 6f")
		})
	})
}
