package cmd

import (
	"fmt"
	"github.com/asyard/gssc/pkg/gssc"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var agendaHeadersCmd = &cobra.Command{
	Use:     "gundem",
	Aliases: []string{"agenda"},
	Short:   "Gundem basliklarini (yeniden eskiye) listeler",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res := gssc.AgendaToday()

		fmt.Println("Gundem basliklari (" + res.Date + ")")
		fmt.Println("Baslik sayisi : " + strconv.Itoa(res.HeaderCount))

		var headerNamesSlice = make([]string, len(res.Headers))
		for i := 0; i < len(res.Headers); i++ {
			headerNamesSlice[i] = res.Headers[i].HeaderText + " (" + strconv.Itoa(res.Headers[i].Count) + ")"
		}

		prompt := promptui.Select{
			Label:    "Okunacak basligi secin",
			Items:    headerNamesSlice, //[]string
			Size:     10,
			HideHelp: true,
			//Templates: templates,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Bir seyler ters gitti muhtemelen\n")
			return
		}

		selected := result[0:strings.LastIndex(result, " (")]
		var selectedHeader = -1
		for _, v := range res.Headers {
			if v.HeaderText == selected {
				selectedHeader = v.HeaderID
				break
			}
		}
		fmt.Println(selectedHeader)
		headerEntriesToday := gssc.EntriesToday(selectedHeader)
		fmt.Println(headerEntriesToday)
		for _, v := range headerEntriesToday.Entryler {
			fmt.Println(v.Mesaj)
			fmt.Println(v.Tarih)
			fmt.Println(v.Yazar)
		}
	},
}

func init() {
	rootCmd.AddCommand(agendaHeadersCmd)
}
