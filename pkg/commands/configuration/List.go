package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
)

func List(context *context.Context) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/configuration/List", context.ApiURL), nil)

	objects := make(map[string]*v1.Configuration)

	bytes, err := json.Marshal(response.Data)

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	err = json.Unmarshal(bytes, &objects)

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("GROUP", "NAME")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, obj := range objects {
		tbl.AddRow(
			obj.Meta.Group,
			obj.Meta.Identifier,
		)
	}

	tbl.Print()
}
