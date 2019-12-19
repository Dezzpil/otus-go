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

// 0000-007f // latin range
// 0410-044f // subrange from cyrillic range
// 0401 // ё
// 0451 // Ё
var CustomRange = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0000, 0x007f, 1},
		{0x0401, 0x0401, 1},
		{0x0410, 0x044f, 1},
		{0x0451, 0x0451, 1},
	},
}

func detectSymbol(char rune) bool {
	if !unicode.Is(CustomRange, char) {
		return true
	}
	return false
}

func main() {
	file := os.Stdin
	// defer file.Close() // Do i need to close stdin?

	sc := bufio.NewScanner(file)
	m := make(map[rune]uint32)
	for sc.Scan() {
		line := sc.Text()
		for _, char := range line {
			if detectSymbol(char) {
				if _, ok := m[char]; !ok {
					m[char] = 0
				}
				m[char] += 1
			}
		}
	}

	err := sc.Err()
	survive(err)

	// --as-csv
	fmt.Printf("unicode;character;count\n")
	for key, val := range m {
		fmt.Printf("\"%U\";%q;%d\n", key, string(key), val)
	}
}

