package network

import (
	"bytes"
	"encoding/json"
	"github.com/qdnqn/smr/pkg/implementations"
	"github.com/qdnqn/smr/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Result struct {
	Data string `json:"data"`
}

func SendFile(client *http.Client, URL string, jsonData string) *implementations.Response {
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Log.Info("invalid response from the smr-agent", zap.String("error", err.Error()))
	}

	if resp.StatusCode == http.StatusOK {
		var response implementations.Response
		err := json.Unmarshal(body, &response)

		if err != nil {
			logger.Log.Info("invalid response from the smr-agent", zap.String("error", err.Error()))
		}

		return &response
	} else {
		logger.Log.Info("invalid response from the smr-agent")
		return nil
	}
}
