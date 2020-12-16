package main

import(
  "gosample/mypkg"
	"fmt"
)


func main(){
  p1 := mypkg.Person{Name: "Go太郎", Age: 10}
	fmt.Println(p1.Name, p1.Age)
	fmt.Println(mypkg.Add(3,5))
  fmt.Println(mypkg.TAX)
}