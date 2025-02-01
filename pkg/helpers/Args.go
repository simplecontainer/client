package helpers

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/smr/pkg/contracts"
	"github.com/simplecontainer/smr/pkg/f"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
	"strings"
)

func GrabArg(index int) string {
	if len(os.Args)-1 >= index {
		return os.Args[index]
	}

	fmt.Println("please provide arguments")
	os.Exit(1)
	return ""
}

func BuildFormat(arg string, group string) (contracts.Format, error) {
	// Build proper format from arg based on info provided
	// Default to prefix=simplecontainer.io, category=kind if missing

	var format contracts.Format
	var err error

	split := strings.Split(arg, "/")

	switch len(split) {
	case 1:
		// kind/name -> read group from flag!
		format, err = f.NewFromString(fmt.Sprintf("%s/%s/%s/%s", static.SMR_PREFIX, "kind", split[0], group))
		break
	case 2:
		// kind/name -> read group from flag!
		format, err = f.NewFromString(fmt.Sprintf("%s/%s/%s/%s/%s", static.SMR_PREFIX, "kind", split[0], group, split[1]))
		break
	case 3:
		// kind/group/name -> read group from arg
		format, err = f.NewFromString(fmt.Sprintf("%s/%s/%s/%s/%s", static.SMR_PREFIX, "kind", split[0], split[1], split[2]))
		break
	case 4:
		// category/kind/group/name
		format, err = f.NewFromString(fmt.Sprintf("%s/%s/%s/%s/%s", static.SMR_PREFIX, split[0], split[1], split[2], split[3]))
		break
	case 5:
		// prefix/category/kind/group/name
		format, err = f.NewFromString(fmt.Sprintf("%s/%s/%s/%s/%s", split[0], split[1], split[2], split[3], split[4]))
		break
	default:
		err = errors.New("valid formats are: [prefix/category/kind/group/name, category/kind/group/name, kind/group/name, kind/name, kind]")
	}

	return format, err
}
