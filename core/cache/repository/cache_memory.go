package repository

import (
	"context"
	"sync"
	"time"
)

type CacheItem struct {
	value     []byte
	expiresAt time.Time
}

type CacheMemory struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

func NewCacheMemory() *CacheMemory {
	return &CacheMemory{
		items: make(map[string]CacheItem),
	}
}

func (m *CacheMemory) Ping(ctx context.Context) error {
	return nil
}

func (m *CacheMemory) Store(ctx context.Context, key string, value []byte, exp time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	expiresAt := time.Now().Add(exp)
	m.items[key] = CacheItem{
		value:     value,
		expiresAt: expiresAt,
	}
	return nil
}

func (m *CacheMemory) Get(ctx context.Context, key string) ([]byte, bool, error) {
	m.mu.RLock()
	item, exists := m.items[key]
	m.mu.RUnlock() // Release read lock before modifying map

	if !exists {
		return nil, false, nil
	}

	// Check if item has expired
	if !item.expiresAt.IsZero() && time.Now().After(item.expiresAt) {
		m.mu.Lock() // Acquire write lock before modifying map
		delete(m.items, key)
		m.mu.Unlock()
		return nil, false, nil
	}

	return item.value, true, nil
}

func (m *CacheMemory) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Clear all items
	m.items = make(map[string]CacheItem)
	return nil
}
