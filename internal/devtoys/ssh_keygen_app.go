package devtoys

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type SSHKeyGeneratorApp struct {
	ctx context.Context
}

func NewSSHKeyGeneratorApp() *SSHKeyGeneratorApp {
	return &SSHKeyGeneratorApp{}
}

func (a *SSHKeyGeneratorApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type SSHKeyGeneratorReq struct {
	Type string `validate:"required,oneof=rsa ecdsa ed25519"`
}

type SSHKeyGeneratorRes struct {
	PrivateKey string
	PublicKey  string
}

func (a *SSHKeyGeneratorApp) GenerateSSHKey(req *SSHKeyGeneratorReq) (*SSHKeyGeneratorRes, error) {
	switch req.Type {
	case "rsa":
		return a.generateRSAKey()
	case "ecdsa":
		return a.generateECDSAKey()
	case "ed25519":
		return a.generateED25519()
	default:
		return nil, errors.New("unsupported key type")
	}
}

func (a *SSHKeyGeneratorApp) generateRSAKey() (*SSHKeyGeneratorRes, error) {
	// 生成 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate private key")
	}

	// 将私钥转换为 PEM 格式
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// 生成公钥
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate public key")
	}

	// 将公钥转换为字符串格式
	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)
	publicKeyStr := string(publicKeyBytes)

	// 返回 SSH 密钥对
	return &SSHKeyGeneratorRes{
		PrivateKey: string(privateKeyPEM),
		PublicKey:  publicKeyStr,
	}, nil
}

func (a *SSHKeyGeneratorApp) generateECDSAKey() (*SSHKeyGeneratorRes, error) {
	// 生成 ECDSA 密钥对
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate ECDSA key")
	}

	// 将私钥转换为 PEM 格式
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal ECDSA private key")
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// 生成公钥
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate ECDSA public key")
	}

	// 将公钥转换为字符串格式
	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)
	publicKeyStr := string(publicKeyBytes)

	// 返回 SSH 密钥对
	return &SSHKeyGeneratorRes{
		PrivateKey: string(privateKeyPEM),
		PublicKey:  publicKeyStr,
	}, nil
}

func (a *SSHKeyGeneratorApp) generateED25519() (*SSHKeyGeneratorRes, error) {
	// 生成 Ed25519 密钥对
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate Ed25519 key")
	}

	// 将私钥转换为 PEM 格式
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: privateKey,
	})

	// 生成公钥
	sshPublicKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate DSA public key")
	}

	// 生成公钥
	publicKeyBytes := ssh.MarshalAuthorizedKey(sshPublicKey)
	publicKeyStr := string(publicKeyBytes)

	// 返回 SSH 密钥对
	return &SSHKeyGeneratorRes{
		PrivateKey: string(privateKeyPEM),
		PublicKey:  publicKeyStr,
	}, nil
}
