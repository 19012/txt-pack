package cmd

import "github.com/spf13/cobra"

var unpackCmd = &cobra.Command{
	Short: "Unpack file",
	Use:   "unpack",
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
