package main

import (
	// "fmt"
	// "fmt"
	"strings"
)

func createLink(baseLink string, new string) string {
	if strings.Contains(new, "https://") {

		return new
	} else {
		if strings.Contains(baseLink, new) {
			s := strings.Split(baseLink, "/")
			var base string
			for i, val := range s {
				if i < 3 {
					base += val + "/"
				}
			}
			// fmt.Println("base:", base)
			return base + new

		} else {
			return baseLink + new
		}

	}
	// return baseLink + new
}
