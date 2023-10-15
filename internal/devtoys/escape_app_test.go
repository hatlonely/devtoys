package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEscapeFunctions(t *testing.T) {
	Convey("TestEscapeString", t, func() {
		So(escapeString(`"hello	world"`), ShouldEqual, `\"hello\tworld\"`)
	})

	Convey("TestUnescapeString", t, func() {
		So(unescapeString(`\"hello\tworld\"`), ShouldEqual, `"hello	world"`)
	})

	Convey("TestEscapeHtml", t, func() {
		So(escapeHtml(`<div class="w-full">`), ShouldEqual, "&lt;div class=&#34;w-full&#34;&gt;")
	})

	Convey("TestUnescapeHtml", t, func() {
		So(unescapeHtml("&lt;div class=&#34;w-full&#34;&gt;"), ShouldEqual, `<div class="w-full">`)
		So(unescapeHtml("&lt;div class=&quot;w-full&quot;&gt;"), ShouldEqual, `<div class="w-full">`)
	})

	Convey("TestEscapeUrl", t, func() {
		So(escapeUrl(`https://example.com/?q=hello world`), ShouldEqual, "https:%2F%2Fexample.com%2F%3Fq=hello%20world")
	})

	Convey("TestUnescapeUrl", t, func() {
		v, err := unescapeUrl("https:%2F%2Fexample.com%2F%3Fq=hello%20world")
		So(err, ShouldBeNil)
		So(v, ShouldEqual, `https://example.com/?q=hello world`)
	})

	Convey("TestEscapeUrlQuery", t, func() {
		So(escapeUrlQuery(`https://example.com/?q=hello world`), ShouldEqual, "https%3A%2F%2Fexample.com%2F%3Fq%3Dhello+world")
	})

	Convey("TestUnescapeUrlQuery", t, func() {
		v, err := unescapeUrlQuery("https%3A%2F%2Fexample.com%2F%3Fq%3Dhello+world")
		So(err, ShouldBeNil)
		So(v, ShouldEqual, `https://example.com/?q=hello world`)
	})
}

func TestEscape(t *testing.T) {
	app := NewEscapeApp()
	Convey("TestEscape", t, func() {
		Convey("escape string", func() {
			res, err := app.Escape(&EscapeReq{
				Text: `"hello	world"`,
				Mode: "escape",
				Type: "string",
			})
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, `\"hello\tworld\"`)
		})

		Convey("unescape string", func() {
			res, err := app.Escape(&EscapeReq{
				Text: `\"hello\tworld\"`,
				Mode: "unescape",
				Type: "string",
			})
			So(err, ShouldBeNil)
			So(res.Text, ShouldEqual, `"hello	world"`)
		})
	})
}
