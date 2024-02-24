package helpers

import "strings"

func GetExtention(fileName string) (string){
	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}