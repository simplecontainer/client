package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/qdnqn/smr/pkg/httpcontract"
	"github.com/qdnqn/smr/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Result struct {
	Data string `json:"data"`
}

func SendFile(client *http.Client, URL string, jsonData string) *httpcontract.ResponseImplementation {
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(jsonData)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Log.Info("invalid response from the smr-agent", zap.String("error", err.Error()))
	}

	var response httpcontract.ResponseImplementation
	err = json.Unmarshal(body, &response)

	if err != nil {
		logger.Log.Info("invalid response from the smr-agent", zap.String("error", err.Error()))
	}

	return &response
}
