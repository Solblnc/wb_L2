package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
//
//
//- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
//- pwd - показать путь до текущего каталога
//- echo <args> - вывод аргумента в STDOUT
//- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
//- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
//
//
//
//
//Так же требуется поддерживать функционал fork/exec-команд
//
//
//Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
//
//
//*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
//в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
//и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

func cd(directory string) {
	if err := os.Chdir(directory); err != nil {
		fmt.Println(err.Error())
	}
}

func pwd() string {
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return d
}

func echo(args []string) {
	for _, arg := range args {
		fmt.Println(arg + " ")
	}
	fmt.Println()
}

func kill(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}

	if err = process.Kill(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

func ps() {
	cmd := exec.Command("ps", "-A")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("ps: ", err)
	}
	fmt.Println(string(output))
}

func execute(query string) {
	commands := strings.Split(query, " | ")
	for _, command := range commands {
		commandSlice := strings.Split(command, " ")
		switch commandSlice[0] {
		case "cd":
			cd(commandSlice[1])
		case "pwd":
			fmt.Println(pwd())
		case "echo":
		case "kill":
			pid, err := strconv.Atoi(commandSlice[1])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			kill(pid)
		case "ps":
			ps()

		}
	}
}

func main() {
	wd, _ := os.Getwd()
	fmt.Print(wd + "> ")

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); fmt.Print("$ ") {
		if query := scanner.Text(); query != "\\quit" {
			execute(query)
		} else {
			break
		}
	}
}
