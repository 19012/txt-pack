package cmd

import (
	"19012/txt-pack/lib/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const unpackedExtention = "txt"

var vlcUnpackCmd = &cobra.Command{
	Short: "Unpack file using variable-length code alg",
	Use:   "vlc",
	Run:   vlcUnpack,
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}

func vlcUnpack(_cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	unpacked := vlc.Decode(string(data))

	if err := os.WriteFile(unpackedFilePath(filePath), []byte(unpacked), 0644); err != nil {
		handleError(err)
	}
}

func unpackedFilePath(filePath string) string {
	fileBase := filepath.Base(filePath)
	return strings.TrimSuffix(fileBase, filepath.Ext(fileBase)) + "." + unpackedExtention
}
