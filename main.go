package main

import (
	"bufio"
	"flag"
	"fmt"
	ud "github.com/shivammg/urbandictionary"
	"os"
	"strings"
)

func main() {
	w := flag.String("w", "", "Word or string to search")
	n := flag.Int("n", 1, "Number of results")
	flag.Parse()

	var q string
	if *w == "" {
		fmt.Print("Search String: ")
		reader := bufio.NewReader(os.Stdin)
		inp, _ := reader.ReadString('\n')
		q = strings.Trim(inp, " \n")
	} else {
		q = *w
	}

	res, _ := ud.Query(q)
	displayDefinitions(res.Results, *n)
}

func displayDefinitions(r []ud.Result, n int) {
	// width
	w := 80
	for i, d := range r {
		if i >= n {
			break
		}
		fmt.Println(strings.Repeat("#", w))
		fmt.Printf("+1: %d\n", d.Upvote)
		fmt.Printf("-1: %d\n", d.Downvote)
		printWrappedString(d.Definition, w)
		fmt.Println()
		printWrappedString(d.Example, w)
	}
}

func printWrappedString(st string, w int) {
	offset := w - len(st)%w
	l := len(st) + offset
	st = st + strings.Repeat(" ", offset)
	for i := 0; i < l; i += w {
		fmt.Println(st[i : i+w])
	}
}
