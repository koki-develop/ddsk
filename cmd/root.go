package cmd

import (
	"os"

	"github.com/koki-develop/ddsk/internal/ddsk"
	"github.com/spf13/cobra"
)

var (
	flagColor   bool
	flagAnimate bool
)

var rootCmd = &cobra.Command{
	Use:   "ddsk",
	Short: "Love Injection",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		ddsk := ddsk.New(&ddsk.Config{
			Writer:  cmd.OutOrStdout(),
			Color:   flagColor,
			Animate: flagAnimate,
		})

		if err := ddsk.Run(); err != nil {
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
	rootCmd.PersistentFlags().BoolVarP(&flagAnimate, "animate", "a", false, "animate output")
}
