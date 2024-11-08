package network

import (
	"fmt"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
	"github.com/simplecontainer/smr/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func SendLogs(client *http.Client, URL string) map[string]map[string]platforms.IContainer {
	var readBytes int
	resp, err := client.Get(URL)

	if err != nil {
		logger.Log.Info("failed to connect to the smr-agent", zap.String("error", err.Error()))
		return nil
	}

	buff := make([]byte, 512)

	for {
		readBytes, err = resp.Body.Read(buff)

		if readBytes == 0 || err == io.EOF {
			err = resp.Body.Close()

			if err != nil {
				return nil
			}

			fmt.Println("server closed connection")
			break
		}

		fmt.Print(string(buff[:readBytes]))
	}

	fmt.Print(string(buff[:readBytes]))

	return nil
}
