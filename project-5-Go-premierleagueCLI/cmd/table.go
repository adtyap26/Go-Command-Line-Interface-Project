/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strings"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Get Premier League table",
	Long:  `Get the current Premier League table`,
	Run: func(cmd *cobra.Command, args []string) {
		link := "https://onefootball.com/en/competition/premier-league-9/table"

		resp, err := http.Get(link)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		pageBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		pageString := string(pageBytes)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageString))
		if err != nil {
			panic(err)
		}

		var table []string

		doc.Find("a.standings__row-grid").Each(func(i int, s *goquery.Selection) {
			table = append(table, s.Text())
		})

		pprint(table)
	},
}

func init() {
	findCmd.AddCommand(tableCmd)
}

func pprint(table []string) {
	for _, row := range table {
		fmt.Println(row)
	}
}
