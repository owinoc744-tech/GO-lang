package main

import "fmt"

func main(){

var name string
var age int
var weight float64

fmt.Println("Enter your name:")
fmt.Scan(&name)

fmt.Println("Enter your age:")
fmt.Scan(&age)

fmt.Println("Enter your weight:")
fmt.Scan(&weight)

fmt.Printf("Hi %v you are %v years oldand your weight is %v.\n",name,age ,weight)

}
