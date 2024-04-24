package context

import (
	"bytes"
	"net/http"
)

type Context struct {
	ApiURL        string
	Name          string
	DirectoryPath string
	CertBundle    string
	Ca            *bytes.Buffer `json:"-"`
	Cert          *bytes.Buffer `json:"-"`
	PrivateKey    *bytes.Buffer `json:"-"`
	ActiveContext string        `json:"-"`
	Client        *http.Client  `json:"-"`
}
