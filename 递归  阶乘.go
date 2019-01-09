package main

import "fmt"
var Sum int=1
func Factorial(n int){

	if n==1{
		return
	}
	Factorial(n-1)
	Sum*=n

}
func  main(){
	n:=10
	Factorial(n)
	fmt.Println(Sum)
}
