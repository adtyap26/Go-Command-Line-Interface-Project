/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(defineCmd)
}

// var defineCmd = &cobra.Command{
// 	Use:   "define",
// 	Short: "Ambil definisi sebuah kata dari laman KBBI",
// 	Args:  cobra.MinimumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		word := strings.Join(args, " ")
// 		fmt.Printf("Searching for word: %s\n", word)
//
// 		// Define a new Colly collector
// 		c := colly.NewCollector()
//
// 		// Define a callback for when we receive a response from the KBBI website
// 		c.OnResponse(func(r *colly.Response) {
// 			fmt.Println("Received response from KBBI website")
// 		})
//
// 		  c.OnHTML("ul.adjusted-par li, ol li", func(e *colly.HTMLElement) {
//         definition := strings.TrimSpace(e.Text)
//         if definition != "" {
//             fmt.Printf("- %s\n", definition)
//         }
//     })
//
//     c.OnError(func(r *colly.Response, err error) {
//         log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
//     })
// 		// Visit the KBBI website with the specified word as the query parameter
// 		c.Visit(fmt.Sprintf("https://kbbi.kemdikbud.go.id/entri/%s", word))
// 	},
// }

var defineCmd = &cobra.Command{
	Use:   "cari [kata]",
	Short: "Mencari definis kata di dalam KBBI",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		word := args[0]
		url := fmt.Sprintf("https://kbbi.kemdikbud.go.id/entri/%s", word)

		fmt.Printf("Mencari definisi: %s\n", word)

		// Instantiate default collector
		c := colly.NewCollector()

		// Define a variable to hold the definitions
		var definitions []string

		// Scrape the definitions
		c.OnHTML("ul.adjusted-par li, ol li", func(e *colly.HTMLElement) {
			definition := strings.TrimSpace(e.Text)
			definitions = append(definitions, definition)
		})

		// Visit the main page
		err := c.Visit(url)
		if err != nil {
			return err
		}

		if len(definitions) == 0 {
			fmt.Printf("Tidak dapat menampilkan definisi '%s'\n", word)
		} else {
			// Create a new table
			table := simpletable.New()

			// Define the table headers
			header := &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "Definisi"},
				},
			}
			table.Header = header

			// Add the definitions to the table
			for _, definition := range definitions {
				row := []*simpletable.Cell{
					{Text: definition},
				}
				table.Body.Cells = append(table.Body.Cells, row)
			}

			// Print the table
			table.SetStyle(simpletable.StyleCompactLite)
			fmt.Println(table.String())
		}

		return nil
	},
}
