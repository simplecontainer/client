package context

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	context2 "github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/spf13/viper"
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

	if viper.GetBool("wait") {
		err = backoff.Retry(func() error {
			var resp *http.Response
			resp, err = context.Client.Get(fmt.Sprintf("%s/healthz", context.ApiURL))

			if err != nil {
				logger.Log.Info("failed to connect to the smr-agent, trying again....")
			} else {
				if resp.StatusCode == http.StatusOK {
					if context.SaveToFile(projectDir) {
						logger.Log.Info("authenticated against the smr-agent")
						return nil
					}
				} else {
					return errors.New("failed to authenticate against the smr-agent")
				}
			}

			return errors.New("context not saved")
		}, backoff.NewExponentialBackOff())

		if err != nil {
			fmt.Println(err)
		}
	} else {
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
}
