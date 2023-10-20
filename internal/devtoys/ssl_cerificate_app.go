package devtoys

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type SSLCertificateApp struct {
	ctx context.Context
}

func NewSSLCertificateApp() *SSLCertificateApp {
	return &SSLCertificateApp{}
}

func (a *SSLCertificateApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type SSLCertificateReq struct {
	Text string
}

type Subject struct {
	Country          []string // 国家
	Organization     []string // 组织
	OrganizationUnit []string // 组织单位
	CommonName       string   // 域名
	Province         []string // 省份
	Locality         []string // 地区
	StreetAddress    []string // 街道地址
	PostalCode       []string // 邮政编码
}

type Issuer struct {
	Country          []string // 国家
	Organization     []string // 组织
	OrganizationUnit []string // 组织单位
	CommonName       string   // 颁发机构
}

type Certificate struct {
	SerialNumber          string   // 证书序列号
	DNSNames              []string // DNS 域名
	NotBefore             string   // 证书生效时间
	NotAfter              string   // 证书失效时间
	CRLDistributionPoints []string // CRL 分发点
	IssuingCertificateURL []string // 颁发者证书 URL
	OCSPServer            []string // OCSP 服务器地址
	PublicKey             string   // 公钥
}

type SSLCertificateRes struct {
	Subject     *Subject
	Issuer      *Issuer
	Certificate *Certificate
}

func (a *SSLCertificateApp) SSLCertificate(req *SSLCertificateReq) (*SSLCertificateRes, error) {
	block, _ := pem.Decode([]byte(req.Text))
	if block == nil {
		return nil, errors.New("pem.Decode failed")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Errorf("x509.ParseCertificate failed")
	}

	fmt.Println(cert.PublicKeyAlgorithm)
	fmt.Println(reflect.TypeOf(cert.PublicKey))

	certificate := &Certificate{
		DNSNames:  cert.DNSNames,
		NotBefore: cert.NotBefore.Format("2006-01-02 15:04:05"),
		NotAfter:  cert.NotAfter.Format("2006-01-02 15:04:05"),
	}

	return &SSLCertificateRes{
		Issuer: &Issuer{
			Country:          cert.Issuer.Country,
			Organization:     cert.Issuer.Organization,
			OrganizationUnit: cert.Issuer.OrganizationalUnit,
			CommonName:       cert.Issuer.CommonName,
		},
		Subject: &Subject{
			Country:          cert.Subject.Country,
			Organization:     cert.Subject.Organization,
			OrganizationUnit: cert.Subject.OrganizationalUnit,
			CommonName:       cert.Subject.CommonName,
			Province:         cert.Subject.Province,
			Locality:         cert.Subject.Locality,
			StreetAddress:    cert.Subject.StreetAddress,
			PostalCode:       cert.Subject.PostalCode,
		},
		Certificate: certificate,
	}, nil
}
