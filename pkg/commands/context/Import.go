package context

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
)

func Import(encrypted string, ctx *context.Context, rootDir string, key string) {
	err := ctx.Import(encrypted, key)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("context imported with success")
	}
}
