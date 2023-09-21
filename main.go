package main

import (
	"fmt"
	"slices"
    "time"

	"github.com/gocolly/colly"
)

func main() {
    // Assume that all figures have been added at the start.
    oldFigures := getNewFigures() 

    for true {
        newFigures := getNewFigures()

        // Compare the lists
        var updatedFigures []string
        updatedFigures = findNewItems(newFigures, oldFigures)
        if len(updatedFigures) != 0 {
            for _, figure := range updatedFigures {
                fmt.Printf("New item: %s\n", figure)
            }
        } else {
            fmt.Println("No new items found")
        }

        // Reassign
        oldFigures = newFigures

        // Wait until tomorrow to check again. Needs tweaks
        time.Sleep(time.Duration(24) * time.Hour)
    }
}

func getNewFigures() []string {
    c := colly.NewCollector()
    
    var figures []string
	c.OnHTML("figure", func(e *colly.HTMLElement) {
		e.ForEach("figcaption", func(i int, h *colly.HTMLElement) {
			figures = append(figures, h.Text)
		})
    })
    defer c.OnHTMLDetach("figure")

    url := "https://alter-web.jp/figure/"
    c.Visit(url)

    return figures 
}

func findNewItems(newList []string, oldList []string) []string {
    var result []string

    // Assumes newList is lager than oldList
    for _, item := range newList {
        if !slices.Contains(oldList, item) {
           result = append(result, item) 
        }
    }
    return result
}
