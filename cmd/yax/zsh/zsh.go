package zsh

import (
	"github.com/ympons/yax/cmd/yax/root"

	"github.com/spf13/cobra"
)

const example = `
	Add zsh configuration
	$ yax zsh

	Restore github configurations
	$ yax zsh restore
`

//
var Command = &cobra.Command{
	Use:     "zsh [<name>...]",
	Short:   "Yax zsh config",
	Example: example,
	RunE:    run,
}

func init() {
	root.Add(Command)
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}
