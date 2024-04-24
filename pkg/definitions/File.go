package definitions

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadFile(filePath string) string {
	var jsonData []byte = nil

	if filePath != "" {
		YAML, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("file does not exist")
			os.Exit(1)
		}

		var body interface{}
		if err := yaml.Unmarshal([]byte(YAML), &body); err != nil {
			panic(err)
		}

		body = convert(body)

		if jsonData, err = json.Marshal(body); err != nil {
			panic(err)
		}
	}

	return string(jsonData)
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
