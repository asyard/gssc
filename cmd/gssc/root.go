package gssc

import (
	"fmt"
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Beklenmedik bir seyler oldu, calismadi ya")
		os.Exit(1)
	}
}
