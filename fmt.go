package main

import "fmt"

func main(){

var name string
fmt.Println("Enter your name:")
fmt.Scan(&name)
var age int
fmt.Println("Enter your age:")
fmt.Scan(&age)
var nation string
fmt.Println("Enter your nationality:")
fmt.Scanln(&nation)

fmt.Printf("My name is %s i am %d years old and i am a %s citizen \n",name ,age ,nation)


}
