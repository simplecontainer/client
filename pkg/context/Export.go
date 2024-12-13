package context

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func (context *Context) Export(API string) (string, string, error) {
	for block, rest := pem.Decode([]byte(context.CertBundle)); block != nil; block, rest = pem.Decode(rest) {
		switch block.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if cert.IsCA {
				pem.Encode(context.Ca, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: cert.Raw,
				})
			} else {
				pem.Encode(context.Cert, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: cert.Raw,
				})
			}

		case "PRIVATE KEY":
			pem.Encode(context.PrivateKey, &pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: block.Bytes,
			})

		default:
			return "", "", errors.New("unknown pem type in the cert bundle")
		}
	}

	context.ApiURL = API

	bytes, err := json.Marshal(context)

	if err != nil {
		panic(err)
	}

	compressed := compress(bytes)

	randbytes := make([]byte, 32)
	if _, err = rand.Read(randbytes); err != nil {
		return "", "", err
	}

	key := hex.EncodeToString(randbytes)

	contextPath := fmt.Sprintf("%s/%s.key", context.Directory, context.Name)
	err = os.WriteFile(contextPath, []byte(key), 0600)
	if err != nil {
		return "", "", err
	}

	encrypted, err := encrypt(string(compressed.Bytes()), key)

	if err != nil {
		return "", "", err
	}

	return encrypted, key, nil
}
