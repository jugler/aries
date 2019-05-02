package listutils

import (
	"math/rand"
	"strings"
)

//RandomList returns the parameter list in a random sort
func RandomList(list []string) []string {
	for i := range list {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}

//FilterByTag returns an array filtered by the parameter filterTag
func FilterByTag(fileNames []string, filterTag string) (filteredFileNames []string) {
	for _, fileName := range fileNames {
		if filterTag == "all" || (filterTag == "sfw" && !strings.Contains(fileName, "nsfw")) {
			filteredFileNames = append(filteredFileNames, fileName)
		} else if filterTag != "sfw" {
			//split image by -
			//filter by tag
			tags := strings.Split(fileName, "-")
			for _, tag := range tags {
				if strings.Contains(tag, filterTag) {
					filteredFileNames = append(filteredFileNames, fileName)
				}
			}
		}
	}
	return filteredFileNames
}
