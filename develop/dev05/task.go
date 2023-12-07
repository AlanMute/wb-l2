package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type strSlice struct {
	after   int
	before  int
	context int
	count   int
}

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Int("c", 0, "количество строк")
	ignorecase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	linenum := flag.Bool("n", false, "печатать номер строки")

	flag.Parse()

	filename := flag.Arg(0)
	if len(filename) == 0 {
		log.Fatalf("Введите название файла!")
	}
	var words string
	br := bufio.NewReader(os.Stdin)
	fmt.Println("Что вы хотите найти в файле:")
	fmt.Fscan(br, &words)
	fmt.Printf("\nРезультаты поиска:\n\n")

	data := ReadFile(filename)

	strslice := strSlice{}

	if *count != 0 {
		strslice.count = *count
	} else {
		strslice.count = len(data)
	}

	switch {
	case *after > 0:
		strslice.after = *after
	case *before > 0:
		strslice.before = *before
	case *context > 0:
		strslice.after, strslice.before = *context, *context

	default:
		strslice.after = len(data)
	}

	for i, s := range data {
		if manGrep(s, words, *ignorecase, *invert, *fixed) {
			a := max(0, i-strslice.before)
			b := min(len(data)-1, i+strslice.after)

			for ; a <= b; a++ {
				if strslice.count > 0 {
					if *linenum {
						fmt.Printf("%d: %s\n", a, data[a])
					} else {
						fmt.Println(data[a])
					}
					strslice.count--
				} else {
					os.Exit(1)
				}
			}
			fmt.Println()
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func manGrep(s string, words string, i, v, F bool) bool {
	var res bool
	if i {
		s = strings.ToLower(s)
		words = strings.ToLower(words)
	}
	if F {
		res = words == s
	} else {
		res = strings.Contains(s, words)
	}
	if v {
		res = !res
	}
	return res
}

func ReadFile(filename string) []string {

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("Ошибка открытия файла", filename)
	}

	sc := bufio.NewReader(file)
	var rows []string

	for {
		l, err := sc.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err.Error())
		}
		l = strings.TrimRight(l, "\r\n")
		rows = append(rows, l)
		if err == io.EOF {
			break
		}
	}

	return rows
}
