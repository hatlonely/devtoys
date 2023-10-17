package devtoys

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
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
	Text   string
	Type   string `validate:"required,oneof=md5 sha1 sha256 sha224 sha512"`
	Hmac   bool
	Key    string
	Encode string `validate:"oneof=hex base64"`
}

type HashRes struct {
	Text string
}

func (a *HashApp) Hash(req *HashReq) (*HashRes, error) {
	var buf []byte
	if req.Hmac {
		switch req.Type {
		case "md5":
			buf = hmacMd5([]byte(req.Key), []byte(req.Text))
		case "sha1":
			buf = hmacSha1([]byte(req.Key), []byte(req.Text))
		case "sha256":
			buf = hmacSha256([]byte(req.Key), []byte(req.Text))
		case "sha224":
			buf = hmacSha224([]byte(req.Key), []byte(req.Text))
		case "sha512":
			buf = hmacSha512([]byte(req.Key), []byte(req.Text))
		}
	} else {
		switch req.Type {
		case "md5":
			buf = md5sum([]byte(req.Text))
		case "sha1":
			buf = sha1sum([]byte(req.Text))
		case "sha256":
			buf = sha256sum([]byte(req.Text))
		case "sha224":
			buf = sha224sum([]byte(req.Text))
		case "sha512":
			buf = sha512sum([]byte(req.Text))
		}
	}

	var text string
	if req.Encode == "hex" {
		text = hex.EncodeToString(buf)
	} else {
		text = base64.StdEncoding.EncodeToString(buf)
	}

	return &HashRes{
		Text: text,
	}, nil
}

func md5sum(text []byte) []byte {
	md5 := md5.Sum([]byte(text))
	return md5[:]
}

func sha1sum(text []byte) []byte {
	sha1 := sha1.Sum([]byte(text))
	return sha1[:]
}

func sha256sum(text []byte) []byte {
	sha256 := sha256.Sum256([]byte(text))
	return sha256[:]
}

func sha224sum(text []byte) []byte {
	sha224 := sha256.Sum224([]byte(text))
	return sha224[:]
}

func sha512sum(text []byte) []byte {
	sha512 := sha512.Sum512([]byte(text))
	return sha512[:]
}

func hmacMd5(key []byte, text []byte) []byte {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(text))
	return hmac.Sum(nil)
}

func hmacSha1(key []byte, text []byte) []byte {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write([]byte(text))
	return hmac.Sum(nil)
}

func hmacSha256(key []byte, text []byte) []byte {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(text))
	return hmac.Sum(nil)
}

func hmacSha224(key []byte, text []byte) []byte {
	hmac := hmac.New(sha256.New224, []byte(key))
	hmac.Write([]byte(text))
	return hmac.Sum(nil)
}

func hmacSha512(key []byte, text []byte) []byte {
	hmac := hmac.New(sha512.New, []byte(key))
	hmac.Write([]byte(text))
	return hmac.Sum(nil)
}
