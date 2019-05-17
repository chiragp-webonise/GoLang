package main 


func main() {
	messages:=make(chan string)
	wg:= *sync.WaitGroup
	<-messages
	go func(){ 
		messages<-"ping"
	}()

	
	//fmt.Println(msg)
}