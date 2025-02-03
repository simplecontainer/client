package static

import (
	_ "embed"
)

var ClientContextDir = "contexts"
var ClientLogDir = "logs"

var ClientStructure = []string{
	ClientContextDir,
	ClientLogDir,
}
