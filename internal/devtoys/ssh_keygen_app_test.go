package devtoys

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSSHKeyGeneratorAppGenerateSSHKey_RSA(t *testing.T) {

	Convey("RSA", t, func() {
		a := NewSSHKeyGeneratorApp()
		res, err := a.GenerateSSHKey(&SSHKeyGeneratorReq{
			Type: "rsa",
		})

		So(err, ShouldBeNil)
		fmt.Println(res.PrivateKey)
		fmt.Println(res.PublicKey)
	})

}

func TestSSHKeyGeneratorAppGenerateSSHKey_ECDSA(t *testing.T) {

	Convey("ECDSA", t, func() {
		a := NewSSHKeyGeneratorApp()
		res, err := a.GenerateSSHKey(&SSHKeyGeneratorReq{
			Type: "ecdsa",
		})

		So(err, ShouldBeNil)
		fmt.Println(res.PrivateKey)
		fmt.Println(res.PublicKey)
	})
}

func TestSSHKeyGeneratorAppGenerateSSHKey_ED25519(t *testing.T) {
	Convey("ED25519", t, func() {
		a := NewSSHKeyGeneratorApp()
		res, err := a.GenerateSSHKey(&SSHKeyGeneratorReq{
			Type: "ed25519",
		})

		So(err, ShouldBeNil)
		fmt.Println(res.PrivateKey)
		fmt.Println(res.PublicKey)
	})
}
