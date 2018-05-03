package Stack

import (
	//"github.com/Stack"
	"testing"
	
)

func TestStackPush(t *testing.T) {
	stack := Newstack()
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	stack.Push(&Node{1})
	stack.Push(&Node{2})
	stack.Push(&Node{3})
	actualValue := stack.Values();
	a1:=actualValue[0].String();
	a2:=actualValue[1].String();
	a3:=actualValue[2].String();
	if   a1 != "1" || a2 != "2" || a3 != "3" {
		t.Errorf("Got %v expected %v", actualValue, "[3,2,1]")
	}
	if actualValue := stack.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := stack.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}
func TestStackPop(t *testing.T) {
	stack := Newstack()
	if actualValue:= stack.Pop(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	stack.Push(&Node{1})
	stack.Push(&Node{2})
	stack.Push(&Node{3})
	stack.Pop()
	actualValue:= stack.Pop(); 
	b1:=actualValue.String();
	if b1 != "2"  {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	actualValue= stack.Pop(); 
	b1=actualValue.String();
	if  b1 != "1" {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue := stack.Empty(); actualValue != true {
	 	t.Errorf("Got %v expected %v", actualValue, true)
	 }
	 if actualValue := stack.Values(); len(actualValue) != 0 {
	 	t.Errorf("Got %v expected %v", actualValue, "[]")
	 }
}
func TestStackPeek(t *testing.T) {
	stack := Newstack()
	if actualValue:= stack.Peek(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	stack.Push(&Node{1})
	stack.Push(&Node{2})
	stack.Push(&Node{3})
	if actualValue:= stack.Peek().String(); actualValue != "3"{
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}
func benchmarkPush(b *testing.B, stack *Stack, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Push(&Node{n})
		}
	}
}
func benchmarkPop(b *testing.B, stack *Stack, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Pop()
		}
	}
}
func BenchmarkArrayStackPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}
func BenchmarkArrayStackPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := Newstack()
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkArrayStackPush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkArrayStackPush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkArrayStackPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := Newstack()
	for n := 0; n < size; n++ {
		stack.Push(&Node{n})
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}
