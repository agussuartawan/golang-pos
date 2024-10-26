package helper

import (
	"html"
	"strings"
)

func TrimSpace(value string) string {
	return html.EscapeString(strings.TrimSpace(value))
}