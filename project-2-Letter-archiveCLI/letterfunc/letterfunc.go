package letterfunc

import (
  "encoding/json"
  "fmt"
  "os"
  "errors"
)

// Letter represents a letter with its properties
type Letter struct {
	No           int    `json:"no"`
	Date         string `json:"date"`
	InOut        string `json:"in_out"`
	Registration string `json:"registration"`
	Sender       string `json:"sender"`
	Recipient    string `json:"recipient"`
	FileName     string `json:"file_name"`
}

// Letters represents a list of letters
type Letters []Letter


const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func ReadLettersFromJSONFile(file string) (Letters, error) {
	var letters Letters
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return letters, nil
		}
		return letters, err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&letters)
	return letters, err
}

func WriteLettersToJSONFile(letters Letters, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(letters)
}

func EditLetter(letters []Letter, no int, content string, fileFlag *string) error {
	index := -1
	for i, letter := range letters {
		if letter.No == no {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Letter not found")
		return errors.New("Letter not found")
	} else {
		return WriteLettersToJSONFile(letters, *fileFlag)
	}
}


//color styling functions

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func Gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
