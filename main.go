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

const (
	NO_OF_RESULTS  = 1
	WRAP_WIDTH     = 80
	DISPLAY_FOOTER = true
)

func main() {
	queryFlag := flag.String("w", "", "Search query")
	nFlag := flag.Int("n", NO_OF_RESULTS, "Number of results to be displayed")
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

	var l int
	if *listSoundsFlag {
		l = len(res.Sounds)
		displaySoundFiles(res.Sounds, *nFlag)
	} else {
		l = len(res.Results)
		displayDefinitions(res.Results, *nFlag)
		if DISPLAY_FOOTER {
			fmt.Println(strings.Repeat("#", WRAP_WIDTH))
			fmt.Printf("Results Fetched: %d, Displayed: %d\n", l, min(l, *nFlag))
		}
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
		fmt.Println(strings.Repeat("#", WRAP_WIDTH))
		fmt.Println(d.Word)
		fmt.Printf("+1: %d\n", d.Upvote)
		fmt.Printf("-1: %d\n", d.Downvote)
		fmt.Println(wordwrap.WrapString(d.Definition, w))
		fmt.Println()
		fmt.Println(wordwrap.WrapString(d.Example, w))
	}
}

func min(x int, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
