package main

import  (
	"fmt"
	"github.com/RomanOrlovDev/wordcount/utility"
	"os"
)

func readScanner() string{
	return os.Args[1]
}

func main(){
	ret := readScanner()
	fmt.Println(utility.Count(ret))
}
