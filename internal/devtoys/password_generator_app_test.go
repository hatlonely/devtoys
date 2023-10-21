package devtoys

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPasswordGeneratorApp(t *testing.T) {
	Convey("TestPasswordGeneratorApp", t, func() {
		app := NewPasswordGeneratorApp()

		res, err := app.GeneratePassword(&GeneratePasswordReq{
			Length:       16,
			CharacterSet: "0123456789",
		})
		So(err, ShouldBeNil)
		fmt.Println(res.Text)
	})
}
