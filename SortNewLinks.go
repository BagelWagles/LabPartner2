package main

import (
	"fmt"
	"strings"
)

/*
	Search the links and the context. Ideally sort
 	into by a score like quantity and value of each word
	like for research > staff
*/

func SortNewLinks(links []SearchLinks, searchTerms []string, remove []string, minScore int) []SearchLinks {
	var result []SearchLinks
	for i, val := range links {
		_, s := SearchedSites[val.NewLink]
		// fmt.Printf("%t -> %s \n", s, val.NewLink)
		if !s {

			var LScore int
			for _, term := range searchTerms {
				if strings.Contains(strings.ToLower(val.Context), term) {
					LScore += 1
				}
			}
			for _, term := range remove {
				if strings.Contains(strings.ToLower(val.Context), term) {
					LScore -= 1
				}
			}
			if LScore > minScore {
				fmt.Printf("goodlink: %s \n", val.NewLink)

				result = append(result, links[i])
			}
		}

	}
	return result
}
