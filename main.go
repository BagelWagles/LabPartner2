package main

import (
	"fmt"
	"strings"

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
	// for _, val := range SeedLinks {
	// 	fmt.Println(val)
	// }
	// fmt.Println("----------")
	crawlPage("https://drexel.edu/medicine/about/departments/pharmacology-physiology/research/barker-lab/")
	// result2 := returnLinks(result1[0].NewLink)
	//https://www.uic.edu/
	// test := []string{"https://www.purdue.edu/"}
	// test := []string{"https://www.uic.edu/"}
	// fmt.Println("Result 1:")
	// result := crawlUic(test, 0, 2, []string{"school", "college", "departments", "of"}, []string{"paying", "for"}, SearchedSites)
	// fmt.Println("--------------")
	// fmt.Println("Result 2:")
	// // fmt.Println(result)

	// result2 := crawlUic(result, 0, 4, []string{"research", "lab", "departments", "program", "center"}, []string{"news"}, SearchedSites)
	// crawlUic(result2, 0, 4, []string{"professor,"principal investigator;"}, []string{}, SearchedSites)

	// fmt.Println("Done checking websites")

}

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
	labKeyTerms := []string{"principle investigator", "members", "lab", "publish"}
	searchKeyTerms := []string{"lab", "department", "center", "program"}
	badSearchTerms := []string{"news", "ethic", "saftey", "grant", "compliance"}
	var newLinks []string
	var siteScore int
	var linkScore int

	//Colly Scrapper
	c := colly.NewCollector()

	// Turns out colly has a .HasVisited() function no map need of searchlinks

	//This Function is called whenever a site is visted.
	//Goal: Check if lab and add to list,else search links
	c.OnResponse(func(e *colly.Response) {
		search = true
		newLinks = nil
		for _, term := range labKeyTerms {
			if strings.Contains(string(e.Body), term) {
				siteScore++
			}
		}
		for _, term := range labKeyTerms {
			if strings.Contains(string(e.Body), term) {
				siteScore--
			}
		}
		if siteScore > 1 {
			//For suspected Labs we save link, and dont search links.
			labLinks = append(labLinks, (e.Request.URL.RawPath))
			search = false
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		for _, term := range searchKeyTerms {
			if strings.Contains(e.Attr("href"), term) || strings.Contains(e.Request.AbsoluteURL(seedLink), term) {
				linkScore++
			}
		}
		for _, term := range badSearchTerms {
			if strings.Contains(e.Attr("href"), term) || strings.Contains(e.Request.AbsoluteURL(seedLink), term) {
				linkScore--
			}
		}
		if linkScore > 1 {
			newLinks = append(newLinks, e.Attr("href"))

		}
	})

	//returning LabsLink
	return labLinks
}
