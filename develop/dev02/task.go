package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func strUnpack(str string) (string, error) {
	builder := strings.Builder{}
	n := len([]rune(str))
	if str == "" {
		return "", nil
	}
	if str[0] >= '0' && str[0] <= '9' {
		return "", fmt.Errorf("Некорректная строка!")
	}

	for i := 0; i < n; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			ind := i - 1
			count := 0
			for ; i < n && str[i] >= '0' && str[i] <= '9'; i++ {
				t, _ := strconv.Atoi(string(str[i]))
				count = 10*count + t
			}

			for j := 1; j < count; j++ {
				builder.WriteByte(str[ind])
			}
			i--
			continue
		}
		if str[i] == '\\' {
			builder.WriteByte(str[i+1])
			i++
			continue
		}
		builder.WriteByte(str[i])
	}

	return builder.String(), nil
}

func main() {
	fmt.Println(strUnpack("qwe\\45"))
}
