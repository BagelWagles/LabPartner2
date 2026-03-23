package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

// "github.com/gocolly/colly"

var UniversityList []University
var Wikilinks []string
var SeedLinks []string
var searchDepth int = 5

// var searchTerms []string = []string{"center", "graduate", "branch", "research", "lab"}
var searchTerms []string = []string{"faculty", "staff"}

// var badTerms []string = []string{"fellowship", "syllabus", "collab", "ethic", "facebook", "news", "compliance", "section", "guides", "development", "patient", "visitor", "board"}
var badTerms []string = []string{}
var scoreMin int = 0

func main() {

	GenerateUniversities()

	// test := FindUrls()
	// goodLinks := TestLinks(test)
	// SeedLinks = CheckRobotTxT(goodLinks)

	// CrawlSeedLink(SeedLinks[0])

	// CrawlSeedList(SeedLinks)

	CrawlSeedLink("https://www.uic.edu/")

	fmt.Println("Done checking websites")

}

func CrawlSeedLink(seedLink string) {

	c := colly.NewCollector()
	c.AllowURLRevisit = true
	c.MaxDepth = 100
	var Pagelinks []string

	// Words to look for: Research, Lab, Study, Medicine,
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// loweredText := strings.ToLower(e.Attr("href"))
		loweredText := strings.Trim(strings.ToLower(e.Attr("href")), seedLink)
		searched := false
		hasRemoveTerms := false
		hasSearchTerm := false
		score := 0
		repeated := false
		for _, val := range Pagelinks {
			if strings.Contains(loweredText, val) {
				// fmt.Println("Removed: ", remove)
				repeated = true
			}
		}
		for _, remove := range badTerms {
			if strings.Contains(loweredText, remove) {
				// fmt.Println("Removed: ", remove)
				hasRemoveTerms = true
			}
		}
		for i, val := range searchTerms {
			if strings.Contains(loweredText, val) {
				score += i
			}

		}
		if strings.HasSuffix(loweredText, "lab") {
			score = 100
		}
		// fmt.Printf("\nsearched:%t, hasRemoveTerms: %t, hasSearchTerms: %t, Link:%s", searched, hasRemoveTerms, hasSearchTerm, loweredText)
		if !repeated && !searched && !hasRemoveTerms && hasSearchTerm {
			searched = true
			newlink, _ := strings.CutPrefix(e.Attr("href"), "/")
			Pagelinks = append(Pagelinks, newlink)
			fmt.Printf("score:%d, Site: %s \n", score, newlink)

			if strings.Contains(newlink, "https://") {
				c.Visit(newlink)
				fmt.Printf("score:%d, Site: %s \n", score, newlink)
			} else {

				c.Visit(seedLink + newlink)
				fmt.Printf("score:%d, Site: %s \n", score, seedLink+newlink)
			}
		}
		// if score == 100 {

		// }

	})

	// c.OnRequest(func(r *colly.Request) {
	// 	// fmt.Println("Visiting", r.URL.String())
	// })

	c.Visit(seedLink)
}

func CrawlSeedList(seedLinks []string) {
	var wg sync.WaitGroup

	for _, val := range seedLinks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CrawlSeedLink(val)
		}()
	}
	wg.Wait()
}
