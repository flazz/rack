package filescommands

import (
	"github.com/codegangsta/cli"
	"github.com/jrperritt/rack/commands/filescommands/containercommands"
	"github.com/jrperritt/rack/commands/filescommands/objectcommands"
)

// Get returns all the commands allowed for a `files` request.
func Get() []cli.Command {
	return []cli.Command{
		{
			Name:        "container",
			Usage:       "Used for Cloud Files container operations",
			Subcommands: containercommands.Get(),
		},
		{
			Name:        "object",
			Usage:       "Used for Cloud Files object operations",
			Subcommands: objectcommands.Get(),
		},
	}
}
