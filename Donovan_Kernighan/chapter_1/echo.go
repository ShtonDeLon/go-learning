// Пакет с упражнениями из первого раздела книги
package chapter1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Упражнение 1.1
// Вывод аргументов командной строки
func Echo1() {
	var s, sep string
	sep = " "
	//_ (blank identifier) пустой идентификатор
	for _, a := range os.Args {
		s += sep + a + "\n"
	}
	fmt.Print(s)
}

// Упражнение 1.2
// Вывод аргументов командной строки с их индексами
func Echo2() {
	var s, sep string
	sep = " "
	for i, a := range os.Args {
		s += strconv.Itoa(i) + sep + a + "\n"
	}
	fmt.Print(s)
}

// Упражнение 1.3
// Вывод аргументов командной строки через string.Join
func Echo3() {
	fmt.Print(strings.Join(os.Args, "\n"))
}

// Упражнение 1.3_1
// Сравнение времении выполнения Echo1 и Echo3
func EchoDelta() {
	t := time.Now().Nanosecond()
	Echo3()
	s := fmt.Sprintf("Время выполнения Echo3 (strings.Join): %v ns\n", time.Now().Nanosecond()-t)
	t = time.Now().Nanosecond()
	Echo1()
	fmt.Printf("Время выполнения Echo1 (for): %v ns\n", time.Now().Nanosecond()-t)
	fmt.Print(s)
}
