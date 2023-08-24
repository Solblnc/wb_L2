package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

//Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
//
//Требования:
//Программа должна быть оформлена как go module
//Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

func main() {
	getTime()
}

func getTime() string {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Errorf("cannot get time: %w", err)
	}
	fmt.Println(time)

	return time.String()
}
