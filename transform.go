package main

import "strings"

type FigureDto struct {
	Name         string
	Materials    []string
	URL          string
	BlogLinks    []string
	Brand        string
	Scale        string
	Measurements string
	Painter      string
	Sculptor     string
	ReleaseDates []Date
	ImageLinks   []string
	Prices       []Price
	Characters   []string
	Series       string
}

type Date struct {
	Year  string
	Month string
}

type Price struct {
	WithTax    string
	WithoutTax string
	Edition    string
}

func normalizeFigureData(figureData FigureData) FigureDto {
	return FigureDto{
		Name:         figureData.Name,
		URL:          figureData.URL,
		Brand:        figureData.Brand,
		Scale:        getScale(figureData.TableData[4]),
		Measurements: getMeasurements(figureData.TableData[4]),
		Sculptor:     figureData.TableData[5],
		Painter:      figureData.TableData[6],
		Materials:    getMaterials(figureData.Material),
		BlogLinks:    getBlogLinks(figureData.BlogLinks),
		ReleaseDates: getReleaseDates(figureData.TableData[2]),
		Prices:       getPrices(figureData.TableData[3]),
		ImageLinks:   getImageLinks(figureData.Name, figureData.TableData[2]),
		Characters:   getCharacters(figureData.TableData[1]),
		Series:       getSeriesName(figureData.TableData[0]),
	}
}

func getScale(size string) string {
	if strings.Contains(size, "スケール") {
		return strings.TrimSpace(size[0:3])
	} else {
		return ""
	}
}

func getMeasurements(data string) string {
	// TODO: implement function logic here
	return ""
}

func getBlogLinks(links string) []string {
	// TODO: implement function logic here
	return []string{}
}

func getMaterials(materials string) []string {
	// TODO: implement function logic here
	return []string{}
}

func getReleaseDates(dates string) []Date {
	// TODO: implement function logic here
	return []Date{}
}

func getPrices(prices string) []Price {
	// TODO: implement function logic here
	return []Price{}
}

func getImageLinks(name string, date string) []string {
	// TODO: implement function logic here
	return []string{}
}

func getCharacters(data string) []string {
	// TODO: implement function logic here
	return []string{}
}

func getSeriesName(data string) string {
	// TODO: implement function logic here
	return ""
}
