package main

import "fmt"

func main(){
// arrays with fixed nember of elements
var name [6] string = [6] string {"john" ,"nigel","ann","peter","lego","ted"}

fmt.Println(name,len(name))
//replacing elements in an array
name[2]="mosby"
fmt.Println(name,len(name))
// we cannot add elements in an array because it has a fixed number of elements
// slices
num := []int{1,1,2,3,4,5, 6}
fmt.Println(num,len(num))
// to append is to add more elements in a slice because slices have no fixed number of elements
num = append(num,7,8,9)
fmt.Println(num,len(num))
// ranges gives you your specified selection in both arrays and slices via  indexing the elements.
Range1 := name[3:]
Range2 := name[:5]
Range3 := num [:8]
fmt.Println(Range1)
fmt.Println(Range2)
fmt.Println(Range3)
}
