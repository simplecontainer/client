package container

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/context"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func List(context *context.Context) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/control/container/list", context.ApiURL), http.MethodGet, nil)

	objects := make(map[string]*v1.ContainerDefinition)

	fmt.Println(string(response.Data))

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
			obj.Meta.Name,
		)
	}

	tbl.Print()
}
