package cmd

import (
	"os"

	"github.com/koki-develop/ddsk/internal/ddsk"
	"github.com/spf13/cobra"
)

var (
	flagColor bool
)

var rootCmd = &cobra.Command{
	Use:  "ddsk",
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		w := cmd.OutOrStdout()
		ddsk := ddsk.New(&ddsk.Config{Color: flagColor})

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

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flagColor, "color", "c", false, "colorize output")
}
