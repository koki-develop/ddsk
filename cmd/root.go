package cmd

import (
	"os"

	"github.com/koki-develop/ddsk/internal/ddsk"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "ddsk",
	RunE: func(cmd *cobra.Command, args []string) error {
		w := cmd.OutOrStdout()
		ddsk := ddsk.New()

		if err := ddsk.Run(w); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
