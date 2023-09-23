package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// Assume that all figures have been added at the start.
	// oldFigures := getFigures()
	// waitOneDay()

	oldFigures := []string{
		"/products/541/",
		"/products/540/",
		"/products/538/",
		"/products/534/",
		"/products/250/",
		"/products/539/",
		"/products/537/",
	}

	for true {
		newFigures := getFigures()

		// Compare the lists
		updatedFigures := findNewItems(newFigures, oldFigures)
		if len(updatedFigures) != 0 {
			for _, figure := range updatedFigures {
				fmt.Printf("New item: %s\n", figure)
				printFigure(figure, "Alter", getRootDir("2024"))
			}
		} else {
			fmt.Println(time.Now().String(), ": No new items found")
		}

		// Reassign
		oldFigures = newFigures
		waitOneDay()
	}
}

func getFigures() []string {
	c := colly.NewCollector()

	var figures []string
	c.OnHTML(".type-a", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, h *colly.HTMLElement) {
			figures = append(figures, h.Attr("href"))
		})
	})
	defer c.OnHTMLDetach(".type-a")

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

func waitOneDay() {
	time.Sleep(time.Duration(24) * time.Hour)
}

func printFigure(link string, brand string, dir string) {
	figure := getNewFigureData(link, brand, dir)

	// Print the fields
	fmt.Printf("Name: %s\n", figure.Name)
	fmt.Printf("Material: %s\n", figure.Material)

	// Print TableData
	fmt.Println("TableData:")
	for i, data := range figure.TableData {
		fmt.Printf("  Element %d: %s\n", i+1, data)
	}

	fmt.Printf("URL: %s\n", figure.URL)
	fmt.Printf("BlogLinks: %s\n", figure.BlogLinks)
	fmt.Printf("Brand: %s\n", figure.Brand)
}
