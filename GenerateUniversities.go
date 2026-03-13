package main

import "strings"

func GenerateUniversities() {
	/*
	* Goes through list of R1 Colleges
	* Generates a unvisity class for each
	* Each contains: Name of Uni, City of campus, State
	* Creates a Wiki Search Link to be used to find college websites.
	 */
	CollegeList := ReadToString("ReadFiles\\ace-institutional-classifications.uid")

	for _, val := range strings.Split(CollegeList, "\n") {
		if len(val) > 0 {
			SeperatedVals := strings.Split(val, "|")
			result := University{university_Name: SeperatedVals[1], city: SeperatedVals[2], state: SeperatedVals[3]}
			Wikilinks = append(Wikilinks, "https://en.wikipedia.org/wiki/"+strings.ReplaceAll(result.university_Name, " ", "_"))
			UniversityList = append(UniversityList, result)
		}
	}
}
