package root

import (
	"github.com/spf13/cobra"
)

//
var Command = &cobra.Command{
	Use:           "yax",
	SilenceErrors: true,
	SilenceUsage:  true,
}

//
func Add(cmd *cobra.Command) {
	Command.AddCommand(cmd)
}
