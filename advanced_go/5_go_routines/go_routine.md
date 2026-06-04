# Go Goroutines & Concurrency: Learning Summary

This document compiles the core concepts for Go's concurrency model (Goroutines) and how to manage them, translated for a JavaScript/TypeScript perspective.

## 1. Concurrency (Goroutines)
* **What it is:** Concurrency allows your program to run multiple tasks simultaneously. Instead of one waiter serving tables one by one (synchronous), it's like having multiple waiters serving different tables at the exact same time.
* **The `go` Keyword:** Goroutines are Go's lightweight threads. You spin one off into the background by simply placing the word `go` in front of a function call (e.g., `go processTruck()`). It's conceptually similar to firing off an async function without awaiting it.

## 2. The Asynchronous Trap
* **The Problem:** When you fire off a Goroutine, it runs in the background. However, if your main `main()` function reaches the end of its code, the program shuts down instantly. 
* **The Result:** Any Goroutines still running in the background are violently canceled and "leaked" before they can finish their work.

## 3. WaitGroups (Go's `Promise.all`)
To prevent the program from shutting down early, you must synchronize your routines using `sync.WaitGroup`. It acts as a bouncer that waits for a collection of Goroutines to finish.
* **`wg.Add(1)`:** Tells the WaitGroup to add one task to its tally. (Place this right before launching the Goroutine).
* **`wg.Done()`:** Placed at the end of the Goroutine's function (usually an anonymous function). It tells the WaitGroup, "I finished my specific task," subtracting 1 from the tally.
* **`wg.Wait()`:** Placed at the end of your main program. It blocks the program from exiting until the WaitGroup tally hits zero.

## 4. Performance Gains
* Running 5 heavy tasks that take 1 second each **synchronously** will take **5 seconds total**.
* By using the `go` keyword and firing them off concurrently, all 5 tasks process at the exact same time, taking only **1 second total**.

## 5. Error Handling Limitations
* If an error occurs *inside* a Goroutine, the main program is blind to it. You cannot simply `return` an error like a normal function. 
* To pass errors or data back and forth between running Goroutines, you must use **Channels**.