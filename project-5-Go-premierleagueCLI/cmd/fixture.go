/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// fixtureCmd represents the fixture command
var fixtureCmd = &cobra.Command{
	Use:   "fixture",
	Short: "Get the list of fixtures",
	Long: `Get the list of fixtures in the Premier League,
premierleagueCLI fixture`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://pl.apir7.repl.co/fixtures"

		// Make HTTP GET request to the API
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		// Decode the response JSON into a slice of strings
		var fixtures []string
		err = json.NewDecoder(resp.Body).Decode(&fixtures)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create a new table
		table := simpletable.New()

		// Set the table headers
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Fixture"},
			},
		}

		// Add the fixture data to the table
		for _, fixture := range fixtures {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: fixture},
			}
			table.Body.Cells = append(table.Body.Cells, row)
		}

		// Set the table style
		table.SetStyle(simpletable.StyleCompactLite)

		// Print the table
		fmt.Println(table.String())
	},
}

func init() {
	findCmd.AddCommand(fixtureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixtureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixtureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
