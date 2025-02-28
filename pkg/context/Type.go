package context

import (
	"bytes"
	"net/http"
)

type Context struct {
	Directory     string
	ApiURL        string
	Name          string
	CertBundle    string
	PKCS12        string
	Ca            *bytes.Buffer `json:"-"`
	Cert          *bytes.Buffer `json:"-"`
	PrivateKey    *bytes.Buffer `json:"-"`
	ActiveContext string        `json:"-"`
	ActiveGroup   string        `json:"-"`
	Client        *http.Client  `json:"-"`
}
