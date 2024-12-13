package context

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func (context *Context) GetActiveContext() bool {
	activeContextPath := fmt.Sprintf("%s/%s", context.Directory, ".active")

	activeContext, err := os.ReadFile(activeContextPath)
	if err != nil {
		return false
	}

	if string(activeContext) == "" {
		activeContext = []byte("default")
	}

	context.ActiveContext = fmt.Sprintf("%s/%s", context.Directory, string(activeContext))
	return true
}

func (context *Context) SetActiveContext(contextName string) bool {
	activeContextPath := fmt.Sprintf("%s/%s", context.Directory, ".active")

	err := os.WriteFile(activeContextPath, []byte(contextName), 0755)
	if err != nil {
		logger.Log.Fatal("active context file not saved", zap.String("error", err.Error()))
	}

	return true
}

func (context *Context) ReadFromFile() bool {
	activeContext, err := os.ReadFile(context.ActiveContext)
	if err != nil {
		logger.Log.Info("active context file not found", zap.String("error", err.Error()))
		return false
	}

	if err = json.Unmarshal(activeContext, &context); err != nil {
		logger.Log.Info("active context file not valid json", zap.String("error", err.Error()))
		return false
	}

	return true
}

func (context *Context) SaveToFile(projectDir string) bool {
	jsonData, err := json.Marshal(context)

	if err != nil {
		logger.Log.Fatal(err.Error())
		return false
	}

	contextPath := fmt.Sprintf("%s/%s", projectDir, context.Name)

	if _, err = os.Stat(contextPath); err == nil {
		if viper.GetBool("y") || helpers.Confirm("Context with the same name already exists. Do you want to overwrite it?") {
			err = os.WriteFile(contextPath, jsonData, 0600)
			if err != nil {
				logger.Log.Fatal("active context file not saved", zap.String("error", err.Error()))
			}

			activeContextPath := fmt.Sprintf("%s/%s", projectDir, ".active")

			err = os.WriteFile(activeContextPath, []byte(context.Name), 0755)
			if err != nil {
				logger.Log.Fatal("active context file not saved", zap.String("error", err.Error()))
			}

			return true
		} else {
			logger.Log.Info("action aborted")
			return false
		}
	} else {
		err = os.WriteFile(contextPath, jsonData, 0600)
		if err != nil {
			logger.Log.Fatal("active context file not saved", zap.String("error", err.Error()))
		}

		activeContextPath := fmt.Sprintf("%s/%s", projectDir, ".active")

		err = os.WriteFile(activeContextPath, []byte(context.Name), 0755)
		if err != nil {
			logger.Log.Fatal("active context file not saved", zap.String("error", err.Error()))
		}

		return true
	}
}
