package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

// Could be refactored for easy maintenance, but maintenance seems unlikely and it works.
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
		url := fmt.Sprintf("https://alter-web.jp%s", h.Attr("src"))
		downloadImg(url, figureDir, "1.jpg")
	})
	i := 1
	c.OnHTML(".imgset li img", func(h *colly.HTMLElement) {
		i++
		url := fmt.Sprintf("https://alter-web.jp%s", h.Attr("src"))
		downloadImg(url, figureDir, fmt.Sprintf("%d.jpg", i))
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

func getRootDir(year string) string {
	root := os.Getenv("FIGURE_IMAGE_ROOT")
	if root == "" {
		log.Fatal("FIGURE_IMAGE_ROOT environment variable not set")
	}

	parentDir := filepath.Join(root, year)
	if _, err := os.Stat(parentDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(parentDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Made %s directory\n", parentDir)
	}

	return parentDir
}

// Generalized function
func downloadImg(url string, directory string, imageName string) {
	// Ensure the directory exists
	if err := os.MkdirAll(directory, 0755); err != nil {
		log.Fatal(err)
	}

	// Get image bytes
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Make image file
	imagePath := filepath.Join(directory, imageName)
	imageFile, err := os.Create(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	// Save bytes to image file
	_, err = io.Copy(imageFile, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Saved image %s to %s\n", imageName, directory)
}
