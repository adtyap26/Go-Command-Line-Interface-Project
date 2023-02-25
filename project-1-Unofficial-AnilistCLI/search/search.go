package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
)

const query = `
query ($search: String) {
  Media (search: $search, type: ANIME) {
    title {
      romaji
      english
    }
    description
    averageScore
    startDate {
      year
      month
      day
    }
    studios {
      nodes {
        name
      }
    }
    popularity
    source
    genres
  }
}
`

// -- SearchAnime function starts here --

func SearchAnime(search string) {
	variables := map[string]interface{}{
		"search": search,
	}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	// Get the media data
	mediaData := data["data"].(map[string]interface{})["Media"]
	var animeData map[string]interface{}
	if media, ok := mediaData.([]interface{}); ok {
		animeData = media[0].(map[string]interface{})
	} else {
		animeData = mediaData.(map[string]interface{})
	}

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true

	fmt.Println("#####################################################################################################")
	table.AddRow("English Title:", animeData["title"].(map[string]interface{})["english"])
	table.AddRow("Romaji Title:", animeData["title"].(map[string]interface{})["romaji"])
	table.AddRow("", "")
	table.AddRow("Start Date:", fmt.Sprintf("%v-%v-%v", animeData["startDate"].(map[string]interface{})["year"], animeData["startDate"].(map[string]interface{})["month"], animeData["startDate"].(map[string]interface{})["day"]))

	studios := animeData["studios"].(map[string]interface{})["nodes"].([]interface{})
	var studioNames []string
	for _, studio := range studios {
		studioNames = append(studioNames, studio.(map[string]interface{})["name"].(string))
	}
	table.AddRow("Studios:", strings.Join(studioNames, ", "))

	table.AddRow("Popularity:", animeData["popularity"])

	table.AddRow("Source:", animeData["source"])
	var genres []string
	for _, genre := range animeData["genres"].([]interface{}) {
		genres = append(genres, genre.(string))
	}

	table.AddRow("Genres:", strings.Join(genres, ", "))
	table.AddRow("", "")

	description := animeData["description"].(string)
	description = strings.Replace(description, "<br>", "", -1)
	table.AddRow("Description:", description)

	table.AddRow("", "")
	table.AddRow("Ratings:", fmt.Sprintf("%v/100", animeData["averageScore"]))

	// Adding color to the table

	tableString := table.String()
	for _, row := range strings.Split(tableString, "\n") {
		columns := strings.SplitN(row, ":", 2)
		if len(columns) == 2 {
			fmt.Print(color.New(color.FgBlue).SprintFunc()(columns[0]) + ": ")
			fmt.Println(columns[1])
		} else {
			fmt.Println(row)
		}
	}
	// fmt.Println(table)
	fmt.Println("#####################################################################################################")

}

// --- TopTenAnime function starts here --

func TopTenAnime(year int) {
	variables := map[string]interface{}{
		"year": year,
	}

	query := `
	query ($year: Int) {
		Page (perPage: 10) {
			media (sort: POPULARITY_DESC, seasonYear: $year, type: ANIME) {
				title {
					romaji
					english
				}
				averageScore
				popularity
			}
		}
	}
	`
	requestBody, _ := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	// Get the media data
	mediaData := data["data"].(map[string]interface{})["Page"].(map[string]interface{})["media"]
	var animeData []map[string]interface{}
	for _, media := range mediaData.([]interface{}) {
		animeData = append(animeData, media.(map[string]interface{}))
	}

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true

	fmt.Println("#####################################################################################################")
	for i, anime := range animeData {
		table.AddRow(color.New(color.FgYellow, color.Bold).SprintfFunc()(fmt.Sprintf("%d. %s", i+1, anime["title"].(map[string]interface{})["romaji"])))
		// table.AddRow(fmt.Sprintf("%d. %s", i+1, anime["title"].(map[string]interface{})["romaji"]))
		// table.AddRow(fmt.Sprintf("%d.", i+1), anime["title"].(map[string]interface{})["romaji"])
		table.AddRow("English Title ", anime["title"].(map[string]interface{})["english"])
		table.AddRow("Popularity ", anime["popularity"])
		table.AddRow("Ratings ", fmt.Sprintf("%v/100", anime["averageScore"]))
		table.AddRow("", "")
	}

	// Adding color to the table

	// tableString := table.String()
	// rows := strings.Split(tableString, "\n")
	//
	//	for i, row := range rows {
	//	    columns := strings.SplitN(row, "", 2)
	//	    if i == 0 {
	//	        fmt.Println(color.New(color.FgYellow, color.Bold).SprintFunc()(columns[0]) + " " + columns[1])
	//	        continue
	//	    }
	//	    if len(columns) == 2 {
	//	        fmt.Print(color.New(color.FgBlue).SprintFunc()(columns[0]) + " ")
	//	        fmt.Println(columns[1])
	//	    } else {
	//	        fmt.Println(row)
	//	    }
	//	}
	tableString := table.String()
	for _, row := range strings.Split(tableString, "\n") {
		columns := strings.SplitN(row, "", 2)
		if len(columns) == 2 {
			fmt.Print(color.New(color.FgBlue).SprintFunc()(columns[0]) + " ")
			fmt.Println(columns[1])
		} else {
			fmt.Println(row)
		}
	}

	// fmt.Print(table)

	fmt.Println("#####################################################################################################")
}
