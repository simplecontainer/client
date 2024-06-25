package helpers

import (
	"os"
	"os/exec"
)

func TmpEditor(bytes []byte) ([]byte, error) {
	f, err := os.CreateTemp("", "edit")

	if err != nil {
		return nil, err
	}

	defer os.Remove(f.Name())

	_, err = f.Write(bytes)

	if err != nil {
		return nil, err
	}

	cmd := exec.Command("vi", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()

	if err != nil {
		return nil, err
	}

	err = cmd.Wait()

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(f.Name()) // just pass the file name
	if err != nil {
		return nil, err
	}

	return data, nil
}
