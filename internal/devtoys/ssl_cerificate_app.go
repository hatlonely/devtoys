package devtoys

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"net/http"

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
	Link string
	Text string `validate:"requiredwithout=Link"`
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
	SignatureAlgorithm    string   // 签名算法
	PublicKeyAlgorithm    string   // 公钥算法
	PublicKey             string   // 公钥
	AuthorityKeyId        string   // 颁发者公钥标识
	SubjectKeyId          string   // 使用者公钥标识

	Md5    string
	Sha1   string
	Sha256 string
}

type SSLCertificateRes struct {
	Subject     *Subject
	Issuer      *Issuer
	Certificate *Certificate
}

func fetchCertificate(link string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(link)
	if err != nil {
		return "", errors.Wrap(err, "client.Get failed")
	}
	defer resp.Body.Close()

	certs := resp.TLS.PeerCertificates
	if len(certs) == 0 {
		return "", errors.New("no certificates found")
	}
	cert := certs[0]

	pemCert := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	return string(pemCert), nil
}

func (a *SSLCertificateApp) SSLCertificate(req *SSLCertificateReq) (*SSLCertificateRes, error) {
	var err error

	if req.Link != "" {
		req.Text, err = fetchCertificate(req.Link)
		if err != nil {
			return nil, errors.WithMessage(err, "fetchCertificate")
		}
	}

	block, _ := pem.Decode([]byte(req.Text))
	if block == nil {
		return nil, errors.New("pem.Decode failed")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Errorf("x509.ParseCertificate failed")
	}

	publicKey := ""
	switch k := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		buf, err := x509.MarshalPKIXPublicKey(k)
		if err != nil {
			break
		}
		pem := pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: buf,
		})
		publicKey = string(pem)
	}

	certificate := &Certificate{
		SerialNumber:          cert.SerialNumber.String(),
		DNSNames:              cert.DNSNames,
		NotBefore:             cert.NotBefore.Format("2006-01-02 15:04:05"),
		NotAfter:              cert.NotAfter.Format("2006-01-02 15:04:05"),
		CRLDistributionPoints: cert.CRLDistributionPoints,
		IssuingCertificateURL: cert.IssuingCertificateURL,
		OCSPServer:            cert.OCSPServer,
		SignatureAlgorithm:    cert.SignatureAlgorithm.String(),
		PublicKeyAlgorithm:    cert.PublicKeyAlgorithm.String(),
		PublicKey:             publicKey,
		AuthorityKeyId:        byteToHex(cert.AuthorityKeyId),
		SubjectKeyId:          byteToHex(cert.SubjectKeyId),
		Md5:                   hex.EncodeToString(md5sum(cert.Raw)),
		Sha1:                  hex.EncodeToString(sha1sum(cert.Raw)),
		Sha256:                hex.EncodeToString(sha256sum(cert.Raw)),
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
