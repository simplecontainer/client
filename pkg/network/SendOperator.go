package network

import (
	"bytes"
	"encoding/json"
	"github.com/simplecontainer/smr/pkg/contracts"
	"io"
	"net/http"
)

func SendOperator(client *http.Client, URL string, data map[string]any) *contracts.ResponseOperator {
	var req *http.Request
	var err error

	if len(data) > 0 {
		marshaled, err := json.Marshal(data)

		if err != nil {
			return &contracts.ResponseOperator{
				HttpStatus:       0,
				Explanation:      "failed to marshal data for sending request",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		req, err = http.NewRequest("POST", URL, bytes.NewBuffer(marshaled))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", URL, nil)
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return &contracts.ResponseOperator{
			HttpStatus:       0,
			Explanation:      "failed to craft request",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		return &contracts.ResponseOperator{
			HttpStatus:       0,
			Explanation:      "failed to connect to the smr-agent",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &contracts.ResponseOperator{
			HttpStatus:       0,
			Explanation:      "invalid response from the smr-agent",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	if resp.StatusCode == http.StatusOK {

		var response contracts.ResponseOperator
		err = json.Unmarshal(body, &response)

		if err != nil {
			return &contracts.ResponseOperator{
				HttpStatus:       0,
				Explanation:      "failed to unmarshal body response from smr-agent",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		return &response
	} else {
		return &contracts.ResponseOperator{
			HttpStatus:       resp.StatusCode,
			Explanation:      string(body),
			ErrorExplanation: "unexpected response from the server",
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}
}
