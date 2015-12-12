package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	ud "github.com/shivammg/urbandictionary"
	"log"
	"os"
	"strings"
)

func main() {
	queryFlag := flag.String("w", "", "Search query")
	nFlag := flag.Int("n", 1, "Number of results to be displayed")
	listSoundsFlag := flag.Bool("s", false, "List sound files instead")
	flag.Parse()

	var query string
	if *queryFlag == "" {
		fmt.Print("Search string: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		query = strings.Trim(input, " \n")
	} else {
		query = *queryFlag
	}

	res, err := ud.Query(query)
	if err != nil {
		log.Println(err)
	}

	if *listSoundsFlag {
		displaySoundFiles(res.Sounds, *nFlag)
	} else {
		displayDefinitions(res.Results, *nFlag)
	}
}

func displaySoundFiles(sounds []string, n int) {
	for i, s := range sounds {
		if i >= n {
			break
		}
		fmt.Println(s)
	}
}

func displayDefinitions(r []ud.Result, n int) {
	// width
	var w uint = 80
	for i, d := range r {
		if i >= n {
			break
		}
		fmt.Println(strings.Repeat("#", int(w)))
		fmt.Println(d.Word)
		fmt.Printf("+1: %d\n", d.Upvote)
		fmt.Printf("-1: %d\n", d.Downvote)
		fmt.Println(wordwrap.WrapString(d.Definition, w))
		fmt.Println()
		fmt.Println(wordwrap.WrapString(d.Example, w))
	}
}
