package gssc

import (
	"fmt"
	"github.com/spf13/cobra"
	"gssc/pkg/gssc"
	"strconv"
)

var authorCmd = &cobra.Command{
	Use:     "author",
	Aliases: []string{"yazar"},
	Short:   "Verilen yazar bilgilerini listeler",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := gssc.GetAuthor(args[0])
		if len(res.Author.Nick) > 0 {
			fmt.Println("Yazar Detaylari : " + res.Author.Nick)
			fmt.Println("Nesil : " + res.Author.Genre)
			fmt.Println("Karma : " + res.Author.Karma)
			if res.Author.Online {
				fmt.Println("Cevrimici : Evet")
			} else {
				fmt.Println("Cevrimici : Hayir")
			}
			fmt.Println("Durumu : " + res.Author.Status)
			fmt.Println("Son Entry'si : " + res.Author.LastEntry)
			fmt.Println("Entry bilgileri : Bugun (" + strconv.Itoa(res.EntryStats.Today) + "), Bu Hafta (" + strconv.Itoa(res.EntryStats.Week) + "), Bu ay (" + strconv.Itoa(res.EntryStats.Month) + "), Toplan entry (" + strconv.Itoa(res.EntryStats.Total) + ")")
			//fmt.Println("Biraz daha detay : ")
			//resEntry := gssc.GetAuthorEntryStats(args[0])
		} else {
			fmt.Println("Boyle bir yazar yok galiba")
		}
	},
}

func init() {
	rootCmd.AddCommand(authorCmd)
}
