package helpers

import (
	"fmt"
	"os"
)

func ExitWithErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
