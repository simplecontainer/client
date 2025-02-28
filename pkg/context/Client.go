package context

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
	"os"
)

func (context *Context) GenerateHttpClient(CertBundle []byte) (*http.Client, error) {
	for block, rest := pem.Decode(CertBundle); block != nil; block, rest = pem.Decode(rest) {
		switch block.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, err
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

		case "EC PRIVATE KEY":
			pem.Encode(context.PrivateKey, &pem.Block{
				Type:  "EC PRIVATE KEY",
				Bytes: block.Bytes,
			})

		default:
			return nil, errors.New("unknown pem type in the cert bundle")
		}
	}

	context.CertBundle = string(CertBundle)

	cert, err := tls.X509KeyPair(context.Cert.Bytes(), context.PrivateKey.Bytes())
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(context.Ca.Bytes())

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}, nil
}

func (context *Context) ConnectionTest() bool {
	if context.Client == nil {
		fmt.Println("no active context found - please add least one context")
		os.Exit(1)
	}

	response := network.Send(context.Client, fmt.Sprintf("%s/connect", context.ApiURL), http.MethodGet, nil)

	if response != nil && response.HttpStatus == http.StatusOK {
		return true
	} else {
		return false
	}
}
