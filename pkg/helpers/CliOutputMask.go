package helpers

import "strings"

func CliMask(condition bool, textTrue string, textFalse string) string {
	if condition {
		return textTrue
	} else {
		return textFalse
	}
}

func CliRemoveComa(text string) string {
	str, _ := strings.CutSuffix(text, ", ")
	return str
}
