package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/smr/pkg/contracts"
	"io"
	"net/http"
	"os"
	"strings"
)

func SendRequest(client *http.Client, URL string, method string, data interface{}) *contracts.Response {
	var req *http.Request
	var err error

	if data != nil {
		var marshaled []byte
		marshaled, err = json.Marshal(data)

		switch v := data.(type) {
		case string:
			marshaled = []byte(v)
			break
		default:
			marshaled, err = json.Marshal(v)
		}

		if err != nil {
			return &contracts.Response{
				HttpStatus:       0,
				Explanation:      "failed to marshal data for sending request",
				ErrorExplanation: err.Error(),
				Error:            true,
				Success:          false,
				Data:             nil,
			}
		}

		req, err = http.NewRequest(method, URL, bytes.NewBuffer(marshaled))

		if err != nil {
			return &contracts.Response{
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
		req, err = http.NewRequest(method, URL, nil)

		if err != nil {
			return &contracts.Response{
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
		return &contracts.Response{
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
		return &contracts.Response{
			HttpStatus:       resp.StatusCode,
			Explanation:      "invalid response from the smr-agent",
			ErrorExplanation: err.Error(),
			Error:            true,
			Success:          false,
			Data:             nil,
		}
	}

	var response *contracts.Response
	err = json.Unmarshal(body, &response)

	if err != nil {
		var reqData []byte
		reqBody := req.Body

		if reqBody != nil {
			reqData, err = io.ReadAll(resp.Body)

			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("simplecontainer returned malformed response - debug information displayed")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println("Request URL: " + req.URL.String())
		fmt.Println("Request method: " + req.Method)
		fmt.Println("Request data: " + string(reqData))
		fmt.Println("Expected json but got: (server response shown below)")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println(string(body))

		os.Exit(1)
	}

	return response
}
