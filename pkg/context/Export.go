package context

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"software.sslmate.com/src/go-pkcs12"
)

func (context *Context) Export(API string) (string, string, error) {
	var cert *x509.Certificate
	var ca *x509.Certificate
	var tmp *x509.Certificate
	var err error

	for block, rest := pem.Decode([]byte(context.CertBundle)); block != nil; block, rest = pem.Decode(rest) {
		switch block.Type {
		case "CERTIFICATE":
			tmp, err = x509.ParseCertificate(block.Bytes)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if tmp.IsCA {
				ca = tmp
				pem.Encode(context.Ca, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: tmp.Raw,
				})
			} else {
				cert = tmp
				pem.Encode(context.Cert, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: tmp.Raw,
				})
			}

		case "EC PRIVATE KEY":
			pem.Encode(context.PrivateKey, &pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: block.Bytes,
			})

		default:
			return "", "", errors.New("unknown pem type in the cert bundle")
		}
	}

	x509KeyCert, err := tls.X509KeyPair(context.Cert.Bytes(), context.PrivateKey.Bytes())

	if err != nil {
		return "", "", err
	}

	password := ""
	pfxData, err := pkcs12.Modern.Encode(x509KeyCert.PrivateKey, cert, []*x509.Certificate{ca}, password)
	if err != nil {
		fmt.Println("Failed to create PKCS#12:", err)
		os.Exit(1)
	}

	context.ApiURL = API
	context.PKCS12 = b64.StdEncoding.EncodeToString(pfxData)

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
