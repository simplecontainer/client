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

func ReadFile(filePath string) (string, error) {
	var jsonData []byte = nil

	if filePath != "" {
		YAML, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}

		var body interface{}
		if err = yaml.Unmarshal([]byte(YAML), &body); err != nil {
			return "", err
		}

		body = convert(body)

		if jsonData, err = json.Marshal(body); err != nil {
			return "", err
		}
	}

	return string(jsonData), nil
}

func DownloadFile(URL *url.URL) (string, error) {
	path := fmt.Sprintf("/tmp/%s", b64.StdEncoding.EncodeToString([]byte(URL.String())))

	out, err := os.Create(path)
	defer out.Close()

	if err != nil {
		return "", err
	}

	resp, err := http.Get(URL.String())
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return "", err
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
