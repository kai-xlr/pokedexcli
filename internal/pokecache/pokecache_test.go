package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		name string
		key  string
		val  []byte
	}{
		{
			name: "simple URL with basic data",
			key:  "https://example.com",
			val:  []byte("testdata"),
		},
		{
			name: "URL with path and more complex data",
			key:  "https://example.com/path",
			val:  []byte("moretestdata"),
		},
		{
			name: "empty string key",
			key:  "",
			val:  []byte("empty key data"),
		},
		{
			name: "empty byte slice value",
			key:  "empty-value",
			val:  []byte{},
		},
		{
			name: "nil byte slice value",
			key:  "nil-value",
			val:  nil,
		},
		{
			name: "large data value",
			key:  "large-data",
			val:  make([]byte, 1024), // 1KB of zeros
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key %q", c.key)
				return
			}
			if len(val) != len(c.val) {
				t.Errorf("expected value length %d, got %d", len(c.val), len(val))
				return
			}
			for i := range val {
				if val[i] != c.val[i] {
					t.Errorf("value mismatch at index %d: expected %v, got %v", i, c.val[i], val[i])
					return
				}
			}
		})
	}
}

func TestCacheGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	// Test getting non-existent key
	val, ok := cache.Get("nonexistent")
	if ok {
		t.Errorf("expected key to not exist")
	}
	if val != nil {
		t.Errorf("expected nil value for non-existent key, got %v", val)
	}

	// Test getting existing key
	testKey := "test-key"
	testVal := []byte("test-value")
	cache.Add(testKey, testVal)

	val, ok = cache.Get(testKey)
	if !ok {
		t.Errorf("expected key to exist")
	}
	if string(val) != string(testVal) {
		t.Errorf("expected %q, got %q", string(testVal), string(val))
	}
}

func TestCacheOverwrite(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	key := "overwrite-test"
	firstVal := []byte("first-value")
	secondVal := []byte("second-value")

	// Add first value
	cache.Add(key, firstVal)
	val, ok := cache.Get(key)
	if !ok || string(val) != string(firstVal) {
		t.Errorf("expected first value %q, got %q (ok=%v)", string(firstVal), string(val), ok)
	}

	// Overwrite with second value
	cache.Add(key, secondVal)
	val, ok = cache.Get(key)
	if !ok || string(val) != string(secondVal) {
		t.Errorf("expected second value %q, got %q (ok=%v)", string(secondVal), string(val), ok)
	}
}

func TestCacheExpiry(t *testing.T) {
	const interval = 100 * time.Millisecond
	cache := NewCache(interval)

	key := "expiry-test"
	val := []byte("expires-soon")

	cache.Add(key, val)

	// Should exist immediately
	if _, ok := cache.Get(key); !ok {
		t.Errorf("expected key to exist immediately after adding")
	}

	// Wait for expiry and reap cycle
	time.Sleep(interval + 50*time.Millisecond) // Wait for reap to run

	// Should be expired and removed
	if _, ok := cache.Get(key); ok {
		t.Errorf("expected key to be expired and removed")
	}
}

func TestCacheReap(t *testing.T) {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mtx:   &sync.Mutex{},
	}

	now := time.Now().UTC()
	oldTime := now.Add(-1 * time.Hour)
	recentTime := now.Add(-1 * time.Minute)

	// Add entries with different ages
	cache.cache["old"] = cacheEntry{createdAt: oldTime, val: []byte("old-data")}
	cache.cache["recent"] = cacheEntry{createdAt: recentTime, val: []byte("recent-data")}

	// Reap entries older than 30 minutes
	cache.reap(now, 30*time.Minute)

	// Old entry should be removed, recent should remain
	if _, ok := cache.cache["old"]; ok {
		t.Errorf("expected old entry to be reaped")
	}
	if _, ok := cache.cache["recent"]; !ok {
		t.Errorf("expected recent entry to remain")
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	const interval = 5 * time.Second
	const numGoroutines = 100
	const numOperations = 10

	cache := NewCache(interval)
	var wg sync.WaitGroup

	// Test concurrent adds and gets
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				val := []byte(fmt.Sprintf("value-%d-%d", id, j))

				// Add
				cache.Add(key, val)

				// Get immediately
				if gotVal, ok := cache.Get(key); !ok {
					t.Errorf("expected to find key %q immediately after adding", key)
				} else if string(gotVal) != string(val) {
					t.Errorf("expected value %q, got %q", string(val), string(gotVal))
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestNewCache(t *testing.T) {
	interval := 1 * time.Second
	cache := NewCache(interval)

	// Verify cache is properly initialized
	if cache.cache == nil {
		t.Errorf("expected cache map to be initialized")
	}
	if cache.mtx == nil {
		t.Errorf("expected mutex to be initialized")
	}

	// Test that reap goroutine is running by adding an entry and waiting
	key := "reap-test"
	val := []byte("should-be-reaped")
	cache.Add(key, val)

	// Wait for more than one interval to ensure reap runs
	time.Sleep(interval + 100*time.Millisecond)

	// Entry should be reaped
	if _, ok := cache.Get(key); ok {
		t.Errorf("expected entry to be reaped by background goroutine")
	}
}
