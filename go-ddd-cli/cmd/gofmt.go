package cmd

import (
	"github.com/spf13/cobra"

	"jettjia/igo-ddd-cli/util"
)

// gofmtCmd represents the gofmt command
var gofmtCmd = &cobra.Command{
	Use:   "gofmt",
	Short: "gofmt your project",
	Long:  `gofmt your project`,
	Run: func(cmd *cobra.Command, args []string) {
		util.GoFmt("./")
	},
}

func init() {
	rootCmd.AddCommand(gofmtCmd)
}
