package main

import (
	"fmt"
	// "regexp"
)

// отделение чисел от букв  и  их вывод

func main() {

	// a := "mfgah134517095aldrfgvh8h"
	// b := regexp.MustCompile("[0-9]+")

	// fmt.Println(b.FindAllString(a,-1))

	a := "45fgah134517095aldrfgvh8h"
	v := TrimFunc(a)

	fmt.Println(v)
}

func TrimFunc(s string) string {
	resp := ""

	for _, v := range s {
		if ok := checkNumber(v); ok {
			resp += string(v)

		}

	}

	return resp
}

func checkNumber(s rune) bool {

	switch s {
	case '0':
		return true
	case '1':
		return true

	case '2':
		return true
	case '3':
		return true
	case '4':
		return true
	case '5':
		return true
	case '6':
		return true
	case '7':
		return true
	case '8':
		return true
	case '9':
		return true

	}

	return false
}
