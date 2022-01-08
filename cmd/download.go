package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"time"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:                   "download <url>",
	Short:                 "download from url passed",
	Long:                  `download from url passed as an argument`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var url string
		if args == nil || len(args) == 0 || args[0] == "" {
			fmt.Println("URL is mandatory!")
			return
		}
		url = args[0]
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error while hitting the url provided : %v", err)
			return
		}
		file, err := os.Create(fmt.Sprint("./downloads/file_", time.Now().UnixNano()))
		if err != nil {
			fmt.Printf("Error while creating file locally : %v", err)
			return
		}
		defer file.Close()
		defer resp.Body.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Printf("Error while copying contents to file : %v", err)
			return
		}
		return

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
