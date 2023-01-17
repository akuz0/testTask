package main

import "strings"

func main() {
	title := "Volvo released"
	shortTitle := generateShortTitle(title)
	print(shortTitle)
}

func generateShortTitle(source string) string {
	lenCut := 25
	
	if len(source) < lenCut {
		return source
	}
	sourceResult := source[0:lenCut]
	if source[lenCut:lenCut+1] != " " && source[lenCut:lenCut+1] != "!" && source[lenCut:lenCut+1] != "." {

		if strings.Contains(sourceResult, " ") {
			sourceResult = sourceResult[:strings.LastIndex(sourceResult, " ")] //todo? " "
		}
	}

	return sourceResult + "..."
}
