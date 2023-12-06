package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	k := flag.Int("k", 1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	filename := flag.Arg(0)
	if len(filename) == 0 {
		log.Fatalf("Введите название файла!")
	}
	data := ReadFile(filename)

	if *k < 1 {
		log.Fatal("Порядковый номер колонки не может быть меньше единицы")
	}
	data = SortByColumn(data, *k, *n)

	if *u {
		data = UniqueStrings(data)
	}
	if *r {
		data = ReverseSlice(data)
	}
	SaveFile(data)

}

func SaveFile(data [][]string) {
	filename := "output.txt"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i, s := range data {
		_, err = f.WriteString(strings.Join(s, " "))
		if err != nil {
			log.Fatal(err.Error())
		}
		if i < len(data)-1 {
			_, err = f.WriteString("\n")
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

func ReverseSlice(data [][]string) [][]string {
	length := len(data)
	for i := 0; i < length/2; i++ {
		data[i], data[length-1-i] = data[length-1-i], data[i]
	}
	return data
}

func UniqueStrings(data [][]string) [][]string {
	var m = make(map[string]int)

	for i := 0; i < len(data); i++ {
		m[strings.Join(data[i], " ")]++
		if m[strings.Join(data[i], " ")] > 1 {
			data = append(data[:i], data[i+1:]...)
			i--
		}
	}

	return data
}

func SortByColumn(data [][]string, k int, n bool) [][]string {
	sort.Slice(data, func(a, b int) bool {
		if n {
			var na, nb int
			var erra, errb error

			if len(data[a]) >= k {
				na, erra = strconv.Atoi(string(data[a][k-1]))
			}
			if len(data[b]) >= k {
				nb, errb = strconv.Atoi(string(data[b][k-1]))
			}
			switch {
			case ((len(data[a]) < k) || erra != nil) && ((len(data[b]) >= k) && errb == nil):
				return false
			case ((len(data[a]) >= k) && erra == nil) && ((len(data[b]) < k) || errb != nil):
				return true
			case ((len(data[a]) < k) || erra != nil) && ((len(data[b]) < k) || errb != nil):
				return data[a][0] < data[b][0]
			default:
				return na < nb
			}
		} else {
			switch {
			case (len(data[a]) < k) && (len(data[b]) >= k):
				return true
			case (len(data[a]) >= k) && (len(data[b]) < k):
				return false
			case (len(data[a]) < k && len(data[b]) < k):
				return data[a][0] < data[b][0]
			default:
				return data[a][k-1] < data[b][k-1]
			}
		}
	})

	return data
}

func ReadFile(filename string) [][]string {

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("Ошибка открытия файла", filename)
	}

	sc := bufio.NewReader(file)
	var rows [][]string

	for {
		l, err := sc.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err.Error())
		}
		l = strings.TrimRight(l, "\r\n")
		rows = append(rows, strings.Split(l, " "))
		if err == io.EOF {
			break
		}
	}

	return rows
}
