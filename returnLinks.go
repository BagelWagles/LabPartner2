package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func returnLinks(seedLink string, SearchedSites map[string]string) []SearchLinks {

	c := colly.NewCollector()
	var links []SearchLinks
	// Words to look for: Research, Lab, Study, Medicine,
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		// fmt.Printf("%s -> %s\n", e.Text, e.Attr("href"))
		_, s := SearchedSites[e.Attr("href")]
		// fmt.Println("s: ", s, "->", e.Attr("href"))
		if !s {

			SearchedSites[e.Attr("href")] = e.Attr("href")
			add := new(SearchLinks)
			add.Context = e.Text
			add.NewLink = createLink(seedLink, e.Attr("href"))
			links = append(links, *add)
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(seedLink)
	return links
}
