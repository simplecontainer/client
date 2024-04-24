package context

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/logger"
	"log"
	"os"
)

func Switch(contextName string, context *context.Context) {
	var dirs []string
	entries, err := os.ReadDir(context.DirectoryPath)
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
			fmt.Println(fmt.Sprintf("active context is %s", contextName))
		} else {
			fmt.Println(fmt.Sprintf("context %s does not exist", contextName))
		}
	} else {
		prompt := promptui.Select{
			Label: "Select a context",
			Items: dirs,
		}

		_, result, err := prompt.Run()

		if err != nil {
			logger.Log.Fatal("failed to select from list of contexts")
		}

		context.SetActiveContext(result)
		fmt.Println(fmt.Sprintf("active context is %s", result))
	}
}
