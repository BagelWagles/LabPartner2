package main

import (
	"fmt"
	"os"
)

// "github.com/gocolly/colly"

var UniversityList []University
var Wikilinks []string
var SeedLinks []string

func main() {
	// c := colly.NewCollector(
	// 	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// 	colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	// )

	// fmt.Println(CollegeList)

	GenerateUniversities()

	// test := FindUrls()
	// goodLinks := TestLinks(test)
	// SeedLinks = CheckRobotTxT(goodLinks)
	Save(SeedLinks)
	fmt.Println("Done checking websites")
	// var r bufio.Reader
	//I am going to test If I can get the links by guess most of their links
	//Sample: https://www.harvard.edu/

}
func Save(s []string) {
	//save file address: //SaveFiles//Save
	file, err := os.Open("//SaveFiles//Save")

}
