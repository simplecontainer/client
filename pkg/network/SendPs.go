package network

import (
	"encoding/json"
	"github.com/simplecontainer/smr/implementations/container/container"
	"github.com/simplecontainer/smr/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func SendPs(client *http.Client, URL string) map[string]map[string]*container.Container {
	resp, err := client.Get(URL)

	if err != nil {
		logger.Log.Info("failed to connect to the smr-agent", zap.String("error", err.Error()))
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Info("invalid response from the smr-agent", zap.String("error", err.Error()))
	}

	if resp.StatusCode == http.StatusOK {
		var data map[string]map[string]*container.Container
		json.Unmarshal(body, &data)

		return data
	} else {
		logger.Log.Info("invalid response from the smr-agent", zap.String("status", resp.Status))
	}

	return nil
}
