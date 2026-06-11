# Go Context & Timeouts: Short & Sweet Summary

Bro, here are the core fundas for `context` in Go. Keep it simple!

## 1. What is Context?
It basically does two main jobs:
* **Passing Data:** Sending global data (like User ID) deep into your functions (just like React Context).
* **Timeouts:** Canceling slow tasks before users get frustrated waiting (like `AbortController` in JS).

## 2. The Golden Rules
* **First Argument Only:** Always pass `ctx context.Context` as the very first parameter in your functions. It is a strict Go community rule.
* **It is Immutable:** You cannot change an existing context. You must take a parent context (like `context.Background()`) and wrap it to create a new one.

## 3. Passing Data (Avoid Collisions!)
* **Never use plain strings as keys.** If you use `"id"`, some other 3rd party package might also use `"id"` and permanently overwrite your user data!
* **The Fix:** Create a custom type for your keys.
```go
//go:build ignore
type mySafeKey string
const userIDKey mySafeKey = "userID"
// Now it is 100% safe from overwriting!
```

## 4. Timeouts & The "Diwali Firecracker" Rule 🧨
* The timer starts the **exact moment** you write `context.WithTimeout`. 
* It does not pause. It does not wait for your slow function to start. The fuse is lit immediately!
* Always create the timeout right before you are ready to pass it to the function.
* **Memory Leak Warning:** Always write `defer cancel()` right after creating a timeout to clean up the background timers.

```go
//go:build ignore
// 1. Timer starts RIGHT HERE! Tick... tick...
ctx, cancel := context.WithTimeout(bgCtx, 2*time.Second)
defer cancel() // Always clean up!

// 2. If you waste time here, the context burns out in the background!
time.Sleep(3 * time.Second)

// 3. By the time it reaches the function, it's already dead!
processSlowTask(ctx) 
```

## 5. The `select` Statement (The Race)
* `select {}` pauses your code and waits for multiple channels.
* It sets up a race between your actual task (e.g., 5 seconds) and your context timeout (e.g., 2 seconds).
* Whichever finishes first, wins! If the context timer wins, it aborts the slow task instantly and triggers `ctx.Done()`.

```go
//go:build ignore
func processSlowTask(ctx context.Context) {
	// The Race begins!
	select {
	case <-time.After(5 * time.Second):
		// Case A: The task finishes first (Wins if task is < 2s)
		fmt.Println("Task finished successfully!")
		
	case <-ctx.Done():
		// Case B: The 2-second timeout hits first and aborts the task!
		fmt.Println("Error:", ctx.Err()) // Prints: "context deadline exceeded"
	}
}
```