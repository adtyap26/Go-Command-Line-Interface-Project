package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
	"letter-archive/letterfunc"
)

func main() {
	// Define command line flags
	addFlag := flag.Bool("add", false, "add a new letter")
	deleteFlag := flag.Int("delete", 0, "delete a letter by number")
	listFlag := flag.Bool("list", false, "list all letters")
	fileFlag := flag.String("file", "letters.json", "JSON file to read/write letters")
	editFlag := flag.Int("edit", 0, "edit letter by number")

	// Parse command line arguments
	flag.Parse()

	// Read letters from JSON file
	letters, err := letterfunc.ReadLettersFromJSONFile(*fileFlag)
	if err != nil {
		fmt.Println("Error reading letters from JSON file:", err)
		os.Exit(1)
	}

	// Add a new letter
	if *addFlag {
		var letter letterfunc.Letter
		fmt.Print("No: ")
		fmt.Scanf("%d", &letter.No)
		fmt.Print("Date: ")
		fmt.Scanf("%s", &letter.Date)
		fmt.Print("In/Out: ")
		fmt.Scanf("%s", &letter.InOut)
		fmt.Print("Registration: ")
		fmt.Scanf("%s", &letter.Registration)
		fmt.Print("Sender: ")
		fmt.Scanf("%s", &letter.Sender)
		fmt.Print("Recipient: ")
		fmt.Scanf("%s", &letter.Recipient)
		fmt.Print("File name: ")
		fmt.Scanf("%s", &letter.FileName)
		letters = append(letters, letter)
		err = letterfunc.WriteLettersToJSONFile(letters, *fileFlag)
		if err != nil {
			fmt.Println("Error writing letters to JSON file:", err)
			os.Exit(1)
		}
	}

	// Delete a letter

	if *deleteFlag > 0 {
		index := -1
		for i, letter := range letters {
			if letter.No == *deleteFlag {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("Letter not found")
		} else {
			letters = append(letters[:index], letters[index+1:]...)
			// Update the "no" field of the remaining letters
			for i := range letters {
				letters[i].No = i + 1
			}
			err = letterfunc.WriteLettersToJSONFile(letters, *fileFlag)
			if err != nil {
				fmt.Println("Error writing letters to JSON file:", err)
				os.Exit(1)
			}
		}
	}

	// edit letter
	if *editFlag > 0 {
		if len(os.Args) < 4 {
			fmt.Println("Error: Please specify which parts of the letter you want to edit.")
			os.Exit(1)
		}

		letterIndex := *editFlag - 1
		if letterIndex >= len(letters) {
			fmt.Printf("Error: Letter with index %d does not exist.\n", letterIndex+1)
			os.Exit(1)
		}

		letter := letters[letterIndex]

		partsToEdit := os.Args[3:]
		for _, part := range partsToEdit {
			switch part {
			case "Date":
				fmt.Print("Date: ")
				fmt.Scanf("%s", &letter.Date)
			case "In/Out":
				fmt.Print("In/Out: ")
				fmt.Scanf("%s", &letter.InOut)
			case "Registration":
				fmt.Print("Registration: ")
				fmt.Scanf("%s", &letter.Registration)
			case "Sender":
				fmt.Print("Sender: ")
				fmt.Scanf("%s", &letter.Sender)
			case "Recipient":
				fmt.Print("Recipient: ")
				fmt.Scanf("%s", &letter.Recipient)
			case "Filename":
				fmt.Print("Filename: ")
				fmt.Scanf("%s", &letter.FileName)
			default:
				fmt.Printf("Error: Invalid part '%s' to edit.\n", part)
				os.Exit(1)
			}
		}

		letters[letterIndex] = letter
		err = letterfunc.WriteLettersToJSONFile(letters, *fileFlag)
		if err != nil {
			fmt.Println("Error writing letters to JSON file:", err)
			os.Exit(1)
		}
	}

	// List all letters
	if *listFlag {
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("No")},
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("Date")},
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("In/Out")},
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("Registration")},
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("Sender")},
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("Recipient")},
				{Align: simpletable.AlignCenter, Text: letterfunc.Blue("Name Of File")},
			},
		}
		for _, letter := range letters {
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", letter.No)},
				{Text: letter.Date},
				{Text: letter.InOut},
				{Text: letter.Registration},
				{Text: letter.Sender},
				{Text: letter.Recipient},
				{Text: letter.FileName},
			})
		}
		table.SetStyle(simpletable.StyleUnicode)
		fmt.Println(table.String())
	}
}
