package devtoys

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/hatlonely/go-kit/strx"
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

func (a *StringApp) String(req *StringReq) (*StringRes, error) {
	return nil, nil
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
