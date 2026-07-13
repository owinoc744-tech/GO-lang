package main

import"fmt"

func main(){
fmt.Println("===Go calculator===")
var num  float64
var num2 float64
operator := "*,+,-,/"
fmt.Println("Enter your calculation:")
fmt.Scanln(&num)
fmt.Scanln(&operator)
fmt.Scanln(&num2)

switch operator{

case "+":
fmt.Println( num + num2)
case "*":
fmt.Println( num * num2)

case "-":
fmt.Println(num - num2)

case "/":
fmt.Println( num/num2)
}
}
