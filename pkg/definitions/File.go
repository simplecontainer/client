package definitions

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"net/url"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	var bytes []byte = nil

	if filePath != "" {
		YAML, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var body interface{}
		if err = yaml.Unmarshal(YAML, &body); err != nil {
			return nil, err
		}

		body = convert(body)

		if bytes, err = json.Marshal(body); err != nil {
			return nil, err
		}
	}

	return bytes, nil
}

func DownloadFile(URL *url.URL) ([]byte, error) {
	path := fmt.Sprintf("/tmp/%s", b64.StdEncoding.EncodeToString([]byte(URL.String())))

	out, err := os.Create(path)
	defer out.Close()

	if err != nil {
		return nil, err
	}

	resp, err := http.Get(URL.String())
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("definition not found at specific URL")
		os.Exit(1)
	}

	return ReadFile(path)
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
