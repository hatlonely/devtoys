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
	In              string
	InType          string `validate:"required,oneof=bin hex dec raw"`
	ToType          string `validate:"required,oneof=bin hex dec raw"`
	AsByte          bool
	LowerCase       bool
	WithoutFillZero bool
}

type TextBinaryConversionRes struct {
	To string
}

func (a *TextBinaryConversionApp) ConvertTextBinary(req *TextBinaryConversionReq) (*TextBinaryConversionRes, error) {
	if req.In == "" {
		return &TextBinaryConversionRes{}, nil
	}
	if req.InType == "" {
		req.InType = "raw"
	}
	if req.ToType == "" {
		req.ToType = "raw"
	}

	if err := validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	var text string
	var err error

	if req.AsByte {
		text, err = convertByte(req.In, req.InType, req.ToType)
	} else {
		text, err = convertText(req.In, req.InType, req.ToType)
	}

	if err != nil {
		return nil, errors.Wrap(err, "convertToText failed")
	}

	if !req.LowerCase && req.ToType == "hex" {
		text = strings.ToUpper(text)
	}

	if req.WithoutFillZero {
		text = strings.Join(funk.Map(strings.Split(text, " "), func(v string) string {
			return strings.TrimLeft(v, "0")
		}).([]string), " ")
	}

	return &TextBinaryConversionRes{
		To: text,
	}, nil
}

func convertText(in string, inType string, toType string) (string, error) {
	var text string
	var err error

	text, err = convertToText(in, inType)
	if err != nil {
		return "", errors.WithMessage(err, "convertToText failed")
	}

	text, err = convertFromText(text, toType)
	if err != nil {
		return "", errors.WithMessage(err, "convertFromText failed")
	}

	return text, nil
}

func convertByte(in string, inType string, toType string) (string, error) {
	var buf []byte
	var err error

	buf, err = convertToByte(in, inType)
	if err != nil {
		return "", errors.WithMessage(err, "convertFromByte failed")
	}

	text, err := convertFromByte(buf, toType)
	if err != nil {
		return "", errors.WithMessage(err, "convertFromByte failed")
	}

	return text, nil
}

func convertToText(in string, inType string) (string, error) {
	switch inType {
	case "bin":
		return binToText(in), nil
	case "hex":
		return hexToText(in), nil
	case "dec":
		return decToText(in), nil
	case "raw":
		return in, nil
	}

	return "", errors.New("invalid inType")
}

func convertToByte(in string, inType string) ([]byte, error) {
	switch inType {
	case "bin":
		return binToByte(in), nil
	case "hex":
		return hexToByte(in), nil
	case "dec":
		return decToByte(in), nil
	case "raw":
		return base64.StdEncoding.DecodeString(in)
	}

	return nil, errors.New("invalid inType")
}

func convertFromText(in string, toType string) (string, error) {
	switch toType {
	case "bin":
		return textToBin(in), nil
	case "hex":
		return textToHex(in), nil
	case "dec":
		return textToDec(in), nil
	case "raw":
		return in, nil
	}

	return "", errors.New("invalid toType")
}

func convertFromByte(in []byte, toType string) (string, error) {
	switch toType {
	case "bin":
		return byteToBin(in), nil
	case "hex":
		return byteToHex(in), nil
	case "dec":
		return byteToDec(in), nil
	case "raw":
		return base64.StdEncoding.EncodeToString(in), nil
	}

	return "", errors.New("invalid toType")
}

func textToBin(text string) string {
	var buf bytes.Buffer

	for _, c := range text {
		v := strconv.FormatInt(int64(c), 2)
		for i := 0; i < (8-len(v)%8)%8; i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(v)
		buf.WriteByte(' ')
	}
	return buf.String()[:buf.Len()-1]
}

func binToText(bin string) string {
	var buf bytes.Buffer

	for _, c := range strings.Split(bin, " ") {
		v, err := strconv.ParseInt(c, 2, 64)
		if err != nil {
			panic(err)
		}
		buf.WriteRune(rune(v))
	}

	return buf.String()
}

func byteToBin(inBytes []byte) string {
	var buf bytes.Buffer

	for _, c := range inBytes {
		v := strconv.FormatInt(int64(c), 2)
		for i := 0; i < (8-len(v)%8)%8; i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(v)
		buf.WriteByte(' ')
	}
	return buf.String()[:buf.Len()-1]
}

func binToByte(bin string) []byte {
	var buf bytes.Buffer

	for _, c := range strings.Split(bin, " ") {
		v, err := strconv.ParseInt(c, 2, 64)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(byte(v))
	}

	return buf.Bytes()
}

func textToHex(text string) string {
	var buf bytes.Buffer

	for _, c := range text {
		v := strconv.FormatInt(int64(c), 16)
		for i := 0; i < (2-len(v)%2)%2; i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(v)
		buf.WriteByte(' ')
	}

	return buf.String()[:buf.Len()-1]
}

func hexToText(hex string) string {
	var buf bytes.Buffer

	for _, c := range strings.Split(hex, " ") {
		v, err := strconv.ParseInt(c, 16, 64)
		if err != nil {
			panic(err)
		}
		buf.WriteRune(rune(v))
	}

	return buf.String()
}

func byteToHex(inBytes []byte) string {
	var buf bytes.Buffer

	for _, c := range inBytes {
		v := strconv.FormatInt(int64(c), 16)
		for i := 0; i < (2-len(v)%2)%2; i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(v)
		buf.WriteByte(' ')
	}

	return buf.String()[:buf.Len()-1]
}

func hexToByte(hex string) []byte {
	var buf bytes.Buffer

	for _, c := range strings.Split(hex, " ") {
		v, err := strconv.ParseInt(c, 16, 64)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(byte(v))
	}

	return buf.Bytes()
}

func textToDec(text string) string {
	var buf bytes.Buffer

	for _, c := range text {
		buf.WriteString(strconv.Itoa(int(c)))
		buf.WriteByte(' ')
	}

	return buf.String()[:buf.Len()-1]
}

func decToText(dec string) string {
	var buf bytes.Buffer

	for _, c := range strings.Split(dec, " ") {
		v, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			panic(err)
		}
		buf.WriteRune(rune(v))
	}

	return buf.String()
}

func byteToDec(inBytes []byte) string {
	var buf bytes.Buffer

	for _, c := range inBytes {
		buf.WriteString(strconv.Itoa(int(c)))
		buf.WriteByte(' ')
	}

	return buf.String()[:buf.Len()-1]
}

func decToByte(dec string) []byte {
	var buf bytes.Buffer

	for _, c := range strings.Split(dec, " ") {
		v, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(byte(v))
	}

	return buf.Bytes()
}
