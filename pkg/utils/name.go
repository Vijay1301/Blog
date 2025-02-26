package utils

import (
	"strings"
	"unicode"
)

func ExtractNameFromEmail(email string) (string, string) {
	parts := strings.Split(email, "@")
	if len(parts) < 2 || len(parts[0]) == 0 {
		return "", ""
	}

	nameParts := strings.FieldsFunc(parts[0], func(r rune) bool {
		return r == '.' || r == '_'
	})

	cleanedParts := []string{}
	for _, part := range nameParts {
		cleanedPart := RmoveNumbers(part)
		if cleanedPart != "" {
			cleanedParts = append(cleanedParts, cleanedPart)
		}
	}

	if len(cleanedParts) == 2 {
		return strings.Title(cleanedParts[0]), strings.Title(cleanedParts[1])
	} else if len(cleanedParts) == 1 {

		cleanedName := strings.Title(cleanedParts[0])
		return cleanedName, cleanedName
	}

	return "", ""
}

func RmoveNumbers(s string) string {
	var result strings.Builder
	for _, r := range s {
		if !unicode.IsDigit(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
