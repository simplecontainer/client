package network

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/smr/pkg/contracts"
	"io"
	"net/http"
)

func SendDelete(client *http.Client, URL string) *contracts.ResponseOperator {
	var req *http.Request
	var err error

	req, err = http.NewRequest("DELETE", URL, nil)

	if err != nil {
		return &contracts.ResponseOperator{
			HttpStatus:       0,
			Explanation:      "failed to create http request",
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

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return &contracts.ResponseOperator{
			HttpStatus:       resp.StatusCode,
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
		return &contracts.ResponseOperator{
			HttpStatus:       resp.StatusCode,
			Explanation:      string(body),
			ErrorExplanation: fmt.Sprintf("unexpected response from the server: %s", req.RequestURI),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}
}
