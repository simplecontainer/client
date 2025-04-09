package context

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/spf13/viper"
	"os"
)

func (c *Context) GetActiveContext() bool {
	activeContextPath := fmt.Sprintf("%s/%s", c.Directory, ".active")

	activeContext, err := os.ReadFile(activeContextPath)
	if err != nil {
		return false
	}

	if string(activeContext) == "" {
		activeContext = []byte("default")
	}

	c.ActiveContext = fmt.Sprintf("%s/%s", c.Directory, string(activeContext))
	return true
}

func (c *Context) SetActiveContext(contextName string) bool {
	activeContextPath := fmt.Sprintf("%s/%s", c.Directory, ".active")

	err := os.WriteFile(activeContextPath, []byte(contextName), 0755)
	if err != nil {
		glog.Fatal("active context file not saved", err.Error())
	}

	return true
}

func (c *Context) ReadFromFile() bool {
	activeContext, err := os.ReadFile(c.ActiveContext)
	if err != nil {
		glog.Info("active context file not found", err.Error())
		return false
	}

	if err = json.Unmarshal(activeContext, &c); err != nil {
		glog.Info("active context file not valid json", err.Error())
		return false
	}

	return true
}

func (c *Context) SaveToFile() error {
	if c.Name == "" {
		c.Name = viper.GetString("context")
	}

	if c.Name == "" {
		return errors.New("context name cannot be empty")
	}

	jsonData, err := json.Marshal(c)

	if err != nil {
		return err
	}

	contextPath := fmt.Sprintf("%s/%s", c.Directory, c.Name)

	if _, err = os.Stat(contextPath); err == nil {
		if viper.GetBool("y") || helpers.Confirm("Context with the same name already exists. Do you want to overwrite it?") {
			err = os.WriteFile(contextPath, jsonData, 0600)
			if err != nil {
				glog.Fatal("active context file not saved", err.Error())
			}

			activeContextPath := fmt.Sprintf("%s/%s", c.Directory, ".active")

			err = os.WriteFile(activeContextPath, []byte(c.Name), 0755)
			if err != nil {
				glog.Fatal("active context file not saved", err.Error())
			}

			return nil
		} else {
			return errors.New("action aborted")
		}
	} else {
		err = os.WriteFile(contextPath, jsonData, 0600)
		if err != nil {
			return err
		}

		activeContextPath := fmt.Sprintf("%s/%s", c.Directory, ".active")

		err = os.WriteFile(activeContextPath, []byte(c.Name), 0755)
		if err != nil {
			return err
		}

		return nil
	}
}
