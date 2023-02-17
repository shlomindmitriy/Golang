package main

import "fmt"
/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
 Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита.
  На вход подается строка-пароль. 
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
	fdsghdfgjsdDD1

Sample Output:
	Ok

*/

func main() {
	var a string
	fmt.Scan(&a)

	res := []rune(a)

	letter := []rune("Q,q,W,w,E,e,R,r,T,t,Y,y,U,u,I,i,O,o,P,p,A,a,S,s,D,d,F,f,G,g,H,h,J,j,K,k,L,l,Z,z,X,x,C,c,V,v,B,b,N,n,M,m,1,2,3,4,5,6,7,8,9,0")

	if len(res) < 5 {
		fmt.Println("длинна не соответствует")
		return
	}

	for i := 0; i < len(res); i++ {
		var flag bool = false
		
		for j := 0; j < len(letter); j++ {

			if res[i] == letter[j] {
				flag = true
				break

			}

		}
		if flag == false {
			fmt.Println("er")
			return
		}
	}
	fmt.Println("ok")
}
