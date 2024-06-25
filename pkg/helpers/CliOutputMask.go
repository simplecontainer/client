package helpers

func CliMask(condition bool, textTrue string, textFalse string) string {
	if condition {
		return textTrue
	} else {
		return textFalse
	}
}
