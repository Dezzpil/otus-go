package lesson3

import (
	"errors"
	"strconv"
	"unicode"
)

/*
Распаковка строки
Цель: Создать Go функцию, осуществляющую примитивную распаковку строки,
содержащую повторяющиеся символы / руны, например:
* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)
* "" => ""

Дополнительное задание: поддержка escape - последовательностей
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)

В случае если была передана некорректная строка функция должна возвращать ошибку.
*/
func Unpack(value string) (string, error) {
	var (
		last rune
		result string
		escape bool
	)
	for _, symbol := range value {

		if string(symbol) == "\\" && escape != true {
			escape = true
			continue
		}

		if unicode.IsNumber(symbol) {
			if !escape {
				if last == 0 {
					return result, errors.New(`value mustn't starts from digit`)
				}
				if digit, err := strconv.Atoi(string(symbol)); err == nil {
					for i := 1; i < digit; i += 1 {
						result += string(last)
					}
				}
				continue
			}
		}

		last = symbol
		result += string(last)
		escape = false
	}

	return result, nil
}
