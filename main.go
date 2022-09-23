package main

import (
	ascii "ascii/drawing"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Wrong number of arguments\nCorrect usage: go run . [STRING]")
	}
	sourceText := os.Args[1]
	bannerName := os.Args[2]
	//fmt.Println(sourceText)
	// fmt.Println(bannerName)
	//ascii.Display(sourceText, bannerName)
	//handle \n from string before calling drawing
	var finalText string
	//test := "hello\n\nthere"
	sourceTextArr := strings.Split(sourceText, "\\n")
	//fmt.Println(sourceTextArr)
	for _, v := range sourceTextArr {
		if len(v) != 0 {
			res := ascii.Display(v, bannerName)
			finalText += res
		} else {
			finalText += "\n"
		}

	}
	fmt.Print(finalText)
	// regQuote1 := `/\\(\s)n`
	// regQuote2 := `/\\n`

}
