package pokeapi 

import (
  "time"
  "sync"
)

type Cache struct {
  Locations map[string]cacheEntry
  Mux *sync.Mutex
}

type cacheEntry struct {
  CreatedAt time.Time
  Val []byte
}

func NewCache(interval time.Duration) Cache {
  c := Cache{
    Locations: make(map[string]cacheEntry),
    Mux: &sync.Mutex{},
  }
  go c.reapLoop(interval)
  return c
}

func (c *Cache) Add(key string, val []byte) {
  entry := cacheEntry {
    CreatedAt: time.Now(),
    Val: val,
  }
  c.Mux.Lock()
  c.Locations[key] = entry
  c.Mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool){
  c.Mux.Lock()
  entry, ok := c.Locations[key]
  c.Mux.Unlock()
  return entry.Val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
  //Remove entries older than interval
  ticker := time.NewTicker(interval)
  defer ticker.Stop()
  for range ticker.C {
    for k, v := range c.Locations {
      c.Mux.Lock()
      if v.CreatedAt.Add(interval).After(time.Now()){
        delete(c.Locations, k)
      }
      c.Mux.Unlock()
    }
  }
}
