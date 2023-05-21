package version

import (
	"os"

	"github.com/gernest/vince/pkg/version"
	"github.com/urfave/cli/v3"
)

func Version() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "prints version information",
		Action: func(ctx *cli.Context) error {
			x := version.Build()
			os.Stdout.WriteString(x.String())
			return nil
		},
	}
}
