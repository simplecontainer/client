package helpers

import (
	"fmt"
	"github.com/docker/docker/client"
)

func GetDomainAndPort(URL string) (string, error) {
	url, err := client.ParseHostURL(URL)

	if err != nil {
		return "", err
	}

	connString := fmt.Sprintf("%s:%s", url.Hostname(), url.Port())
	return connString, nil
}
