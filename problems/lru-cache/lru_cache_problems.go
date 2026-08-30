package lru_cache

import "container/list"

// --- LRU Cache (LeetCode #146) ---

type lruEntry struct {
	key, val int
}

type LRUCache struct {
	cap     int
	list    *list.List
	hashMap map[int]*list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cap:     capacity,
		list:    list.New(),
		hashMap: make(map[int]*list.Element),
	}
}

// Get returns the value for key, or -1 if not found; moves item to front (MRU).
func (c *LRUCache) Get(key int) int {
	if el, ok := c.hashMap[key]; ok {
		c.list.MoveToFront(el)
		return el.Value.(*lruEntry).val
	}
	return -1
}

// Put inserts/updates key-value; evicts LRU item when over capacity.
func (c *LRUCache) Put(key, value int) {
	if el, ok := c.hashMap[key]; ok {
		el.Value.(*lruEntry).val = value
		c.list.MoveToFront(el)
		return
	}
	if c.list.Len() == c.cap {
		back := c.list.Back()
		c.list.Remove(back)
		delete(c.hashMap, back.Value.(*lruEntry).key)
	}
	el := c.list.PushFront(&lruEntry{key, value})
	c.hashMap[key] = el
}

// --- LFU Cache (LeetCode #460) ---

type lfuEntry struct {
	key, val, freq int
}

type LFUCache struct {
	cap     int
	minFreq int
	keyMap  map[int]*list.Element       // key -> *list.Element (lfuEntry)
	freqMap map[int]*list.List          // freq -> doubly linked list of lfuEntry
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		cap:     capacity,
		keyMap:  make(map[int]*list.Element),
		freqMap: make(map[int]*list.List),
	}
}

func (c *LFUCache) Get(key int) int {
	if el, ok := c.keyMap[key]; ok {
		c.incrementFreq(el)
		return el.Value.(*lfuEntry).val
	}
	return -1
}

func (c *LFUCache) Put(key, value int) {
	if c.cap == 0 {
		return
	}
	if el, ok := c.keyMap[key]; ok {
		el.Value.(*lfuEntry).val = value
		c.incrementFreq(el)
		return
	}
	if len(c.keyMap) == c.cap {
		// Evict least frequent (from minFreq list's back)
		lst := c.freqMap[c.minFreq]
		back := lst.Back()
		lst.Remove(back)
		delete(c.keyMap, back.Value.(*lfuEntry).key)
	}
	entry := &lfuEntry{key: key, val: value, freq: 1}
	if c.freqMap[1] == nil {
		c.freqMap[1] = list.New()
	}
	el := c.freqMap[1].PushFront(entry)
	c.keyMap[key] = el
	c.minFreq = 1
}

func (c *LFUCache) incrementFreq(el *list.Element) {
	entry := el.Value.(*lfuEntry)
	oldFreq := entry.freq
	entry.freq++
	// Remove from old freq list
	c.freqMap[oldFreq].Remove(el)
	if c.freqMap[oldFreq].Len() == 0 {
		delete(c.freqMap, oldFreq)
		if c.minFreq == oldFreq {
			c.minFreq++
		}
	}
	// Add to new freq list
	newFreq := entry.freq
	if c.freqMap[newFreq] == nil {
		c.freqMap[newFreq] = list.New()
	}
	newEl := c.freqMap[newFreq].PushFront(entry)
	c.keyMap[entry.key] = newEl
}

// --- Hit Counter (LeetCode #362) ---
// Counts hits in the past 300 seconds.

type HitCounter struct {
	times  [300]int
	counts [300]int
}

func NewHitCounter() *HitCounter {
	return &HitCounter{}
}

// Hit records a hit at the given timestamp (in seconds).
func (h *HitCounter) Hit(timestamp int) {
	idx := (timestamp - 1) % 300
	if h.times[idx] != timestamp {
		h.times[idx] = timestamp
		h.counts[idx] = 0
	}
	h.counts[idx]++
}

// GetHits returns the number of hits in the past 300 seconds (inclusive).
func (h *HitCounter) GetHits(timestamp int) int {
	total := 0
	for i := 0; i < 300; i++ {
		if timestamp-h.times[i] < 300 {
			total += h.counts[i]
		}
	}
	return total
}
