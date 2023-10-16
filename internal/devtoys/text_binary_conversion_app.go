package devtoys

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
)

type TextBinaryConversionApp struct {
	BaseApp
}

func NewTextBinaryConversionApp() *TextBinaryConversionApp {
	return &TextBinaryConversionApp{}
}

type TextBinaryConversionReq struct {
	Text            string
	To              string `validate:"required,oneof=bin hex dec base64"`
	LowerCase       bool
	WithoutSpace    bool
	WithoutFillZero bool
}

type TextBinaryConversionRes struct {
	Text string
}

func (a *TextBinaryConversionApp) ConvertTextBinary(req *TextBinaryConversionReq) (*TextBinaryConversionRes, error) {
	if req.Text == "" {
		return &TextBinaryConversionRes{}, nil
	}

	if err := validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	var text string
	switch req.To {
	case "bin":
		text = strToBin(req.Text)
	case "hex":
		text = strToHex(req.Text)
	case "dec":
		text = strToDec(req.Text)
	case "base64":
		text = strToBase64(req.Text)
	}

	if !req.LowerCase && req.To == "hex" {
		text = strings.ToUpper(text)
	}

	if req.WithoutFillZero {
		text = strings.Join(funk.Map(strings.Split(text, " "), func(v string) string {
			return strings.TrimLeft(v, "0")
		}).([]string), " ")
	}

	if req.WithoutSpace {
		text = strings.ReplaceAll(text, " ", "")
	}

	return &TextBinaryConversionRes{
		Text: text,
	}, nil
}

func strToBin(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		v := strconv.FormatInt(int64(c), 2)
		for i := 0; i < (8-len(v)%8)%8; i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(v)
		buf.WriteByte(' ')
	}
	return buf.String()[:buf.Len()-1]
}

func strToHex(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		v := strconv.FormatInt(int64(c), 16)
		for i := 0; i < (2-len(v)%2)%2; i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(v)
		buf.WriteByte(' ')
	}

	return buf.String()[:buf.Len()-1]
}

func strToDec(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		buf.WriteString(strconv.Itoa(int(c)))
		buf.WriteByte(' ')
	}

	return buf.String()[:buf.Len()-1]
}

func strToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
