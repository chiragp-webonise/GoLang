package main
import (
        "fmt"
        "sync"
        "github.com/Scheduling_assignment/interval"
        )
var wg sync.WaitGroup

    func m1(){
        fmt.Println("From method m1")
        
    }
    func m2(){
        fmt.Println("From method m2")
        
    }
    func m3(){
        fmt.Println("From method m3")
        
    }
    func main() {
    	
    	interval.Exe(m1,&wg,500)
    	interval.Exe(m2,&wg,1500)
    	interval.Exe(m3,&wg,3000)
    	wg.Wait()
    }