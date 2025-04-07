package static

import (
	_ "embed"
	"github.com/simplecontainer/smr/pkg/static"
)

var ClientContextDir = "contexts"
var ClientLogDir = "logs"
var ClientConfigDir = static.CONFIGDIR

var ClientStructure = []string{
	ClientContextDir,
	ClientLogDir,
	ClientConfigDir,
}
