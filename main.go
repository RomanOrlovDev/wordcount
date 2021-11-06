package main

import (
	"bufio"
	"fmt"
	"github.com/RomanOrlovDev/wordcount/utility"
	"os"
)

func readScanner() (error, string){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Err(), scanner.Text()
}

func main(){

	err, ret := readScanner()
	if err != nil {
		panic("Bad")
	}

	fmt.Println(utility.Count(ret))
}
