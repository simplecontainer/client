package network

import (
	"encoding/json"
	"github.com/simplecontainer/smr/pkg/contracts"
	"io"
	"net/http"
)

func SendGet(client *http.Client, URL string) *contracts.ResponseImplementation {
	if client == nil {
		return &contracts.ResponseImplementation{
			HttpStatus:       0,
			Explanation:      "",
			ErrorExplanation: "client is invalid",
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	resp, err := client.Get(URL)

	if err != nil {
		return &contracts.ResponseImplementation{
			HttpStatus:       0,
			Explanation:      "",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &contracts.ResponseImplementation{
			HttpStatus:       resp.StatusCode,
			Explanation:      "",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	if resp.StatusCode == http.StatusOK {
		var data *contracts.ResponseImplementation
		err := json.Unmarshal(body, &data)

		if err != nil {
			return nil
		}

		return data
	} else {
		return &contracts.ResponseImplementation{
			HttpStatus:       resp.StatusCode,
			Explanation:      string(body),
			ErrorExplanation: "unexpected response from the server",
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}
}
