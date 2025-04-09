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

func (c *Context) GenerateHttpClient(CertBundle []byte) (*http.Client, error) {
	for block, rest := pem.Decode(CertBundle); block != nil; block, rest = pem.Decode(rest) {
		switch block.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, err
			}

			if cert.IsCA {
				pem.Encode(c.Ca, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: cert.Raw,
				})
			} else {
				pem.Encode(c.Cert, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: cert.Raw,
				})
			}

		case "EC PRIVATE KEY":
			pem.Encode(c.PrivateKey, &pem.Block{
				Type:  "EC PRIVATE KEY",
				Bytes: block.Bytes,
			})

		default:
			return nil, errors.New("unknown pem type in the cert bundle")
		}
	}

	c.CertBundle = string(CertBundle)

	cert, err := tls.X509KeyPair(c.Cert.Bytes(), c.PrivateKey.Bytes())
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(c.Ca.Bytes())

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}, nil
}

func (c *Context) ConnectionTest() bool {
	if c.Client == nil {
		fmt.Println("no active context found - please add least one context")
		os.Exit(1)
	}

	response := network.Send(c.Client, fmt.Sprintf("%s/connect", c.ApiURL), http.MethodGet, nil)

	if response != nil && response.HttpStatus == http.StatusOK {
		return true
	} else {
		return false
	}
}
