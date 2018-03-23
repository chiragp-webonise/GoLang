package interval
import (
        "sync"
        "time"
        )
   func Exe(function func(),wg *sync.WaitGroup,t time.Duration) {
        //for f := range funcChan {
        //i++
        ticker := time.NewTicker( t * time.Millisecond)
        wg.Add(1)

        go func() {
            wg.Add(1)
            for _ = range ticker.C {
                go function()
            }
            wg.Done()
        }()
    }
    
    