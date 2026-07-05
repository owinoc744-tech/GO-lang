package main

import"fmt"

func main(){
fmt.Println("===Go calculator===")
var num ,num2 float64
operator := ("*,+,-,/")
fmt.scan(&num)
fmt.Scan(&operator)
fmt.Scan(&num2)

switch operator{

case "+"
result = num + num2
case "*"
result = num * num2

case "-"
result = num - num2

case "/"
result = num/num2
}
}
