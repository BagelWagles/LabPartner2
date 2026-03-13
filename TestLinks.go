package main

import (
	// cls"fmt"

	"fmt"
	"net/http"
	"time"
)

var WorkingLinks []string

func TestLinks(links map[string]string) []string {

	for _, link := range links {

		go func() {

			resp, err := http.Get(link)
			if err != nil {
				// fmt.Println("This link doesn't work: ", link)
			}
			if resp != nil {
				fmt.Println(":) -> ", link)
				WorkingLinks = append(WorkingLinks, link)

			}

		}()

	}

	time.Sleep(10 * time.Second)
	fmt.Println("Terminated early")

	return WorkingLinks

}
