//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// --- 1. THE STANDARD MUTEX (Slow for readers) ---
type slowManager struct {
	mu   sync.Mutex
	data map[string]string
}

func (m *slowManager) ReadData(wg *sync.WaitGroup) {
	defer wg.Done()
	m.mu.Lock() // 🚪 Locks out EVERYONE (even other readers!)
	defer m.mu.Unlock()
	time.Sleep(1 * time.Millisecond) // Simulate 1ms of reading work
	_ = m.data["key"]
}

// --- 2. THE RWMUTEX (Lightning fast for readers) ---
type fastManager struct {
	mu   sync.RWMutex
	data map[string]string
}

func (m *fastManager) ReadData(wg *sync.WaitGroup) {
	defer wg.Done()
	m.mu.RLock() // 📖 Let ALL readers in at the exact same time!
	defer m.mu.RUnlock()
	time.Sleep(1 * time.Millisecond) // Simulate 1ms of reading work
	_ = m.data["key"]
}

// --- 3. NO MUTEX (The Crash Simulator) ---
func crashTest() {
	var wg sync.WaitGroup
	badMap := make(map[string]int)

	// 100 workers rushing the map with NO locks
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			badMap["crash"] = val // 🚨 FATAL ERROR: concurrent map writes
		}(i)
	}
	wg.Wait()
}

func main() {
	fmt.Println("🚀 Starting the Mutex Showdown!\n")

	// --- TEST 1: Standard Mutex ---
	slow := slowManager{data: map[string]string{"key": "value"}}
	var wg1 sync.WaitGroup

	startSlow := time.Now()
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go slow.ReadData(&wg1) // Spawn 1000 readers
	}
	wg1.Wait()

	fmt.Printf("🐢 Standard Mutex Time: %v (Readers waited in line)\n", time.Since(startSlow))

	// --- TEST 2: RWMutex ---
	fast := fastManager{data: map[string]string{"key": "value"}}
	var wg2 sync.WaitGroup

	startFast := time.Now()
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go fast.ReadData(&wg2) // Spawn 1000 readers
	}
	wg2.Wait()

	fmt.Printf("⚡ RWMutex Time:       %v (Readers read simultaneously!)\n\n", time.Since(startFast))

	// --- TEST 3: No Mutex (Uncomment to crash!) ---
	fmt.Println("🚨 Try uncommenting 'crashTest()' in the code to see the race condition panic!")
	// crashTest()
}
