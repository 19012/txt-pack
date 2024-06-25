package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const packedExtention = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

var vlcCmd = &cobra.Command{
	Short: "Pack file using variable-length code alg",
	Use:   "vlc",
	Run:   vlcPack,
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

func vlcPack(_cmd *cobra.Command, args []string) {
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

	fmt.Println(string(data)) // TODO: used Encode func{

	packed := ""
	if err := os.WriteFile(packedFilePath(filePath), []byte(packed), 0644); err != nil {
		handleError(err)
	}
}

func packedFilePath(filePath string) string {
	fileBase := filepath.Base(filePath)
	return strings.TrimSuffix(fileBase, filepath.Ext(fileBase)) + "." + packedExtention
}
