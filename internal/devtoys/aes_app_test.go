package devtoys

import (
	"encoding/base64"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAesCBC(t *testing.T) {
	app := NewAESApp()

	Convey("TestAESCBC", t, func() {
		iv := base64.StdEncoding.EncodeToString([]byte("123456"))

		eres, err := app.AES(&AESReq{
			Function:       "encrypt",
			Text:           "hello world",
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "cbc",
		})
		So(err, ShouldBeNil)
		So(eres, ShouldNotBeNil)
		So(eres.Text, ShouldEqual, "NbyzrtCoJcn1fZtwrewELw==")

		dres, err := app.AES(&AESReq{
			Function:       "decrypt",
			Text:           eres.Text,
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "cbc",
		})
		So(err, ShouldBeNil)
		So(dres, ShouldNotBeNil)
		So(dres.Text, ShouldEqual, "hello world")
	})
}

func TestAesCFB(t *testing.T) {
	app := NewAESApp()

	Convey("TestAESCFB", t, func() {
		iv := base64.StdEncoding.EncodeToString([]byte("123456"))

		eres, err := app.AES(&AESReq{
			Function:       "encrypt",
			Text:           "hello world hello world hello world hello world hello world hello world hello world hello world hello world",
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "cfb",
		})
		So(err, ShouldBeNil)
		So(eres, ShouldNotBeNil)
		So(eres.Text, ShouldEqual, "fhUb7huWfmFmWw+/QTdGPjikDSSboQo0GzBzK+RCSqwoeA/i2sCiO9/9tMU5rlwTw/FWeKYBswdQPkKxaF+QSkyQjw1QQ9PJWwL2jTIU43j0OkvxCxEMB/AHIh03OqkAWlbh2FVfNzppaU0Slk6rTg==")

		dres, err := app.AES(&AESReq{
			Function:       "decrypt",
			Text:           eres.Text,
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "cfb",
		})
		So(err, ShouldBeNil)
		So(dres, ShouldNotBeNil)
		So(dres.Text, ShouldEqual, "hello world hello world hello world hello world hello world hello world hello world hello world hello world")
	})
}

func TestAesCTR(t *testing.T) {
	app := NewAESApp()

	Convey("TestAESCFB", t, func() {
		iv := base64.StdEncoding.EncodeToString([]byte("123456"))

		eres, err := app.AES(&AESReq{
			Function:       "encrypt",
			Text:           "hello world hello world hello world hello world hello world hello world hello world hello world hello world",
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "ctr",
		})
		So(err, ShouldBeNil)
		So(eres, ShouldNotBeNil)
		So(eres.Text, ShouldEqual, "fhUb7huWfmFmWw+/QTdGPpL2vmPO3FgXCN5MvWWSRVkmXa1IyZvn8jlJTejNhDKFiEQpKCxVAEHnVWYZRjQsCpgaxf/BmjivrDiRXmg74mzqn5nXb0fp4nUwxWvIOZ126F9cWbVu6xnqPYHSQvtl9A==")

		dres, err := app.AES(&AESReq{
			Function:       "decrypt",
			Text:           eres.Text,
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "ctr",
		})
		So(err, ShouldBeNil)
		So(dres, ShouldNotBeNil)
		So(dres.Text, ShouldEqual, "hello world hello world hello world hello world hello world hello world hello world hello world hello world")
	})
}

func TestAesOFB(t *testing.T) {
	app := NewAESApp()

	Convey("TestAESCFB", t, func() {
		iv := base64.StdEncoding.EncodeToString([]byte("123456"))

		eres, err := app.AES(&AESReq{
			Function:       "encrypt",
			Text:           "hello world hello world hello world hello world hello world hello world hello world hello world hello world",
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "ofb",
		})
		So(err, ShouldBeNil)
		So(eres, ShouldNotBeNil)
		So(eres.Text, ShouldEqual, "fhUb7huWfmFmWw+/QTdGPoORYB7KAda7aOSTR1O7o+8Wn/UBccnH9FHcEmFuN0kUerOSqxs7+ZRk+O7/B+/FwML0gkHBFXFb5q9E8RpUpTAoboOMUkzDT+FFkvw2+ho6ScRhvS6TcAPMzH6oRH4/Dg==")

		dres, err := app.AES(&AESReq{
			Function:       "decrypt",
			Text:           eres.Text,
			IV:             iv,
			Base64Text:     false,
			Key:            "1234567890123456",
			Base64Key:      false,
			Type:           "AES128",
			EncryptionMode: "ofb",
		})
		So(err, ShouldBeNil)
		So(dres, ShouldNotBeNil)
		So(dres.Text, ShouldEqual, "hello world hello world hello world hello world hello world hello world hello world hello world hello world")
	})
}
