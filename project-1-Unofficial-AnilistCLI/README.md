# Unofficial AniList CLI

This is a simple command-line application that allows you to search for anime on AniList and display information about them. The application uses the AniList GraphQL API to fetch the data.

## Installation

### Prerequisites

You must have Go installed on your system to be able to build and run this application. If you don't have Go installed, you can download it from the official website [https://go.dev/dl/](https://go.dev/dl/)

https://user-images.githubusercontent.com/101618848/212522094-3727bf6a-eaed-40f7-b8cf-c0327c31f4bd.mp4

### Windows

1. Install Go by running the installer and following the prompts.
2. Open the Command Prompt and navigate to the directory where you want to clone the repository
3. navigate to the project directory
4. Build the application by running the following command:

```
go build
```

5. Run the application by running the following command:

```
Unofficial-AniList-CLI.exe
```

### Linux and MacOS

1. Install Go by following the instructions on the official website [https://go.dev/doc/install](https://go.dev/doc/install)
2. Open the Terminal and navigate to the directory where you want to clone the repository
3. navigate to the project directory
4. Build the application by running the following command:

```
go build
```

5. Run the application by running the following command:

```
./Unofficial-AniListCLI
```

## Usage

To use the application, simply run the binary file in your terminal and you will be prompted with the following:

Looking for anime or the top ten list?

1.  Search for anime
2.  See the top ten list.

Type the title of the anime you want to see and hit enter. The application will then display the description and ratings of the anime.

## libraries

This application is developed in Go and uses the AniList GraphQL API to fetch the data. The project uses the following libraries:

- bufio
- bytes
- encoding/json
- fmt
- net/http
- os

- github.com/gosuri/uitable
- github.com/janeczku/go-spinner
- github.com/fatih/color
- github.com/mgutz/ansi

## Disclaimer

This is an unofficial command-line application for AniList and is not affiliated with AniList or its creators in any way.
