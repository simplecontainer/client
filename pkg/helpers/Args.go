package helpers

import (
	"fmt"
	"os"
)

func GrabArg(index int) string {
	if len(os.Args)-1 >= index {
		return os.Args[index]
	}

	fmt.Println("please provide arguments")
	os.Exit(1)
	return ""
}
