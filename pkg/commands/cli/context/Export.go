package context

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"log"
	"os"
)

func Export(contextName string, ctx *context.Context, rootDir string, API string) {
	var dirs []string
	entries, err := os.ReadDir(ctx.Directory)
	if err != nil {
		log.Fatal(err)
	}

	var validCtxProvided = false
	for _, e := range entries {
		if e.Name() == ".active" {
			continue
		}

		if contextName != "" {
			if contextName == e.Name() {
				validCtxProvided = true
			}
		}

		dirs = append(dirs, e.Name())
	}

	if contextName != "" {
		if validCtxProvided {
			ctx.SetActiveContext(contextName)
			ctx.LoadContext()

			var encrypted string
			encrypted, _, err = ctx.Export(API)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(encrypted)
		} else {
			fmt.Println(fmt.Sprintf("context %s does not exist", contextName))
		}
	} else {
		if ctx != nil {
			var encrypted string
			encrypted, _, err = ctx.Export(API)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(encrypted)
		}
	}
}
