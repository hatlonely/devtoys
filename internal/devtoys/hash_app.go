package devtoys

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type HashApp struct {
	ctx context.Context
}

func NewHashApp() *HashApp {
	return &HashApp{}
}

func (a *HashApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type HashReq struct {
	Text string
	Type string `validate:"required,oneof=md5 sha1 sha256 sha224 sha512"`
}

type HashRes struct {
	Text string
}

func (a *HashApp) Hash(req *HashReq) (*HashRes, error) {
	res := &HashRes{}
	switch req.Type {
	case "md5":
		res.Text = md5sum(req.Text)
	case "sha1":
		res.Text = sha1sum(req.Text)
	case "sha256":
		res.Text = sha256sum(req.Text)
	case "sha224":
		res.Text = sha224sum(req.Text)
	case "sha512":
		res.Text = sha512sum(req.Text)
	default:
		res.Text = "unknown hash type"
	}
	return res, nil
}

func md5sum(text string) string {
	md5 := md5.Sum([]byte(text))
	return hex.EncodeToString(md5[:])
}

func sha1sum(text string) string {
	sha1 := sha1.Sum([]byte(text))
	return hex.EncodeToString(sha1[:])
}

func sha256sum(text string) string {
	sha256 := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sha256[:])
}

func sha224sum(text string) string {
	sha224 := sha256.Sum224([]byte(text))
	return hex.EncodeToString(sha224[:])
}

func sha512sum(text string) string {
	sha512 := sha512.Sum512([]byte(text))
	return hex.EncodeToString(sha512[:])
}
