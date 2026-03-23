package main

import (
	"fmt"
	"strings"
)

func FindUrls() map[string]string {

	var result = make(map[string]string)
	for _, uni := range UniversityList {

		title := strings.Trim(uni.university_Name, "University")
		title = strings.Trim(title, "College")
		title = strings.Trim(title, "  ")
		title = strings.Trim(title, "of")
		title = strings.ReplaceAll(title, " .edu", ".edu")
		title = strings.ReplaceAll(title, " ", "_")
		title = strings.ToLower(title)
		// GuessLink1 := fmt.Sprintf("https://www.%s.edu/", title)
		var GuessLink2 string
		sep := strings.Split(strings.ToLower(uni.university_Name), " ")
		for _, word := range sep {
			WholeWord := []rune(word)

			GuessLink2 = GuessLink2 + string(WholeWord[0])

		}
		GuessLink2 = fmt.Sprintf("https://www.%s.edu", GuessLink2)
		// result = append(result, GuessLink1)

		result[fmt.Sprintf("https://www.%s.edu", sep[0])] = fmt.Sprintf("https://www.%s.edu", sep[0])
		_, repeat := result[GuessLink2]
		if !repeat {
			result[GuessLink2] = GuessLink2
		}

	}

	return result
}
