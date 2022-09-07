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
	var lineT [][]byte
	for i := 1; ; i++ {
		line, err = r.ReadBytes('\n')
		//fmt.Printf(" type: %T", line)
		//fmt.Printf("[line:%d] %q\n", i, line)
		//fmt.Println(line)
		if i == 12 {
			lineT = append(lineT, line)
		}

		if i == 21 {
			lineT = append(lineT, line)
		}

		if i == 30 {
			lineT = append(lineT, line)
		}

		if i == 39 {
			lineT = append(lineT, line)
			break
		}
		if err != nil {
			break
		}
	}

	fmt.Printf("%q", lineT)
	if err != io.EOF {
		log.Fatal(err)
	}
}
