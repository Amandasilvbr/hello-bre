// 1- Escreva um programa em Go que imprima um triângulo de asteriscos de acordo com o número de linhas que o usuário desejar;

package main

import "fmt"

func main(){
	var num int
	fmt.Scan(&num)
	simb := "*"

	for i:= 1; i <= num; i++ {
		fmt.Println(simb)
		simb = simb + "*" 
	}
}