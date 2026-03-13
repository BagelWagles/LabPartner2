package main

import (
	"fmt"
	"net/http"
	"strings"
)

func CheckRobotTxT(Links []string) []string {
	var results []string
	for _, link := range Links {
		go func() {
			resp, err := http.Get(fmt.Sprintf("%s/robots.txt", link))
			if err != nil {
				// fmt.Println("doesn't have a Robot.Txt")
			}
			// fmt.Println(resp.Status)

			// This whole if just removes allow the status that don't allow for scraping.
			if !(strings.EqualFold(string(resp.Status), "403 Forbidden") || strings.EqualFold(string(resp.Status), "502 Bad Gateway") || strings.EqualFold(string(resp.Status), "405 Method Not Allowed") || strings.EqualFold(string(resp.Status), "406 Not Acceptable")) {

				fmt.Println("Good")
				results = append(results, link)
			}
		}()

	}
	return results
}
