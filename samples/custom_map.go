package samples

import (
	"fmt"
	"sync"
)

// KeyValuePair represents a key-value pair
type KeyValuePair struct {
	Key   string
	Value int
}

// CustomMap represents a custom map-like data structure
type CustomMap struct {
	mu    sync.RWMutex
	pairs []KeyValuePair
}

// NewCustomMap creates a new instance of CustomMap
func NewCustomMap() *CustomMap {
	return &CustomMap{
		pairs: make([]KeyValuePair, 0),
	}
}

// Set adds or updates a key-value pair in the map
func (m *CustomMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the key already exists
	for i := range m.pairs {
		if m.pairs[i].Key == key {
			// Update the existing value
			m.pairs[i].Value = value
			return
		}
	}

	// Add a new key-value pair
	m.pairs = append(m.pairs, KeyValuePair{Key: key, Value: value})
}

// Get retrieves the value associated with the given key from the map
func (m *CustomMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Find the key-value pair
	for _, pair := range m.pairs {
		if pair.Key == key {
			return pair.Value, true
		}
	}

	// Key not found
	return 0, false
}

// Delete removes the key-value pair associated with the given key from the map
func (m *CustomMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Find the index of the key-value pair
	for i, pair := range m.pairs {
		if pair.Key == key {
			// Remove the pair from the slice
			m.pairs = append(m.pairs[:i], m.pairs[i+1:]...)
			return
		}
	}
}

func mainCustomMap() {
	// Create a new instance of CustomMap
	customMap := NewCustomMap()

	// Set key-value pairs in the map
	customMap.Set("key1", 10)
	customMap.Set("key2", 20)

	// Get values from the map
	value1, exists1 := customMap.Get("key1")
	value2, exists2 := customMap.Get("key2")
	value3, exists3 := customMap.Get("key3")

	fmt.Println("Value1:", value1, "Exists1:", exists1)
	fmt.Println("Value2:", value2, "Exists2:", exists2)
	fmt.Println("Value3:", value3, "Exists3:", exists3)

	// Delete a key from the map
	customMap.Delete("key2")

	// Get the updated value after deletion
	value2, exists2 = customMap.Get("key2")
	fmt.Println("Value2:", value2, "Exists2:", exists2)
}
