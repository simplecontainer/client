package gitops

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/network"
	gitopsBase "github.com/simplecontainer/smr/implementations/gitops/gitops"
)

func List(context *context.Context) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops/List", context.ApiURL), nil)

	gitopsObj := make(map[string]*gitopsBase.Gitops)

	bytes, err := json.Marshal(response.Data)

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	err = json.Unmarshal(bytes, &gitopsObj)

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("GROUP", "NAME", "REPOSITORY", "REVISION", "SYNCED", "AUTO", "STATE")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, g := range gitopsObj {
		certRef := fmt.Sprintf("%s.%s", g.Auth.CertKeyRef.Group, g.Auth.CertKeyRef.Name)
		httpRef := fmt.Sprintf("%s.%s", g.Auth.HttpAuthRef.Group, g.Auth.HttpAuthRef.Name)

		if certRef == "." {
			certRef = ""
		}

		if httpRef == "." {
			httpRef = ""
		}

		if g.Definition == nil {
			continue
		}

		tbl.AddRow(g.Definition.Meta.Group,
			g.Definition.Meta.Name,
			helpers.CliMask(g.Commit.ID().IsZero(), fmt.Sprintf("%s (Not pulled)", g.RepoURL), fmt.Sprintf("%s (%s)", g.RepoURL, g.Commit.ID().String()[:7])),
			g.Revision,
			helpers.CliMask(g.Status.LastSyncedCommit.IsZero(), "Never synced", g.Status.LastSyncedCommit.String()[:7]),
			g.AutomaticSync,
			helpers.CliMask(g.Status.InSync, "InSync", "Drifted"),
		)
	}

	tbl.Print()
}
