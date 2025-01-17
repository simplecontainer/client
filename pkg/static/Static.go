package static

import (
	_ "embed"
)

const CONFIGDIR string = "config"

var CLIENT_CONTEXT_DIR = "contexts"
var CLIENT_LOG_DIR = "logs"

var CLIENT_STRUCTURE = []string{
	CLIENT_CONTEXT_DIR,
	CLIENT_LOG_DIR,
}
