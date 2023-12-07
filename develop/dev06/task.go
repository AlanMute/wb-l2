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
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fields := flag.Int("f", 1, "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	if *fields < 1 {
		*fields = 1
	}

	filename := flag.Arg(0)
	if filename == "" {
		log.Fatal("Введите название файла!")
	}

	data := ReadFile(filename)
	fmt.Printf("Результат:\n\n")

	res := cut(data, *fields, *delimiter, *separated)
	for _, s := range res {
		fmt.Println(s)
	}
}

func cut(data []string, fields int, delimiter string, separated bool) []string {
	var res []string

	for _, s := range data {
		t := strings.Split(s, delimiter)
		if separated {
			if t[0] != s {
				if fields < len(t) {
					res = append(res, t[fields-1])
				}
			}
		} else {
			if fields < len(t) {
				res = append(res, t[fields-1])
			}
		}
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
