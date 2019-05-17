package main 

import "fmt"

type A struct {
	age int
}

func (a *A) display() int{
	return a.age
}
type B struct{
	A
	name string
}
func (b *B) display() string{
	a.age=12
	re
	turn b.name
}
func main(){

  m:=A{age:23}
  x:=B{name:"chirag"}
  fmt.Println(m)
  fmt.Println(x)
}
