package cmd

import "github.com/spf13/cobra"

var packCmd = &cobra.Command{
	Short: "Pack file",
	Use:   "pack",
}

func init() {
	rootCmd.AddCommand(packCmd)
}
