package context

import (
	"bytes"
	"fmt"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/static"
	"go.uber.org/zap"
)

func NewContext(projectDir string) *Context {
	return &Context{
		ApiURL:        "",
		Name:          "",
		Directory:     fmt.Sprintf("%s/%s", projectDir, static.CLIENT_CONTEXT_DIR),
		CertBundle:    "",
		ActiveContext: "",
		PrivateKey:    &bytes.Buffer{},
		Cert:          &bytes.Buffer{},
		Ca:            &bytes.Buffer{},
	}
}

func (context *Context) LoadContext() *Context {
	if context.GetActiveContext() {
		if context.ReadFromFile() {
			var err error
			context.Client, err = context.GenerateHttpClient([]byte(context.CertBundle))

			if err != nil {
				logger.Log.Info("failed to generate http client", zap.String("error", err.Error()))
				return nil
			}

			return context
		}
	}

	return nil
}
