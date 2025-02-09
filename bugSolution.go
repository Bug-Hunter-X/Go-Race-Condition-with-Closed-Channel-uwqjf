```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 10) // Buffered channel

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        val := <-ch
                        fmt.Printf("Goroutine %d received %d from channel\n", i, val)
                }(i)
        }

        for i := 0; i < 10; i++ {
                ch <- i
        }
        close(ch)
        wg.Wait()
}
```