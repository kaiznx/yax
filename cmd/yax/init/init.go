package init

import (
	"github.com/spf13/cobra"
	"github.com/ympons/yax/cmd/yax/root"
)

//
var Command = &cobra.Command{
	Use:   "init",
	Short: "Initialize workspace",
	RunE:  run,
}

func init() {
	root.Add(Command)
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}
