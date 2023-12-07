package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() && (sc.Text() != `\quit`) {
		UnixShel(sc.Text())
	}
}

func UnixShel(s string) {
	switch strings.Split(s, " ")[0] {
	case "cd":
		cdCmd(s)
	case "pwd":
		pwdCmd()
	case "echo":
		echoCmd(s)
	case "kill":
		killCmd(s)
	case "ps":
		psCmd()
	default:
		fmt.Println("Error! Invalid command!")
	}
}

func cdCmd(cmd string) {
	cmd = string([]rune(cmd)[3:])

	err := os.Chdir(cmd)
	if err != nil {
		fmt.Println(err)
	}
}

func pwdCmd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path)
}

func echoCmd(cmd string) {
	cmd = string([]rune(cmd)[5:])
	fmt.Println(cmd)
}

func killCmd(cmd string) {
	cmd = string([]rune(cmd)[5:])
	pid, err := strconv.Atoi(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	proc.Kill()
}

func psCmd() {
	sliceProc, _ := ps.Processes()
	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
}
