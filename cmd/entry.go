package cmd

import (
	"fmt"
	"github.com/asyard/gssc/srv"
	"github.com/spf13/cobra"
)

var entryCmd = &cobra.Command{
	Use:     "girdi",
	Aliases: []string{"entry"},
	Short:   "Verilen id'ye ait entry'yi listeler",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := srv.GetEntry(args[0])
		fmt.Println(res.EntryText)
		fmt.Println("Yazar : " + res.Author)
		fmt.Println("Tarih : " + res.Date) //todo add modified
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
}
