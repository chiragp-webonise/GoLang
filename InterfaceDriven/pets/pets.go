package pets

import "fmt"

type Dog struct {
    Name  string
    Breed string
}

func (d Dog) Walk() {
    fmt.Println(d.Name, "walks across the room")
}
func (d Dog) Sit() {
    fmt.Println(d.Name, "sits down")
}
func (d Dog) Fetch() {
    fmt.Println(d.Name, "fetches a toy")
}

type Cat struct {
    Name  string
    Breed string
}

func (c Cat) Walk() {
    fmt.Println(c.Name, "walks across the room")
}
func (c Cat) Sit() {
    fmt.Println(c.Name, "sits down")
}
func (c Cat) Purr() {
    fmt.Println(c.Name, "purrs")
}