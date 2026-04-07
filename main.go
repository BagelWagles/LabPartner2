package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// "github.com/gocolly/colly"

var UniversityList []University
var Wikilinks []string
var SeedLinks []string
var SearchedSites map[string]string
var searchDepth int = 5

func main() {
	// SearchedSites := make(map[string]string)

	// GenerateUniversities()
	// TestTime = 10
	// test := FindUrls()
	// goodLinks := TestLinks(test)
	// SeedLinks = CheckRobotTxT(goodLinks)

	//https://www.uic.edu/
	CrawlUni("https://www.uic.edu/")
	// fmt.Println(CrawlUni("https://alsberglab.weebly.com/"))

	// fmt.Println("Done checking websites")

}

/*
New Design Idea...
1.Seed links
2.Check if pages are Labs ie(contains the word Principle Investigator)
2.5. If page is a Lab site add to final Lab links array
3.If not Get all links that look like could help find the labs ie(link or context contains(lab,research,program,department))
4.Get rid of links that contain(News,blog, other stuff.)

*/

/*
Need to make manager Function that manages running through school sites.
Inputs: SeedLinks
Outputs: LabLink
*/
func CrawlUni(seedLink string) []string {
	//Anouncing Start of function
	fmt.Printf("Searching: %v", seedLink)

	//function Variables
	var labLinks []string
	var search bool
	labKeyTerms := []string{"principal investigator", "members", "lab", "publish"}
	searchKeyTerms := []string{"research", "lab", "department", "center"}
	badSearchTerms := []string{"news", "ethic", "saftey", "grant", "compliance", "wellness"}
	var newLinks []string
	var siteScore int
	var linkScore int

	//Colly Scrapper
	c := colly.NewCollector()
	c.MaxDepth = 6

	// Turns out colly has a .HasVisited() function no map need of searchlinks

	//This Function is called whenever a site is visted.
	//Goal: Check if lab and add to list,else search links
	c.OnResponse(func(e *colly.Response) {
		fmt.Println("-----------------------")
		fmt.Printf("Searching: %v\n", e.Request.AbsoluteURL(e.Request.URL.Path))

		search = true

		for _, term := range labKeyTerms {
			if strings.Contains(string(e.Body), term) {
				siteScore++
			}
		}
		if strings.Contains(string(e.Body), labKeyTerms[0]) {
			siteScore += 10
		}
		for _, term := range labKeyTerms {
			if strings.Contains(string(e.Body), term) {
				siteScore--
			}
		}
		if siteScore > 9 {
			//For suspected Labs we save link, and dont search links.
			labLinks = append(labLinks, (e.Request.URL.RawPath))
			search = false
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// fmt.Printf("a[href]: %v ,AbsoluteUrl(seedlink): %v, AbsoluteURL(e.attr(href)): %v \n", e.Attr("href"), e.Request.AbsoluteURL(seedLink), e.Request.AbsoluteURL(e.Attr("href")))
		linkScore = 0
		tempLink := e.Request.AbsoluteURL(e.Attr("href"))
		for _, term := range searchKeyTerms {
			if strings.Contains(tempLink, term) {
				// fmt.Println(tempLink, "|", term)
				linkScore++
			}
		}
		for _, term := range badSearchTerms {
			if strings.Contains(tempLink, term) {
				linkScore += -3
			}
		}
		b, _ := c.HasVisited(tempLink)
		if linkScore >= 1 && !b {
			fmt.Printf("New Link:%v, Score: %v\n", tempLink, linkScore)
			newLinks = append(newLinks, tempLink)
		}
	})

	//Once onHtml is done this function goes through and visit all the new links
	c.OnScraped(func(e *colly.Response) {
		if search {
			// This won't work because it will simply overide new links after the first one.
			//maybe just add a list of all intresting links and go through them like a heap or stack

			time.Sleep(250 * time.Millisecond)
			if len(newLinks) > 0 {
				e.Request.Visit(newLinks[0])
				newLinks = newLinks[1:]
			}
			//removes link ounce it has been visited

		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request failed:", r.Request.URL, "Error:", err)
	})

	//Init Page Search
	c.Visit(seedLink)
	//returning LabsLink
	return labLinks
}
