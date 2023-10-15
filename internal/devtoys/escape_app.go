package devtoys

import (
	"bytes"
	"html"
	"net/url"

	"github.com/pkg/errors"
)

type EscapeApp struct {
	BaseApp
}

func NewEscapeApp() *EscapeApp {
	return &EscapeApp{}
}

type EscapeReq struct {
	Text string
	Mode string `validate:"oneof=escape unescape"`
	Type string `validate:"oneof=html string url urlquery"`
}

type EscapeRes struct {
	Text string
}

func (a *EscapeApp) Escape(req *EscapeReq) (*EscapeRes, error) {
	if err := validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	var text string
	if req.Mode == "escape" {
		switch req.Type {
		case "html":
			text = escapeHtml(req.Text)
		case "url":
			text = escapeUrl(req.Text)
		case "urlquery":
			text = escapeUrlQuery(req.Text)
		case "string":
			text = escapeString(req.Text)
		}
	} else if req.Mode == "unescape" {
		switch req.Type {
		case "html":
			text = unescapeHtml(req.Text)
		case "url":
			t, err := unescapeUrl(req.Text)
			if err != nil {
				return nil, errors.Wrap(err, "unescapeUrl failed")
			}
			text = t
		case "urlquery":
			t, err := unescapeUrlQuery(req.Text)
			if err != nil {
				return nil, errors.Wrap(err, "unescapeUrlQuery failed")
			}
			text = t
		case "string":
			text = unescapeString(req.Text)
		}
	}

	return &EscapeRes{
		Text: text,
	}, nil
}

func escapeHtml(text string) string {
	return html.EscapeString(text)
}

func unescapeHtml(text string) string {
	return html.UnescapeString(text)
}

func escapeUrl(text string) string {
	return url.PathEscape(text)
}

func unescapeUrl(text string) (string, error) {
	return url.PathUnescape(text)
}

func escapeUrlQuery(text string) string {
	return url.QueryEscape(text)
}

func unescapeUrlQuery(text string) (string, error) {
	return url.QueryUnescape(text)
}

func escapeString(text string) string {
	var buf bytes.Buffer
	for _, r := range text {
		switch r {
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case '\\':
			buf.WriteString(`\\`)
		case '"':
			buf.WriteString(`\"`)
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func unescapeString(text string) string {
	var buf bytes.Buffer
	for i := 0; i < len(text); i++ {
		if text[i] == '\\' && i+1 < len(text) {
			switch text[i+1] {
			case 'n':
				buf.WriteRune('\n')
				i++
			case 'r':
				buf.WriteRune('\r')
				i++
			case 't':
				buf.WriteRune('\t')
				i++
			case '\\':
				buf.WriteRune('\\')
				i++
			case '"':
				buf.WriteRune('"')
				i++
			default:
				buf.WriteByte(text[i])
			}
		} else {
			buf.WriteByte(text[i])
		}
	}
	return buf.String()
}
