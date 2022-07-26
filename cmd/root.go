package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "gssc",
	Version: version,
	Short:   "gssc - a simple CLI to retrieve data from rerererarara.net",
	Long: `gssc is a simple, unofficial CLI which uses
    
https://rerererarara.net/api/`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

/*func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}*/

func Execute() {
	/*var cmdFound bool
	cmd := rootCmd.Commands()
	for _, a := range cmd {
		for _, b := range os.Args[1:] {
			if a.Name() == b || contains(a.Aliases, b) {
				cmdFound = true
				break
			}
		}
	}*/

	/*if !cmdFound {
		//args := append([]string{"today"}, os.Args[1:]...)
		//rootCmd.SetArgs(args)
	}*/

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}
