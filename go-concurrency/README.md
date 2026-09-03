# Go Concurrency Learning Path

1. Goroutines
   - Creating Goroutines
   - Goroutine Lifecycle
   - Concurrency and Parallelism
   - Goroutine Scheduling

2. Synchronization
   - `sync.WaitGroup`
   - `sync.Mutex`
   - `sync.RWMutex`
   - `sync.Once`

3. Channels
   - Unbuffered Channels
   - Buffered Channels
   - Send and Receive Operations
   - Directional Channels
   - Closing Channels
   - Ranging Over Channels

4. Select
   - Multiple Channel Operations
   - Timeouts
   - Default Case
   - Non-Blocking Operations

5. Context
   - Cancellation
   - Deadlines
   - Timeouts
   - Context Propagation

6. Concurrency Patterns
   - Worker Pool
   - Producer-Consumer
   - Fan-Out
   - Fan-In
   - Pipeline
   - Task Queue
   - Rate Limiter
   - Pub/Sub

7. Error Handling
   - Error Channels
   - Concurrent Error Propagation
   - `errgroup`
   - Partial Failure Handling

8. Shared State
   - Data Races
   - Atomic Operations
   - `sync.Map`
   - Channels Versus Mutexes

9. Goroutine Lifecycle Management
   - Graceful Shutdown
   - Goroutine Leak Prevention
   - Backpressure
   - Bounded Concurrency

10. Testing and Debugging
    - Race Detector
    - Concurrent Unit Tests
    - Deadlock Detection
    - Goroutine Leak Detection
    - Benchmarks

11. Code-Level System Design
    - Concurrent Cache
    - Connection Pool
    - Job Scheduler
    - Rate Limiter Service
    - In-Memory Message Broker
    - Concurrent Web Crawler
