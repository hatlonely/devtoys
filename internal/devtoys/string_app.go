package devtoys

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/hatlonely/go-kit/strx"
	"github.com/pkg/errors"
)

type StringApp struct {
	BaseApp
}

func NewStringApp() *StringApp {
	return &StringApp{}
}

type StringReq struct {
	Text string
	Type string
}

type StringRes struct {
	Text string
}

func (a *StringApp) ConvertString(req *StringReq) (*StringRes, error) {
	switch req.Type {
	case "title":
		return &StringRes{Text: titleCase(req.Text)}, nil
	case "lower":
		return &StringRes{Text: lowerCase(req.Text)}, nil
	case "upper":
		return &StringRes{Text: upperCase(req.Text)}, nil
	case "sentence":
		return &StringRes{Text: sentenceCase(req.Text)}, nil
	case "camel":
		return &StringRes{Text: camelCase(req.Text)}, nil
	case "pascal":
		return &StringRes{Text: pascalCase(req.Text)}, nil
	case "snake":
		return &StringRes{Text: snakeCase(req.Text)}, nil
	case "kebab":
		return &StringRes{Text: kebabCase(req.Text)}, nil
	case "snakeCaseAllCaps":
		return &StringRes{Text: snakeCaseAllCaps(req.Text)}, nil
	case "kebabCaseAllCaps":
		return &StringRes{Text: kebabCaseAllCaps(req.Text)}, nil
	case "train":
		return &StringRes{Text: trainCase(req.Text)}, nil
	default:
		return nil, errors.Errorf("unknown type: %v", req.Type)
	}
}

func titleCase(text string) string {
	var buf bytes.Buffer
	prevChar := ' '

	for _, c := range text {
		if prevChar == ' ' || prevChar == '-' || prevChar == '_' ||
			prevChar > unicode.MaxASCII {
			buf.WriteRune(unicode.ToUpper(c))
		} else {
			buf.WriteRune(c)
		}
		prevChar = c
	}

	return buf.String()
}

func lowerCase(text string) string {
	return strings.ToLower(text)
}

func upperCase(text string) string {
	return strings.ToUpper(text)
}

func sentenceCase(text string) string {
	var buf bytes.Buffer
	prevChar := '.'

	for _, c := range text {
		if prevChar == '.' || prevChar == '!' || prevChar == '?' || prevChar == '\n' || prevChar == '\r' {
			buf.WriteRune(unicode.ToUpper(c))
		} else {
			buf.WriteRune(c)
		}
		if c != ' ' && c != '\t' {
			prevChar = c
		}
	}

	return buf.String()
}

func camelCase(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, strx.CamelName(line))
	}
	return strings.Join(strs, "\n")
}

func pascalCase(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, strx.PascalName(line))
	}
	return strings.Join(strs, "\n")
}

func snakeCase(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, strx.SnakeName(line))
	}
	return strings.Join(strs, "\n")
}

func kebabCase(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, strx.KebabName(line))
	}
	return strings.Join(strs, "\n")
}

func snakeCaseAllCaps(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, strings.ToUpper(strx.SnakeName(line)))
	}
	return strings.Join(strs, "\n")
}

func kebabCaseAllCaps(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, strings.ToUpper(strx.KebabName(line)))
	}
	return strings.Join(strs, "\n")
}

func trainCase(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var strs []string
	for _, line := range lines {
		strs = append(strs, titleCase(strx.KebabName(line)))
	}
	return strings.Join(strs, "\n")
}
