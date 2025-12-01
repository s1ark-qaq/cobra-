package cmd

import (
	"github.com/spf13/cobra"
	docd "github.com/spf13/cobra/doc"
)

func init() {
	RootCmd.AddCommand(doc)
}

var doc = &cobra.Command{
	Use:   "doc",
	Short: "Generate documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return docd.GenMarkdownTree(RootCmd, "./docs")
	},
}
