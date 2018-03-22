package main 
import "fmt"

type person struct{

	name string
	age int
}
func main() {
	fmt.Println(person{"chirag",22})
	fmt.Println(person{name:"niketa",age:70})
	fmt.Println(person{name:"nikhil"})
	fmt.Println(&person{name:"sagar",age:22})

	s:=person{name:"smit",age:22}
	fmt.Println(s.name)
	s1:=&s
	s1.name="hardik"
	fmt.Println(s1.name)
}