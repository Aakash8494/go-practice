package main

import (
	"container/list"
	"sync"
	"time"
)

// 1. The Data Container (What goes inside the list)
type cacheItem struct {
	key       string
	value     string
	expiresAt time.Time // The TTL expiration stamp!
}

// 2. The Manager
type LRUCache struct {
	sync.Mutex                          // 🚪 Our steel door!
	capacity   int                      // Max items allowed
	ttl        time.Duration            // How long items live
	items      map[string]*list.Element // The Phonebook (Points to list items)
	lineup     *list.List               // The Order (Tracks who is newest/oldest)
}

// Constructor
func NewLRUCache(capacity int, ttl time.Duration) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ttl:      ttl,
		items:    make(map[string]*list.Element),
		lineup:   list.New(),
	}
}

// --- WRITE DATA ---
func (c *LRUCache) Set(key string, value string) {
	c.Lock() // 🚪 Lock the door!
	defer c.Unlock()

	// If it already exists, update it and move it to the front of the line
	if element, exists := c.items[key]; exists {
		c.lineup.MoveToFront(element) // "You were just used, go to the front!"
		item := element.Value.(*cacheItem)
		item.value = value
		item.expiresAt = time.Now().Add(c.ttl) // Reset the expiration clock
		return
	}

	// If it's brand new, put it at the front of the line
	newItem := &cacheItem{
		key:       key,
		value:     value,
		expiresAt: time.Now().Add(c.ttl), // Stamp it with an expiration date
	}
	element := c.lineup.PushFront(newItem)
	c.items[key] = element

	// 🚨 LRU ENFORCEMENT: If we are over capacity, delete the guy at the very back!
	if c.lineup.Len() > c.capacity {
		oldest := c.lineup.Back()
		if oldest != nil {
			c.lineup.Remove(oldest) // Remove from the list
			oldItem := oldest.Value.(*cacheItem)
			delete(c.items, oldItem.key) // Remove from the map
		}
	}
}

// --- READ DATA ---
func (c *LRUCache) Get(key string) (string, bool) {
	c.Lock() // 🚪 Lock the door!
	defer c.Unlock()

	element, exists := c.items[key]
	if !exists {
		return "", false
	}

	item := element.Value.(*cacheItem)

	// 🚨 TTL ENFORCEMENT: Check if the milk is expired!
	if time.Now().After(item.expiresAt) {
		// It's expired! Delete it and pretend it doesn't exist.
		c.lineup.Remove(element)
		delete(c.items, key)
		return "", false
	}

	// If it's valid, move it to the front of the line because it was just used!
	c.lineup.MoveToFront(element)
	return item.value, true
}
