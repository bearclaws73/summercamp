package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func LongestWord(rdr io.Reader) (string, int) {
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	var longestWord string
	longWordLength := 0
	for scanner.Scan() {
		wordOrManyWords := scanner.Text()
		wordOrManyWords = strings.Replace(wordOrManyWords, "-", " ", -1)
		myword := scanner.Text()
		if len(myword) > longWordLength {
			longWordLength = len(myword)
			longestWord = myword
		}
	}
	return longestWord, longWordLength
}

func main() {
	srcFile, err := os.Open("../mobyDick/moby.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()
	longestWord, longWordLength := LongestWord(srcFile)

	fmt.Println(longestWord, longWordLength)

}
