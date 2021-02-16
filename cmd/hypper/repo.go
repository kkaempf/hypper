package main

import (
	"github.com/Masterminds/log-go"

	logio "github.com/Masterminds/log-go/io"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/cmd/helm/require"
)

var repoHypper = `
This command consists of multiple subcommands to interact with chart repositories.
It can be used to add, remove, list, and index chart repositories.
`

func newRepoCmd(logger log.Logger) *cobra.Command {
	wInfo := logio.NewWriter(logger, log.InfoLevel)
	cmd := &cobra.Command{
		Use:   "repo add|remove|list|index|update [ARGS]",
		Short: "add, list, remove, update, and index chart repositories",
		Long:  repoHypper,
		Args:  require.NoArgs,
	}

	cmd.AddCommand(
		newRepoAddCmd(wInfo),
	)

	return cmd
}
