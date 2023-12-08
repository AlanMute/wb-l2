package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Flags struct {
	host    string
	port    string
	timeout time.Duration
}

func flagInit() *Flags {
	var flags *Flags = new(Flags)
	t := *flag.String("timeout", "10s", "connection timeout")
	flag.Parse()

	value, err := strconv.Atoi(t[:len(t)-1])
	if err != nil {
		log.Fatal(err)
	}

	switch t[len(t)-1] {
	case 's':
		flags.timeout = time.Duration(value) * time.Second
	case 'm':
		flags.timeout = time.Duration(value) * time.Minute
	case 'h':
		flags.timeout = time.Duration(value) * time.Hour
	default:
		flags.timeout = 10 * time.Second
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	if len(host) == 0 || len(port) == 0 {
		log.Fatal("Введите хост и порт!")
	}

	_, err = strconv.Atoi(port)
	if err != nil {
		log.Fatal("Неккоректно введен порт!")
	}

	flags.host = host
	flags.port = port

	return flags
}

func telnetClient(flags *Flags) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	host := flags.host + ":" + flags.port

	con, err := net.DialTimeout("tcp", host, flags.timeout)

	if err != nil {
		log.Fatal("Connect error: ", err)
	}

	go func(con net.Conn) {
		<-sigs
		err := con.Close()
		if err != nil {
			log.Println("Невозможно закрыть сокет")
			os.Exit(1)
		}
		os.Exit(0)
	}(con)

	go func() {
		_, err := io.Copy(os.Stdout, con)
		if err != nil {
			fmt.Println("Сокет закрыт")
			os.Exit(0)
		}
	}()

	for {
		_, err = io.Copy(con, os.Stdin)
		if err != nil {
			fmt.Println("Сокет закрыт")
			os.Exit(0)
		}
	}
}

func main() {
	flags := flagInit()
	telnetClient(flags)
}
