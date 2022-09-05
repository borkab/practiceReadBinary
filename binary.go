package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}
func main() {
	//we read an image and print it in hexadecimal format

	file, err := os.Open("20220831_231054.jpg")
	check(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		//the Dump returns a string that contains a hex dump of the given data
		fmt.Printf("%s", hex.Dump(buf))
	}

}
