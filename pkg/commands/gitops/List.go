package gitops

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
	gitopsSmrAgent "github.com/qdnqn/smr/pkg/gitops"
	"github.com/rodaine/table"
)

func List(context *context.Context) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops/List", context.ApiURL), nil)

	gitops := make(map[string]*gitopsSmrAgent.Gitops)

	bytes, err := json.Marshal(response.Data)

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	err = json.Unmarshal(bytes, &gitops)

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Group", "Name", "Repository", "Revision", "Last Synced Commit", "Automatic Sync", "Pooling Interval", "Ssh", "Http Auth")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, g := range gitops {
		lastSyncedCommit := g.LastSyncedCommit.String()

		if g.LastSyncedCommit.IsZero() {
			lastSyncedCommit = "Never synced"
		}

		certRef := fmt.Sprintf("%s.%s", g.CertKeyRef.Group, g.CertKeyRef.Identifier)
		httpRef := fmt.Sprintf("%s.%s", g.HttpAuthRef.Group, g.HttpAuthRef.Identifier)

		if certRef == "." {
			certRef = ""
		}

		if httpRef == "." {
			httpRef = ""
		}

		tbl.AddRow(g.Definition.Meta.Group,
			g.Definition.Meta.Identifier,
			g.RepoURL,
			g.Revision,
			lastSyncedCommit,
			g.AutomaticSync,
			g.PoolingInterval,
			certRef,
			httpRef)
	}

	tbl.Print()
}
