package utils

import (
	"strings"
	"time"
)

// DateTimeUTC returns the current UTC time in RFC3339 format.
func DateTimeUTC() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// PanicOnError will panic if the error is not nil.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// MountableMarkdown adds the mountable class to all headers and paragraphs.
func MountableMarkdown(html string) string {
	processTag := func(code string, tag string) string {
		return strings.Replace(code, "<"+tag+">", "<"+tag+" class='mountable'>", -1)
	}

	html = processTag(html, "p")
	html = processTag(html, "ul")
	html = processTag(html, "ol")
	html = processTag(html, "blockquote")
	html = processTag(html, "pre")
	html = processTag(html, "h2")
	html = processTag(html, "h3")
	html = processTag(html, "h4")
	html = processTag(html, "h5")
	html = processTag(html, "h6")

	return html
}
