package main

import (
	"fmt"
	"slices"

	"github.com/gocolly/colly"
)

func main() {
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

    // In RAM for now...
    oldFigures := []string {
        "大鳳　春の暁に鳳歌うVer.",
        "ブレマートン　アクションクルーズVer.",
        "渡辺 曜",
        "津島 善子",
        "コルネリア",
        "ライダー／紫式部",
    }

    // Compare the lists
    var updatedFigures []string
    if (len(figures) != len(oldFigures)) {
        updatedFigures = findNewItems(figures, oldFigures)
    }
    for _, figure := range updatedFigures {
        fmt.Printf("New item: %s\n", figure)
    }
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
