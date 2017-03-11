package github

import (
	"github.com/ympons/yax/cmd/yax/root"

	"github.com/spf13/cobra"
)

const example = `
	Add github configuration
	$ yax github

	Restore github configurations
	$ yax github restore
`

//
var Command = &cobra.Command{
	Use:     "github [<name>...]",
	Short:   "Yax github config",
	Example: example,
	RunE:    run,
}

func init() {
	root.Add(Command)
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}
