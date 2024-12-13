package helpers

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func Confirm(message string) bool {
	ask := promptui.Select{
		Label: fmt.Sprintf("%s [y/n]", message),
		Items: []string{"y", "n"},
	}

	_, result, err := ask.Run()
	if err != nil {
		// if err provide simple yes no
		return false
	}

	return result == "y"
}
