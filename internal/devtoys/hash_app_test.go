package devtoys

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMd5(t *testing.T) {
	Convey("TestMd5", t, func() {
		So(md5sum("hello world"), ShouldEqual, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	})
}

func TestSha1(t *testing.T) {
	Convey("TestSha1", t, func() {
		So(sha1sum("hello world"), ShouldEqual, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	})
}

func TestSha256(t *testing.T) {
	Convey("TestSha256", t, func() {
		So(sha256sum("hello world"), ShouldEqual, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
	})
}

func TestSha224(t *testing.T) {
	Convey("TestSha224", t, func() {
		So(sha224sum("hello world"), ShouldEqual, "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
	})
}

func TestSha512(t *testing.T) {
	Convey("TestSha512", t, func() {
		So(sha512sum("hello world"), ShouldEqual, "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
	})
}
