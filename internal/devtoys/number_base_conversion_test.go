package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumberBaseConversionApp(t *testing.T) {
	app := NewNumberBaseConversionApp()

	Convey("Given a NumberBaseConversionApp", t, func() {
		Convey("When converting a number from one base to another", func() {
			req := &NumberBaseConversionReq{
				Number: "1010",
				InBase: 2,
				ToBase: 16,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "A")
		})

		Convey("When converting a number from decimal to hex 1", func() {
			req := &NumberBaseConversionReq{
				Number: "255",
				InBase: 10,
				ToBase: 16,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "FF")
		})

		Convey("When converting a number from decimal to hex 2", func() {
			req := &NumberBaseConversionReq{
				Number: "4096",
				InBase: 10,
				ToBase: 16,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "1000")
		})

		Convey("When converting a number from decimal to hex 3", func() {
			req := &NumberBaseConversionReq{
				Number: "65535",
				InBase: 10,
				ToBase: 16,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "FFFF")
		})

		Convey("When converting a number from hexadecimal to octal 1", func() {
			req := &NumberBaseConversionReq{
				Number: "FF",
				InBase: 16,
				ToBase: 8,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "377")
		})

		Convey("When converting a number from hexadecimal to octal 2", func() {
			req := &NumberBaseConversionReq{
				Number: "ABCDEF",
				InBase: 16,
				ToBase: 8,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "52746757")
		})

		Convey("When converting big number from hexadecimal to decimal", func() {
			req := &NumberBaseConversionReq{
				Number: "ABCDEFABCDEFABCDEFABCDEFABCDEFABCDEF",
				InBase: 16,
				ToBase: 10,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "14966277357100432636419616004569301003259375")
		})

		Convey("When converting big number from decimal to hexadecimal", func() {
			req := &NumberBaseConversionReq{
				Number: "14966277357100432636419616004569301003259375",
				InBase: 10,
				ToBase: 16,
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "ABCDEFABCDEFABCDEFABCDEFABCDEFABCDEF")
		})

		Convey("When converting use custom base characters", func() {
			req := &NumberBaseConversionReq{
				Number:           "14966277357100432636419616004569301003259375",
				InBase:           10,
				ToBaseCharacters: "🌵🐱🚗🌍🌞🍀💎💰🐼🍋🐵🌴💜🐭🌝",
			}
			res, err := app.ConvertNumberBase(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Number, ShouldEqual, "💎💜🌴🌴🐱🍀🐼🌝🍋🌞🌞🐵🐼🐼🐱🚗🚗💜🐵🐱💰🚗🌵💎💜🌴🌞🚗🐭🌍🌍🐵🍋💎🐵🌵🌵")
		})
	})
}
