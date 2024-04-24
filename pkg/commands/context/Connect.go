package context

import (
	"fmt"
	context2 "github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func Connect(URL string, CertBundlePath string, projectDir string) {
	context := context2.NewContext(projectDir)
	context.ApiURL = URL

	logger.Log.Info("trying to read cert bundle", zap.String("file", CertBundlePath))

	CertBundle, err := os.ReadFile(CertBundlePath)
	if err != nil {
		logger.Log.Info("certbundle file not found", zap.String("error", err.Error()))
		return
	}

	context.Client, err = context.GenerateHttpClient(CertBundle)

	if err != nil {
		logger.Log.Info("failed to generate http client", zap.String("error", err.Error()))
		return
	}

	resp, err := context.Client.Get(fmt.Sprintf("%s/healthz", context.ApiURL))

	if err != nil {
		logger.Log.Info("failed to connect to the smr-agent", zap.String("error", err.Error()))
		return
	}

	if resp.StatusCode == http.StatusOK {
		if context.SaveToFile(projectDir) {
			logger.Log.Info("authenticated against the smr-agent")
		}
	} else {
		logger.Log.Fatal("failed to authenticate against the smr-agent")
	}
}
