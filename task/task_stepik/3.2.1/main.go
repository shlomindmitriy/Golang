package main
/*
Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.
Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.
Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы, вам ничего импортировать не нужно!
Считывать и выводить ничего не нужно!
Ваша функция должна называться adding() !
Считайте что функция и пакет main уже объявлены!

Sample Input:
	%^80 hhhhh20&&&&nd
Sample Output:
	100
*/
import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var strs [2]string
	for i, _ := range strs {
		fmt.Scan(&strs[i])

	}

	fmt.Println(adding(strs[0], strs[1]))
}

func adding(str1, str2 string) int64 {
	mas := []rune(str1)

	rez := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	res := make([]rune, 0)

	for i := 0; i < len(mas); i++ {
		for j := 0; j < len(rez); j++ {
			if mas[i] == rez[j] {
				res = append(res, mas[i])
				break

			}
		}

	}
	summa, err := strconv.Atoi(string(res))
	if err != nil {
		log.Println("ошибка парсинга Atoi(в первом fore)")
	}

	mas1 := []rune(str2)
	res1 := make([]rune, 0)
	for i := 0; i < len(mas1); i++ {
		for j := 0; j < len(rez); j++ {
			if mas1[i] == rez[j] {
				res1 = append(res1, mas1[i])
				break
			}
		}

	}
	summa1, err := strconv.Atoi(string(res1))
	if err != nil {
		log.Println("ошибка парсинга Atoi(во втором fore)")

	}

	return int64(summa + (summa1))

}
