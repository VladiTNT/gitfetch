package files

import "strings"

func GetFileExtension(name string) string {
	subStrings := strings.SplitAfter(name, ".")
	return subStrings[len(subStrings)-1]
}
