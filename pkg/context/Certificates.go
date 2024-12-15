package context

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/simplecontainer/smr/pkg/keys"
	"github.com/simplecontainer/smr/pkg/static"
	"net/http"
	"os"
)

func (context *Context) ImportCertificates(key string) error {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/ca", context.ApiURL), http.MethodGet, nil)

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

		err = importedKeys.Server.Write("/home/qdnqn/.ssh/simplecontainer")

		if err != nil {
			return err
		}

		for user, client := range importedKeys.Clients {
			err = client.Write(fmt.Sprintf("%s/.ssh/simplecontainer", os.Getenv("HOME")), user)
			if err != nil {
				return err
			}

			importedKeys.GeneratePemBundle(static.SMR_SSH_HOME, "root", importedKeys.Clients[user])
		}

		return nil
	} else {
		return errors.New(response.ErrorExplanation)
	}
}
