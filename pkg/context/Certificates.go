package context

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/smr/pkg/keys"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
	"os"
)

func (c *Context) ImportCertificates(key string) error {
	response := network.Send(c.Client, fmt.Sprintf("%s/fetch/certs", c.ApiURL), http.MethodGet, nil)

	if response.Success {
		keysEncrypted := keys.Encrypted{}
		bytes, err := response.Data.MarshalJSON()

		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, &keysEncrypted)

		if err != nil {
			return err
		}

		decrypted, _ := decrypt(keysEncrypted.Keys, key)

		var importedKeys keys.Keys
		err = json.Unmarshal([]byte(decrypted), &importedKeys)

		if err != nil {
			return err
		}

		err = importedKeys.CA.Write("/home/qdnqn/.ssh/simplecontainer")

		if err != nil {
			return err
		}

		for user, client := range importedKeys.Clients {
			err = client.Write(fmt.Sprintf("%s/.ssh/simplecontainer", os.Getenv("HOME")), user)
			if err != nil {
				return err
			}

			err = importedKeys.GeneratePemBundle(fmt.Sprintf("%s/.ssh/simplecontainer", os.Getenv("HOME")), user, importedKeys.Clients[user])

			if err != nil {
				fmt.Println(err)
			}
		}

		return nil
	} else {
		return errors.New(response.ErrorExplanation)
	}
}
