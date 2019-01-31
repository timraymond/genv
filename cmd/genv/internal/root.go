package internal

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var rootCmd = &cobra.Command{
	Use:   "genv",
	Short: "genv helps you manage GOPATHs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Please select a subcommand")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
