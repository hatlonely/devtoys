package devtoys

import (
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMd5(t *testing.T) {
	Convey("TestMd5", t, func() {
		buf := md5sum([]byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	})
}

func TestSha1(t *testing.T) {
	Convey("TestSha1", t, func() {
		buf := sha1sum([]byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	})
}

func TestSha256(t *testing.T) {
	Convey("TestSha256", t, func() {
		buf := sha256sum([]byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
	})
}

func TestSha224(t *testing.T) {
	Convey("TestSha224", t, func() {
		buf := sha224sum([]byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
	})
}

func TestSha512(t *testing.T) {
	Convey("TestSha512", t, func() {
		buf := sha512sum([]byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
	})
}

func TestHmacMd5(t *testing.T) {
	Convey("TestHmacMd5", t, func() {
		buf := hmacMd5([]byte("123456"), []byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "9028905855a81e0f3f76a72212e6c488")
	})
}

func TestHmacSha1(t *testing.T) {
	Convey("TestHmacSha1", t, func() {
		buf := hmacSha1([]byte("123456"), []byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "4873e80b320876fef98bd5857a7146db412007c7")
	})
}

func TestHmacSha256(t *testing.T) {
	Convey("TestHmacSha256", t, func() {
		buf := hmacSha256([]byte("123456"), []byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "83b3eb2788457b46a2f17aaa048f795af0d9dabb8e5924dd2fc0ea682d929fe5")
	})
}

func TestHmacSha224(t *testing.T) {
	Convey("TestHmacSha224", t, func() {
		buf := hmacSha224([]byte("123456"), []byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "4f358a44c7bfaf92a2aa3e015d06a6913cc22dd1d684efbb3ce71755")
	})
}

func TestHmacSha512(t *testing.T) {
	Convey("TestHmacSha512", t, func() {
		buf := hmacSha512([]byte("123456"), []byte("hello world"))
		So(hex.EncodeToString(buf), ShouldEqual, "6c9c251365f3507dc923023fd8e180925eee0dc0bb467d156edc21b9889fc1115cbd7a948090abb59b31718e83978900d7582993392d90d2835ee13c9f2fbb69")
	})
}
