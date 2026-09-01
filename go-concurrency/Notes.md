# Go-routine 
- It is a lightweight function that runs concurrently with the other functions in a go programm .
- You start one by placing go before a function call . 
```go 

package main 
import "fmt"

func sayHello() {
    fmt.Println("Hello from go routine")
}
func main () {
    go sayHello() {
        fmt.Println("Hello from main")
    }
}

// the output mau onlu show Hello from main 
// because the main function can finish before the go routine runs .
```
# Important concepts 
- Concurrency means managing multiple tasks during the same period .
- Parallelism means tasks actually run at the same time on different CPU cores. Goroutines can
  provide concurrency and, when resources allow, parallelism. 
- Go routine are not exactly os level threads . Goroutines are managed by the Go runtime. Go schedules many goroutines over a smaller number of
operating-system threads.

- This makes goroutines:

  - Lightweight
  - Fast to create
  - Suitable for thousands of concurrent tasks
  - Less expensive than manually creating threads

- In go also main does not wait automatically .The programmg