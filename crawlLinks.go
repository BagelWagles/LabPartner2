package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func crawlLinks(seedLinks []string, i int, maxI int, search []string, remove []string, searched map[string]string) []string {
	// fmt.Println(seedLinks)
	var result []string

	if i == maxI {
		return nil
	}
	// var wg sync.WaitGroup
	for _, link := range seedLinks {
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()

		links := returnLinks(link, searched)
		Results := SortNewLinks(links, search, remove, 0)
		for _, newLink := range Results {
			// resp, err := http.Get(newLink.NewLink)

			// if err != nil {
			// 	fmt.Println(err)
			// }

			// if resp.StatusCode >= 200 && resp.StatusCode <= 299 /*&& 1000 > unsafe.Sizeof(resp.Body)*/ {

			// fmt.Println(resp.StatusCode)

			result = append(result, newLink.NewLink)
			// }
			// fmt.Println(newLink)
			// defer resp.Body.Close()
		}
		// }()

	}
	// wg.Wait()
	return append(result, crawlLinks(result, i+1, maxI, search, remove, searched)...)
}
func crawlPage(url string) {
	c := colly.NewCollector()

	defer c.Visit(url)

	c.OnResponse(func(e *colly.Response) {
		// fmt.Println(string(e.Body))
		count := strings.Count(string(e.Body), "Principal Investigator")
		if count > 0 {
			fmt.Printf("%v is likely a Lab as Principle Investigator is mentioned %v times", url, count)
		}

	})
}
