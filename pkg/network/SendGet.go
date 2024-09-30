package network

import (
	"encoding/json"
	"github.com/simplecontainer/client/pkg/httpcontract"
	"github.com/simplecontainer/smr/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func SendGet(client *http.Client, URL string) *httpcontract.ResponseImplementation {
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
		var data *httpcontract.ResponseImplementation
		err := json.Unmarshal(body, &data)

		if err != nil {
			return nil
		}

		return data
	} else {
		logger.Log.Info("invalid response from the smr-agent", zap.String("status", resp.Status))
	}

	return nil
}
