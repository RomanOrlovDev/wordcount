package utility

import "strings"

func Count(str string) (numberOfWords int){
	return len(strings.Split(str, " "))
}