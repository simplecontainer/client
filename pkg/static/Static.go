package static

import (
	_ "embed"
)

const CONFIGDIR string = "config"

var CLIENT_CONTEXT_DIR = "contexts"

var CLIENT_STRUCTURE = []string{
	CLIENT_CONTEXT_DIR,
}

//go:embed resources/git/version
var SMR_VERSION string
