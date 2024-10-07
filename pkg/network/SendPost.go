package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/smr/pkg/httpcontract"
	"io"
	"net/http"
)

func SendPost(client *http.Client, URL string, data map[string]any) *httpcontract.ResponseOperator {
	var req *http.Request
	var err error

	if len(data) > 0 {
		var marshaled []byte
		marshaled, err = json.Marshal(data)

		if err != nil {
			return &httpcontract.ResponseOperator{
				HttpStatus:       0,
				Explanation:      "failed to marshal data for sending request",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		req, err = http.NewRequest("POST", URL, bytes.NewBuffer(marshaled))

		if err != nil {
			return &httpcontract.ResponseOperator{
				HttpStatus:       0,
				Explanation:      "failed to create http request",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("POST", URL, nil)

		if err != nil {
			return &httpcontract.ResponseOperator{
				HttpStatus:       0,
				Explanation:      "failed to create http request",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)

	if err != nil {
		return &httpcontract.ResponseOperator{
			HttpStatus:       0,
			Explanation:      "failed to connect to the smr-agent",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return &httpcontract.ResponseOperator{
			HttpStatus:       resp.StatusCode,
			Explanation:      "invalid response from the smr-agent",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	if resp.StatusCode == http.StatusOK {
		var response httpcontract.ResponseOperator
		err = json.Unmarshal(body, &response)

		if err != nil {
			return &httpcontract.ResponseOperator{
				HttpStatus:       resp.StatusCode,
				Explanation:      "failed to unmarshal body response from smr-agent",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		return &response
	} else {
		return &httpcontract.ResponseOperator{
			HttpStatus:       resp.StatusCode,
			Explanation:      string(body),
			ErrorExplanation: fmt.Sprintf("unexpected response from the server: %s", req.RequestURI),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}
}
