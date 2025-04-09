package context

import (
	"github.com/golang/glog"
	context "github.com/simplecontainer/client/pkg/context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func Connect(URL string, CertBundlePath string, rootDir string) error {
	ctx := context.NewContext(rootDir)
	ctx.ApiURL = URL

	glog.Info("trying to read cert bundle", zap.String("file", CertBundlePath))

	CertBundle, err := os.ReadFile(CertBundlePath)
	if err != nil {
		return err
	}

	ctx.Client, err = ctx.GenerateHttpClient(CertBundle)

	if err != nil {
		return err
	}

	err = ctx.Connect(viper.GetBool("wait"))

	if err != nil {
		return err
	}

	return ctx.SaveToFile()
}
