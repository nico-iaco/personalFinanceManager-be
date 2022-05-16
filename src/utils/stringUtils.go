package utils

import "strings"

func SanitizeString(s string) string {
	escapedString := strings.Replace(s, "\n", "", -1)
	escapedString = strings.Replace(escapedString, "\r", "", -1)
	return escapedString
}
