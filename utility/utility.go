package utility

import "strings"

func Count(str string) (numberOfWords int){
	if str == ""{
		return 0
	}else{
		return len(strings.Split(str, " "))
	}
}