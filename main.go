package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var line []byte
	var lines [][]byte
	for i := 1; i <= 855; i++ {
		line, err = r.ReadBytes('\n')
		lines = append(lines, line)
	}
	var dictionary [][][]byte
	var char [][]byte
	count := 0
	//i <= len(lines)-1
	// fmt.Println(lines[0:9])
	// fmt.Println(lines[847:855])

	for i := 0; i <= len(lines)-1; i++ {
		if len(lines[i]) == 1 {
			continue
		} else {
			lines[i] = lines[i][:len(lines[i])-1]
			//fmt.Println("this is lines after trim: \n", lines[i])
			char = append(char, lines[i])
			count++
			if count == 8 {
				// char = append(char, lines[i])
				dictionary = append(dictionary, char)
				char = [][]byte{}
				count = 0
			}
		}

	}
	//fmt.Println(dictionary)
	//fmt.Printf("%q\n", dictionary)
	//fmt.Println(char)
	fmt.Println(dictionary[72])
	test := "he"
	fmt.Println(len(dictionary))
	//var display [][][]byte
	for _, e := range test {
		char = dictionary[e-32]
		fmt.Println(char)
		// for i := 0; i < 8; i++ {
		// 	display = append(display, dictionary[e-32])
		// }
	}
	// for _, e := range display {
	// 	fmt.Printf("%s\n", e)
	// }

	if err != io.EOF {
		//log.Fatal(err)
	}
}
