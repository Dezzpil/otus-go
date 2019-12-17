package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func survive(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./query-result.tsv")
	survive(err)
	defer file.Close()

	m := make(map[rune]uint32)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if detectSymbol(char) {
				//fmt.Printf("%U %q\n", char, char)
				if _, ok := m[char]; !ok {
					m[char] = 0
				}
				m[char] += 1
			}
		}
	}

	err = scanner.Err();
	survive(err)

	fmt.Printf("unicode;character;count\n")
	for key, val := range m {
		fmt.Printf("\"%U\";%q;%d\n", key, string(key), val)
	}
}

func detectSymbol(char rune) bool {
	if !unicode.IsNumber(char) {
		if !unicode.IsControl(char) {
			if !unicode.Is(unicode.Common, char) {
				if !unicode.Is(unicode.Cyrillic, char) {
					// if unicode.Is(unicode.Latin, char) && unicode.Is{
						return true
					//}
				}
			}
		}
	}
	return false
}