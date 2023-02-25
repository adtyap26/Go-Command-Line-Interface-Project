package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"
	"time"

	"CLIAPPS/search"
	// "CLIAPPS/layout"

	"github.com/fatih/color"
	"github.com/janeczku/go-spinner"
	"github.com/mgutz/ansi"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("#####################################################################################################")
	green := ansi.ColorCode("green+b")
	reset := ansi.ColorCode("reset")
	fmt.Println(green + `
      _          _ _     _     _    ____ _     ___ 
     / \   _ __ (_) |   (_)___| |_ / ___| |   |_ _|
    / _ \ | '_ \| | |   | / __| __| |   | |    | | 
   / ___ \| | | | | |___| \__ \ |_| |___| |___ | | 
  /_/   \_\_| |_|_|_____|_|___/\__|\____|_____|___|
                                                    
           "https://github.com/adtyap26"
` + reset)
	fmt.Print(color.New(color.FgHiYellow, color.Bold).SprintFunc()(" Looking for anime or the top ten list? \n 1) Search for anime.\n 2) See the top ten list.\n\n"))
	fmt.Printf(" ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choiceInt, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if choiceInt == 1 {
		fmt.Print(color.New(color.FgHiYellow, color.Bold).SprintFunc()(" What anime do you want to see? \n\n"))
		fmt.Printf(" ")
		searchQuery, _ := reader.ReadString('\n')
		searchQuery = strings.TrimSpace(searchQuery)
		s := spinner.StartNew("Processing...")
		s.SetSpeed(100 * time.Millisecond)
		// s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})

		s.SetCharset([]string{
			"▱▱▱▱▱",
			"▰▱▱▱▱",
			"▰▰▱▱▱",
			"▰▰▰▱▱",
			"▰▰▰▰▱",
			"▰▰▰▰▰",
			"▰▰▰▰▱",
			"▰▰▰▱▱",
			"▰▰▱▱▱",
			"▰▱▱▱▱",
			"▱▱▱▱▱",
		})

		time.Sleep(3 * time.Second)
		s.Stop()
		fmt.Println(color.New(color.FgGreen, color.Bold).SprintFunc()(" Found!\n"))
		search.SearchAnime(searchQuery)
	} else if choiceInt == 2 {
		fmt.Print(color.New(color.FgHiYellow, color.Bold).SprintFunc()(" What year do you want to see top ten anime? \n\n"))
		fmt.Printf(" ")
		year, _ := reader.ReadString('\n')
		year = strings.TrimSpace(year)
		s := spinner.StartNew("Processing...")
		s.SetSpeed(100 * time.Millisecond)
		// s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
		s.SetCharset([]string{
			"▱▱▱▱▱",
			"▰▱▱▱▱",
			"▰▰▱▱▱",
			"▰▰▰▱▱",
			"▰▰▰▰▱",
			"▰▰▰▰▰",
			"▰▰▰▰▱",
			"▰▰▰▱▱",
			"▰▰▱▱▱",
			"▰▱▱▱▱",
			"▱▱▱▱▱",
		})

		time.Sleep(3 * time.Second)
		s.Stop()
		fmt.Println(color.New(color.FgGreen, color.Bold).SprintFunc()(" Found!\n"))

		yearInt, err := strconv.Atoi(year)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		search.TopTenAnime(yearInt)
	}
}
