package Stack

import (
	"fmt"
)

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack returns a new stack.
func Newstack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Values() []*Node {
		return s.nodes[:s.count]
}
func (s *Stack) Empty() bool {
		if s.count == 0{
			return true
		}
		return false

}
func (s *Stack) Size() int {
		return s.count
	}
// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
// func main() {
// 	s := Newstack()
	
// 	s.Push(&Node{1})
// 	s.Push(&Node{2})
// 	s.Push(&Node{3})
// 	vals:=s.Values()
// 	fmt.Println(s.Empty())
// 	fmt.Println(vals[0])
// 	fmt.Println(s.Pop(), s.Pop(), s.Pop())
// }
 