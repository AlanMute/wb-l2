package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
*/

func main() {
	flag.Parse()
	site := flag.Arg(0)
	if len(site) == 0 {
		log.Fatal("Введите url сайта!")
	}
	Wget(site)
	fmt.Println("Скачивание прошло успешно!")
}

func GetPathFile(s string) string {
	sr := strings.Split(s, "/")[2] + ".html"
	return sr
}

func Wget(s string) {
	resp, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	path := GetPathFile(s)

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}
