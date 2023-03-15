/*
Copyright ©permaditya 2023 <permanaaditya2606@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

type player struct {
	DateOfBirth   string     `json:"Date of Birth"`
	Nationality   string     `json:"Nationality"`
	Club          string     `json:"club"`
	CompleteStats [][]string `json:"complete stats"`
	Height        string     `json:"height"`
	KeyStats      [][]string `json:"key_stats"`
	Name          string     `json:"name"`
	Position      string     `json:"position"`
}

// playerCmd represents the player command
var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Find player by name",
	Long: `Find player by name in the Premier League,
	premierleagueCLI find player <name of the player>`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := fmt.Sprintf("https://pl.apir7.repl.co/players/%s", name)

		// Make HTTP GET request to the API
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		// Decode the response JSON into a player struct
		var p player
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the player information
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Club: %s\n", p.Club)
		fmt.Printf("Position: %s\n", p.Position)
		fmt.Printf("Nationality: %s\n", p.Nationality)
		fmt.Printf("Height: %s\n", p.Height)

		// Create a table for complete stats
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Text: "Statistic"},
				{Text: "Value"},
			},
		}
		for _, stat := range p.CompleteStats {
			r := []*simpletable.Cell{
				{Text: stat[0]},
				{Text: stat[1]},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleUnicode)
		fmt.Println("Complete Stats:")
		table.Println()
	},
}

func init() {
	findCmd.AddCommand(playerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
