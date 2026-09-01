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

- In go also main does not wait automatically .The program exits when main exit . You need synchronization to wait . 

```go 

 package main

  import (
        "fmt"
        "sync"
  )

  func sayHello() {
        fmt.Println("Hello from goroutine")
  }

  func main() {
        var wg sync.WaitGroup

        wg.Add(1)

        go func() {
                defer wg.Done()
                sayHello()
        }()

        wg.Wait()
  }
```
- Each goroutine has its own execution flow .
- A normal function run directly.
- A go routine runs independently . 
- Goroutine need communication . GoRoutine often need to share results or coordinate work . In go this is commonly done with 
 - Channels
 - sync.WaitGroups 
 - sync.Mutex 
 - context.Context 

