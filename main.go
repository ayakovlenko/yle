package main

import (
	"flag"
	"fmt"
	"os"
)

var limit int

func main() {

	if limit <= 0 {
		fmt.Println("-n must be bigger than 0")
		os.Exit(1)
	}

	articles, err := scrapeArticles()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, a := range articles {
		if i+1 > limit {
			break
		}

		fmt.Printf("%s\n\u001B[32m%s\u001B[0m\n%s\n\n",
			a.time.Format("January 02, 15:04"),
			a.title,
			a.url,
		)
	}
}

func init() {
	flag.IntVar(&limit, "n", 5, "number of articles to display")
	flag.Parse()
}
