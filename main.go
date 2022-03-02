package main

import (
	"fmt"
	"os"
	"pw_lab2/internal/pwtcp"
)

func main() {
	cl := pwtcp.NewClient()
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Printf("No option specified. Use -h for help\n")
		return
	}

	fmt.Println(arguments)

	switch arguments[0] {
	case "-s":
		if len(arguments) != 2 {
			fmt.Printf("Not enough parameters. Expected 1 search term (string)\n")
			break
		}

		// Search here with arguments[0] return plain string and error
		searchResult, err := cl.SearchGoogle(arguments[1])
		if err != nil {
			fmt.Printf("Search on %s failed: %v\n", arguments[1], err)
		} else {
			fmt.Printf("%s\n", searchResult)
		}
	case "-u":
		if len(arguments) != 2 {
			fmt.Printf("Not enough parameters. Expected 1 url (string)\n")
			break
		}

		// Get here with arguments[0] return plain string and error
		getResult, err := cl.Get(arguments[1])
		if err != nil {
			fmt.Printf("Get request on %s failed: %v\n", arguments[1], err)
		} else {
			fmt.Printf("%s\n", getResult)
		}
	case "-h":
		fmt.Printf("go2web -s <<search term>> - get top 10 results from search engine by search term\n")
		fmt.Printf("go2web -u \"<<url link>>\"- get the plain text of the webpage of the input url link")
	}
}
