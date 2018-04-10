package util

import (
	"github.com/headzoo/surf/browser"
)

// Bowsable represents an bower surf utils
type Bowsable interface {
	// GetHTML get HTML from browser
	GetHTML(bow *browser.Browser) string
}

// Format Json to readable string
func GetHTML(bow *browser.Browser) string {
	return "<html>" + bow.Body() + "</html>"
}
