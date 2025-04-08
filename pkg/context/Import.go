package context

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func (c *Context) Import(encrypted string, key string) error {
	ctx := NewContext(c.Directory)

	if key == "" {
		return errors.New("key is empty")
	}

	decrypted, err := decrypt(encrypted, key)

	if err != nil {
		return err
	}

	decompressed := decompress([]byte(decrypted))

	err = json.Unmarshal([]byte(decompressed), ctx)

	if err != nil {
		return err
	}

	ctx.Client, err = ctx.GenerateHttpClient([]byte(ctx.CertBundle))

	if err != nil {
		return err
	}

	if ctx.ConnectionTest() {
		viper.Set("y", true)
		err = ctx.SaveToFile()

		if err != nil {
			glog.Fatal(err.Error())
		}

		fmt.Println("Successfully imported context and connected to simplecontainer!")
	} else {
		fmt.Println(fmt.Sprintf("Failed to connect to the %s with imported context", ctx.ApiURL))
	}

	return nil
}
