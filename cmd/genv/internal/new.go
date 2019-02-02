package internal

import (
	"context"

	"github.com/gobuffalo/genny"
	"github.com/spf13/cobra"

	"github.com/timraymond/genv/pathgen"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "new generates new Go projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		gen := pathgen.New("changeme", "timraymond")
		run.With(gen)
		return run.Run()
	},
}
