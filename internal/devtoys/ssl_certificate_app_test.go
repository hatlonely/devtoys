package devtoys

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/hatlonely/go-kit/strx"
	. "github.com/smartystreets/goconvey/convey"
)

var cert = `-----BEGIN CERTIFICATE-----
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
Mb7kB97cNdEzacv03lwuL+JIa2g=
-----END CERTIFICATE-----
`

func TestX509(t *testing.T) {
	block, _ := pem.Decode([]byte(cert))
	cert, _ := x509.ParseCertificate(block.Bytes)
	fmt.Println(strx.JsonMarshalIndent(cert))
}

func TestFetchCertificate(t *testing.T) {
	txt, err := fetchCertificate("https://www.baidu.com")
	fmt.Println(err)
	fmt.Println(txt)
}

func TestSSLCertificateApp(t *testing.T) {
	Convey("Test SSLCertificateApp.SSLCertificate", t, func() {
		app := NewSSLCertificateApp()

		res, err := app.SSLCertificate(&SSLCertificateReq{
			Text: cert,
		})
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		fmt.Println(strx.JsonMarshalIndent(res))
	})
}
