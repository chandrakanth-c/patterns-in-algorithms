package lru_cache

import "testing"

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if got := c.Get(1); got != 1 {
		t.Errorf("LRUCache.Get(1): got %d, want 1", got)
	}
	c.Put(3, 3) // evicts key 2
	if got := c.Get(2); got != -1 {
		t.Errorf("LRUCache.Get(2): got %d, want -1 (evicted)", got)
	}
	c.Put(4, 4) // evicts key 1
	if got := c.Get(1); got != -1 {
		t.Errorf("LRUCache.Get(1): got %d, want -1 (evicted)", got)
	}
	if got := c.Get(3); got != 3 {
		t.Errorf("LRUCache.Get(3): got %d, want 3", got)
	}
	if got := c.Get(4); got != 4 {
		t.Errorf("LRUCache.Get(4): got %d, want 4", got)
	}
}

func TestLFUCache(t *testing.T) {
	c := NewLFUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if got := c.Get(1); got != 1 { // freq(1)=2, freq(2)=1
		t.Errorf("LFUCache.Get(1): got %d, want 1", got)
	}
	c.Put(3, 3) // evicts key 2 (lowest freq)
	if got := c.Get(2); got != -1 {
		t.Errorf("LFUCache.Get(2): got %d, want -1 (evicted)", got)
	}
	if got := c.Get(3); got != 3 {
		t.Errorf("LFUCache.Get(3): got %d, want 3", got)
	}
}

func TestHitCounter(t *testing.T) {
	hc := NewHitCounter()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	if got := hc.GetHits(4); got != 3 {
		t.Errorf("HitCounter.GetHits(4): got %d, want 3", got)
	}
	hc.Hit(300)
	if got := hc.GetHits(300); got != 4 {
		t.Errorf("HitCounter.GetHits(300): got %d, want 4", got)
	}
	if got := hc.GetHits(301); got != 1 { // hit@1 is now outside the 300s window
		t.Errorf("HitCounter.GetHits(301): got %d, want 1", got)
	}
}
