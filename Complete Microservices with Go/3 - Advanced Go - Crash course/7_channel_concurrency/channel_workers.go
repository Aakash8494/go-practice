//go:build ignore

package main

import (
	"fmt"
	"time"
)

// 1. THE WORKER FUNCTION (The Waiter)
// Notice the arrows:
// <-chan means this pipe is READ-ONLY for the waiter
// chan<- means this pipe is WRITE-ONLY for the waiter
func restaurantWaiter(waiterID int, incomingOrdersPipe <-chan int, finishedFoodPipe chan<- string) {

	// The waiter constantly watches the pipe.
	// As soon as an order arrives, he grabs it!
	for currentOrder := range incomingOrdersPipe {
		fmt.Printf("-> Waiter %d grabbed Order %d\n", waiterID, currentOrder)

		// Simulating 1 second of cooking time...
		time.Sleep(1 * time.Second)

		// 2. PUSH RESULT: Food is ready, push it to the finished pipe!
		finishedMessage := fmt.Sprintf(
			"Order %d successfully cooked by Waiter %d!",
			currentOrder,
			waiterID,
		)
		finishedFoodPipe <- finishedMessage
	}
}

func main() {
	totalOrders := 5
	totalWaiters := 3

	// 3. CREATE THE PIPES (Buffered Channels)
	incomingOrdersPipe := make(chan int, totalOrders)
	finishedFoodPipe := make(chan string, totalOrders)

	fmt.Println("=== RESTAURANT OPENED ===")

	// 4. HIRE THE WAITERS (Spin up the Goroutines)
	// We spin up 3 workers. They will just stand there waiting for orders to arrive.
	for w := 1; w <= totalWaiters; w++ {
		go restaurantWaiter(w, incomingOrdersPipe, finishedFoodPipe)
	}

	// 5. RECEIVE ORDERS (Push jobs into the pipe)
	for jobID := 1; jobID <= totalOrders; jobID++ {
		incomingOrdersPipe <- jobID
	}

	// 6. LOCK THE JOBS PIPE
	// We tell the waiters: "No more new customers are coming!"
	// This allows their 'for range' loop to safely stop once the queue is empty.
	// close(incomingOrdersPipe)

	// 7. DELIVER THE FOOD (Read from the results pipe)
	// We know exactly 5 orders were placed, so we wait to pull exactly 5 results out.
	fmt.Println("\n=== DELIVERING FOOD ===")
	for i := 1; i <= totalOrders; i++ {
		completedOrderMessage := <-finishedFoodPipe
		fmt.Println(completedOrderMessage)
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
	fmt.Println("===============================")
	fmt.Println("===============================")
	fmt.Println("===============================")

	incomingOrdersPipe <- 6
	test := <-finishedFoodPipe
	fmt.Println(test)

	incomingOrdersPipe <- 7
	test1 := <-finishedFoodPipe
	fmt.Println(test1)

	incomingOrdersPipe <- 8
	test2 := <-finishedFoodPipe
	fmt.Println(test2)

}
