package main

import "strings"

func main() {
	title := "\"Vol releas a new car for fdfdd\""
	shortTitle := generateShortTitle(title)
	print(shortTitle)
}

func generateShortTitle(source string) string {
	lenCut := 21
	quotationMark := "\""
	if len(source) <= lenCut {
		return source
	}
	sourceResult := source[0:lenCut]
	if source[lenCut:lenCut+1] != " " && source[lenCut:lenCut+1] != "!" && source[lenCut:lenCut+1] != "." {

		if strings.Contains(sourceResult, " ") {
			sourceResult = sourceResult[:strings.LastIndex(sourceResult, " ")] //todo? " "
		}
	}
	if sourceResult[0:1] != quotationMark {
		return quotationMark + sourceResult + "..." + quotationMark
	}
	return sourceResult + "..." + quotationMark
}
