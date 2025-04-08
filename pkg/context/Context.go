package context

import (
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"github.com/simplecontainer/client/pkg/static"
	"go.uber.org/zap"
)

func NewContext(projectDir string) *Context {
	return &Context{
		ApiURL:        "",
		Name:          "",
		Directory:     fmt.Sprintf("%s/%s", projectDir, static.ClientContextDir),
		CertBundle:    "",
		ActiveContext: "",
		PrivateKey:    &bytes.Buffer{},
		Cert:          &bytes.Buffer{},
		Ca:            &bytes.Buffer{},
	}
}

func (c *Context) LoadContext() *Context {
	if c.GetActiveContext() {
		if c.ReadFromFile() {
			var err error
			c.Client, err = c.GenerateHttpClient([]byte(c.CertBundle))

			if err != nil {
				glog.Info("failed to generate http client", zap.String("error", err.Error()))
				return nil
			}

			return c
		}
	}

	return nil
}
