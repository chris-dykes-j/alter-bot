package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

type FigureData struct {
	Name      string
	TableData []string
	Material  string
	URL       string
	BlogLinks string
	Brand     string
}

func getNewFigureData(link string, brand string, root string) FigureData {
	c := colly.NewCollector()
	var data FigureData

	c.OnHTML(".hl06", func(h *colly.HTMLElement) {
		data.Name = h.Text
	})

	data.URL = fmt.Sprintf("https://alter-web.jp%s", link)
	data.Brand = brand

	// Add figure images
	figureDir := filepath.Join(root, "temp")
	c.OnHTML(".item-mainimg figure img", func(h *colly.HTMLElement) {
		downloadImg(h.Attr("src"), figureDir, "1.jpg")
	})
	i := 1
	c.OnHTML(".imgset li img", func(h *colly.HTMLElement) {
		i++
		downloadImg(h.Attr("src"), figureDir, fmt.Sprintf("%d.jpg", i))
	})
	defer c.OnHTMLDetach(".item-mainimg figure img")
	defer c.OnHTMLDetach(".imgset li img")

	// Get Figure Table
	c.OnHTML(".tbl-01 > tbody", func(e *colly.HTMLElement) {
		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			text, err := el.DOM.Html()
            if err != nil {
                fmt.Printf("%s\n", err)
            }
            text = strings.ReplaceAll(text, "<br/>", " ")
			data.TableData = append(data.TableData, strings.Join(strings.Fields(text), " "))
		})
    })
	defer c.OnHTMLDetach(".tbl-01 > tbody")

	// Get Material
	c.OnHTML(".spec > .txt", func(h *colly.HTMLElement) {
		data.Material = strings.Join(strings.Fields(h.Text), " ")
	})
	defer c.OnHTMLDetach(".spec > .txt")

	// Get Blog links
	var blogLinks []string
	c.OnHTML(".imgtxt-type-b", func(h *colly.HTMLElement) {
		h.ForEach("a", func(_ int, el *colly.HTMLElement) {
			blogLink := el.Attr("href")
			blogLink = fmt.Sprintf("https://alter-web.jp%s", blogLink)
			blogLinks = append(blogLinks, blogLink)
		})
	})
	defer c.OnHTMLDetach(".imgtxt-type-b")

	url := fmt.Sprintf("https://alter-web.jp%s", link)
	err := c.Visit(url)
	if err != nil {
		fmt.Printf("Error: %s, %s\n", err, link)
	}

	// Add blogLinks
	data.BlogLinks = strings.Join(blogLinks, ",")

    newDir := filepath.Join(root, data.Name) 
    os.Rename(figureDir, newDir)

	return data
}
