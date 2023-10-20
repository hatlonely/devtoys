package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTitleCase(t *testing.T) {
	Convey("TestTitleCase", t, func() {
		So(titleCase("hello world"), ShouldEqual, "Hello World")
		So(titleCase("你好hello_world"), ShouldEqual, "你好Hello_World")
	})
}

func TestSentenceCase(t *testing.T) {
	Convey("TestSentenceCase", t, func() {
		So(sentenceCase("hello world"), ShouldEqual, "Hello world")
		So(sentenceCase("hello. world"), ShouldEqual, "Hello. World")
		So(sentenceCase("你好hello_world"), ShouldEqual, "你好hello_world")
		So(sentenceCase("hello\nworld"), ShouldEqual, "Hello\nWorld")
	})
}

func TestCamelCase(t *testing.T) {
	Convey("TestCamelCase", t, func() {
		So(camelCase("hello world"), ShouldEqual, "helloWorld")
		So(camelCase("HelloWorld"), ShouldEqual, "helloWorld")
		So(camelCase("Hello World"), ShouldEqual, "helloWorld")
		So(camelCase("hello world\nhello world"), ShouldEqual, "helloWorld\nhelloWorld")
		So(camelCase("你好hello_world"), ShouldEqual, "helloWorld")
	})
}

func TestTrainCase(t *testing.T) {
	Convey("TestTrainCase", t, func() {
		So(trainCase("hello world"), ShouldEqual, "Hello-World")
		So(trainCase("HelloWorld"), ShouldEqual, "Hello-World")
		So(trainCase("Hello World"), ShouldEqual, "Hello-World")
		So(trainCase("hello_world"), ShouldEqual, "Hello-World")
	})
}
