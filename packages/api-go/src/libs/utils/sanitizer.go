package utils

import "github.com/microcosm-cc/bluemonday"

var htmlSanitizer = bluemonday.UGCPolicy()

// SanitizeHTML sanitizes HTML input by removing potentially unsafe content
// It uses the bluemonday UGCPolicy which allows safe HTML elements and attributes
// while stripping unsafe content like scripts and event handlers
//
// Parameters:
//   - html: The HTML string to sanitize
//
// Returns:
//   - string: The sanitized HTML with unsafe content removed
func SanitizeHTML(html string) string {
	return htmlSanitizer.Sanitize(html)
}
