/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// color
const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

// List of options after command
var (
	provinsi string
	kabKota  string
)

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "cari",
	Short: "Menampilkan Jadwal Imsakiyah Berdasarkan Wilayah",
	Long:  `Menggunakan teknik web scraping untuk menampilkan jadwal-imsakiyah berdasarkan wilayah `,
	// Run the collect command
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("https://www.viva.co.id/jadwal-imsakiyah/%s/%s", provinsi, kabKota)

		// Make HTTP GET request
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the table element with class "table-jadwal"
		table := doc.Find("table.table-jadwal")

		// Find all the rows in the table's body
		rows := table.Find("tbody").Find("tr")

		// Create a new table object
		st := simpletable.New()

		// Add the table header
		header := []*simpletable.Cell{
			{Text: "No."},
			{Text: "Tanggal"},
			{Text: red("Imsak")},
			{Text: "Subuh"},
			{Text: "Zuhur"},
			{Text: "Asar"},
			{Text: blue("Maghrib")},
			{Text: "Isya"},
		}
		st.Header = &simpletable.Header{
			Cells: header,
		}

		// Iterate through the rows and add them to the table
		var rowCounter int = 1
		rows.Each(func(i int, row *goquery.Selection) {
			cols := row.Find("td")
			data := []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", rowCounter)},
				{Text: cols.Eq(1).Text()},
				{Text: red(cols.Eq(2).Text())},
				{Text: cols.Eq(3).Text()},
				{Text: cols.Eq(4).Text()},
				{Text: cols.Eq(5).Text()},
				{Text: blue(cols.Eq(6).Text())},
				{Text: cols.Eq(7).Text()},
			}
			st.Body.Cells = append(st.Body.Cells, data)
			rowCounter++
		})

		// Print the table
		st.SetStyle(simpletable.StyleUnicode)
		fmt.Println(st.String())
	},
}

func init() {
	rootCmd.AddCommand(collectCmd)

	// Add flags to specify provinsi and kabKota
	collectCmd.Flags().StringVarP(&provinsi, "provinsi", "p", "", "provinsi name")
	collectCmd.Flags().StringVarP(&kabKota, "kabKota", "k", "", "kabKota name")

	// Mark flags as required
	collectCmd.MarkFlagRequired("provinsi")
	collectCmd.MarkFlagRequired("kabKota")
}

// color format

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
