package devtoys

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
)

type AESApp struct {
	ctx context.Context
}

func NewAESApp() *AESApp {
	return &AESApp{}
}

func (a *AESApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type AESReq struct {
	Function       string `validate:"oneof=encrypt decrypt"`
	Text           string `validate:"required"`
	IV             string
	Base64Text     bool
	Key            string `validate:"required"`
	Base64Key      bool
	Type           string `validate:"oneof=AES128 AES192 AES256"`
	EncryptionMode string `validate:"oneof=cbc cfb ctr ofb"`
}

type AESRes struct {
	Text string
}

func (a *AESApp) AES(req *AESReq) (*AESRes, error) {
	if req.Function == "encrypt" {
		return a.doAESEncrypt(req)
	} else if req.Function == "decrypt" {
		return a.doAESDecrypt(req)
	}

	return nil, errors.New("invalid mode")
}

func (a *AESApp) generateKey(key string, type_ string, base64Key bool) ([]byte, error) {
	var err error
	var keyBytes []byte
	if base64Key {
		keyBytes, err = base64.StdEncoding.DecodeString(key)
		if err != nil {
			return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString failed")
		}
	} else {
		keyBytes = []byte(key)
	}

	var padKeyBytes []byte
	switch type_ {
	case "AES128":
		if len(keyBytes) > 16 {
			return nil, errors.New("key length should not be greater than 16")
		}
		padKeyBytes = make([]byte, 16)
		copy(padKeyBytes, keyBytes)
	case "AES192":
		if len(keyBytes) > 24 {
			return nil, errors.New("key length should not be greater than 24")
		}
		padKeyBytes = make([]byte, 24)
		copy(padKeyBytes, keyBytes)
	case "AES256":
		if len(keyBytes) > 32 {
			return nil, errors.New("key length should not be greater than 32")
		}
		padKeyBytes = make([]byte, 32)
		copy(padKeyBytes, keyBytes)
	default:
		return nil, errors.New("invalid type")
	}

	return padKeyBytes, nil
}

func (a *AESApp) doAESEncrypt(req *AESReq) (*AESRes, error) {
	var err error
	if err = validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	key, err := a.generateKey(req.Key, req.Type, req.Base64Key)
	if err != nil {
		return nil, errors.WithMessage(err, "generateKey failed")
	}

	// 创建加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "aes.NewCipher failed")
	}

	// 填充加密字符串
	var plainText []byte
	if req.Base64Text {
		plainText, err = base64.StdEncoding.DecodeString(req.Text)
		if err != nil {
			return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString failed")
		}
	} else {
		plainText = []byte(req.Text)
	}

	blockSize := block.BlockSize()
	padtext, err := pkcs7pad(plainText, blockSize)
	if err != nil {
		return nil, errors.WithMessage(err, "pkcs7pad failed")
	}

	// 加密
	iv := make([]byte, blockSize)
	if req.IV != "" {
		buf, err := base64.StdEncoding.DecodeString(req.IV)
		if err != nil {
			return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString failed")
		}
		if len(buf) > blockSize {
			return nil, errors.New("iv length should not be greater than block size")
		}
		copy(iv, buf)
	}
	cipherText := make([]byte, len(padtext))
	switch req.EncryptionMode {
	case "cbc":
		stream := cipher.NewCBCEncrypter(block, iv)
		stream.CryptBlocks(cipherText, padtext)
	case "cfb":
		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(cipherText, padtext)
	case "ctr":
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(cipherText, padtext)
	case "ofb":
		stream := cipher.NewOFB(block, iv)
		stream.XORKeyStream(cipherText, padtext)
	default:
		return nil, errors.New("invalid encryption mode")
	}

	// 加密

	return &AESRes{
		Text: base64.StdEncoding.EncodeToString(cipherText),
	}, nil
}

func (a *AESApp) doAESDecrypt(req *AESReq) (*AESRes, error) {
	var err error
	if err = validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	key, err := a.generateKey(req.Key, req.Type, req.Base64Key)
	if err != nil {
		return nil, errors.WithMessage(err, "generateKey failed")
	}

	var cipherText []byte
	cipherText, err = base64.StdEncoding.DecodeString(req.Text)
	if err != nil {
		return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString failed")
	}

	// 创建加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "aes.NewCipher failed")
	}

	// 解密
	plainText := make([]byte, len(cipherText))

	// 创建加密器
	iv := make([]byte, block.BlockSize())
	if req.IV != "" {
		buf, err := base64.StdEncoding.DecodeString(req.IV)
		if err != nil {
			return nil, errors.Wrap(err, "base64.StdEncoding.DecodeString failed")
		}
		if len(buf) > block.BlockSize() {
			return nil, errors.New("iv length should not be greater than block size")
		}
		copy(iv, buf)
	}

	switch req.EncryptionMode {
	case "cbc":
		stream := cipher.NewCBCDecrypter(block, iv)
		stream.CryptBlocks(plainText, cipherText)
	case "cfb":
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(plainText, cipherText)
	case "ctr":
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(plainText, cipherText)
	case "ofb":
		stream := cipher.NewOFB(block, iv)
		stream.XORKeyStream(plainText, cipherText)
	default:
		return nil, errors.New("invalid encryption mode")
	}

	// 去除填充
	plainText, err = pkcs7strip(plainText, block.BlockSize())
	if err != nil {
		return nil, errors.WithMessage(err, "pkcs7strip failed")
	}

	var encodedText string
	if req.Base64Text {
		encodedText = base64.StdEncoding.EncodeToString(plainText)
	} else {
		encodedText = string(plainText)
	}

	return &AESRes{
		Text: encodedText,
	}, nil
}

// pkcs7strip remove pkcs7 padding
func pkcs7strip(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7: Data is empty")
	}
	if length%blockSize != 0 {
		return nil, errors.New("pkcs7: Data is not block-aligned")
	}
	padLen := int(data[length-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)
	if padLen > blockSize || padLen == 0 || !bytes.HasSuffix(data, ref) {
		return nil, errors.New("pkcs7: Invalid padding")
	}
	return data[:length-padLen], nil
}

// pkcs7pad add pkcs7 padding
func pkcs7pad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 1 || blockSize >= 256 {
		return nil, fmt.Errorf("pkcs7: Invalid block size %d", blockSize)
	} else {
		padLen := blockSize - len(data)%blockSize
		padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
		return append(data, padding...), nil
	}
}
