package context

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	context "github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func Connect(URL string, CertBundlePath string, rootDir string) error {
	ctx := context.NewContext(rootDir)
	ctx.ApiURL = URL

	logger.Log.Info("trying to read cert bundle", zap.String("file", CertBundlePath))

	CertBundle, err := os.ReadFile(CertBundlePath)
	if err != nil {
		return err
	}

	ctx.Client, err = ctx.GenerateHttpClient(CertBundle)

	if err != nil {
		return err
	}

	if viper.GetBool("wait") {
		err = backoff.Retry(func() error {
			var resp *http.Response
			resp, err = ctx.Client.Get(fmt.Sprintf("%s/healthz", ctx.ApiURL))

			if err != nil {
				return err
			} else {
				if resp.StatusCode == http.StatusOK {
					if ctx.SaveToFile() {
						return nil
					}
				} else {
					return errors.New("failed to authenticate against the smr-agent")
				}
			}

			return errors.New("context not saved")
		}, backoff.NewExponentialBackOff())

		if err != nil {
			return err
		}

		return nil
	} else {
		resp, err := ctx.Client.Get(fmt.Sprintf("%s/healthz", ctx.ApiURL))

		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			if ctx.SaveToFile() {
				return nil
			}

			return errors.New("failed to save context to contexts directory")
		} else {
			return errors.New("failed to authenticate against the smr-agent")
		}
	}
}
