package helpers

import (
	"fmt"
	"io"
)

func PrintBytes(reader io.ReadCloser) error {
	buff := make([]byte, 512)

	for {
		bytes, err := reader.Read(buff)

		if bytes == 0 || err == io.EOF {
			err = reader.Close()

			if err != nil {
				return err
			}

			fmt.Print(string(buff[:bytes]))
			break
		}

		fmt.Print(string(buff[:bytes]))
	}

	return nil
}
