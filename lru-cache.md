# 23 - Cache Design Patterns (LRU & LFU)

## 23.1 - Overview & Theoretical Foundations (CLRS Chapter 10 & 11)

* A **Least Recently Used (LRU) Cache** organizes items in order of access history, evicting the least recently accessed item when capacity is exceeded.
* **Composite Data Structure Pattern:**
  * Direct lookup in a Doubly Linked List is $\mathcal{O}(n)$.
  * Direct removal/reordering in a Hash Table is undefined (no sequential order).
  * By combining a **Hash Table** (mapping `key -> Node*` in $\mathcal{O}(1)$) with a **Doubly Linked List** (performing $\mathcal{O}(1)$ node splice, insertion, and deletion), both `get` and `put` execute in strict $\mathcal{O}(1)$ worst-case time.
* **Sentinel Nodes (CLRS 10.2):**
  * Using dummy `head` and `tail` sentinel nodes eliminates edge cases when inserting at the boundary or removing the last remaining element.

---

## 23.2 - Classic Example: LRU Cache Implementation

### Java Implementation

```java
import java.util.HashMap;
import java.util.Map;

public class LRUCache {

    private static class Node {
        int key, value;
        Node prev, next;
        Node(int key, int value) {
            this.key = key;
            this.value = value;
        }
    }

    private final int capacity;
    private final Map<Integer, Node> map;
    private final Node head, tail; // Sentinel boundary nodes

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.map = new HashMap<>();
        this.head = new Node(0, 0);
        this.tail = new Node(0, 0);
        head.next = tail;
        tail.prev = head;
    }

    public int get(int key) {
        if (!map.containsKey(key)) return -1;
        Node node = map.get(key);
        moveToHead(node); // Update access recency
        return node.value;
    }

    public void put(int key, int value) {
        if (map.containsKey(key)) {
            Node node = map.get(key);
            node.value = value;
            moveToHead(node);
        } else {
            if (map.size() >= capacity) {
                // Evict least recently used (node before tail)
                Node lru = tail.prev;
                removeNode(lru);
                map.remove(lru.key);
            }
            Node newNode = new Node(key, value);
            map.put(key, newNode);
            addNodeToHead(newNode);
        }
    }

    private void addNodeToHead(Node node) {
        node.next = head.next;
        node.prev = head;
        head.next.prev = node;
        head.next = node;
    }

    private void removeNode(Node node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }

    private void moveToHead(Node node) {
        removeNode(node);
        addNodeToHead(node);
    }
}
```

---

### Go Implementation

```go
package main

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cacheMap   map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	node, exists := c.cacheMap[key]
	if !exists {
		return -1
	}
	c.moveToHead(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cacheMap[key]; exists {
		node.value = value
		c.moveToHead(node)
	} else {
		if len(c.cacheMap) >= c.capacity {
			// Evict LRU
			lru := c.tail.prev
			c.removeNode(lru)
			delete(c.cacheMap, lru.key)
		}
		newNode := &Node{key: key, value: value}
		c.cacheMap[key] = newNode
		c.addNodeToHead(newNode)
	}
}

func (c *LRUCache) addNodeToHead(node *Node) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addNodeToHead(node)
}
```

---

## 23.3 - Time & Space Complexity Analysis

* **`get(key)`:** $\mathcal{O}(1)$ constant time.
* **`put(key, value)`:** $\mathcal{O}(1)$ constant time.
* **Space Complexity:** $\mathcal{O}(\text{Capacity})$ to maintain the Hash Table entries and doubly linked list nodes.

---

## 23.4 - Classic LeetCode & System Design Benchmarks

### 1. LRU Cache (LeetCode #146)
* **Problem Statement**: Design a data structure that follows the constraints of a Least Recently Used (LRU) cache. Implement the `LRUCache` class with `get` and `put` methods.
* **Solution Link**: [problems/lru-cache/lru_cache_problems.go](problems/lru-cache/lru_cache_problems.go) (`LRUCache`)
* **Explanation**: Uses a combination of a Hash Map for $\mathcal{O}(1)$ lookup and a Doubly Linked List to maintain the order of elements based on their access recency.
* **Conceptual Link**: Demonstrates the **Composite Data Structure Pattern**, where multiple data structures are combined to achieve optimal time complexity for different operations.

### 2. LFU Cache (LeetCode #460)
* **Problem Statement**: Design and implement a data structure for a Least Frequently Used (LFU) cache. It should support `get` and `put` operations. When the cache is full, it should evict the least frequently used item.
* **Solution Link**: [problems/lru-cache/lru_cache_problems.go](problems/lru-cache/lru_cache_problems.go) (`LFUCache`)
* **Explanation**: Maintains two maps: one for keys to node elements and another for frequencies to doubly linked lists of nodes with that frequency. It also tracks the `minFreq` to facilitate $\mathcal{O}(1)$ eviction.
* **Conceptual Link**: An extension of the LRU pattern that incorporates frequency-based eviction logic using multiple linked lists.

### 3. Design Hit Counter (LeetCode #362)
* **Problem Statement**: Design a hit counter which counts the number of hits received in the past 5 minutes (i.e., the past 300 seconds).
* **Solution Link**: [problems/lru-cache/lru_cache_problems.go](problems/lru-cache/lru_cache_problems.go) (`HitCounter`)
* **Explanation**: Uses two arrays of size 300 to store timestamps and hit counts. The index is determined by `timestamp % 300`, effectively creating a circular buffer/sliding window.
* **Conceptual Link**: Related to cache eviction and sliding window concepts, where data outside a certain time range is effectively "evicted" or ignored.

### 4. Design In-Memory File System (LeetCode #588)
* **Problem Statement**: Design an in-memory file system that supports basic operations like `ls`, `mkdir`, `addContentToFile`, and `readContentFromFile`.
* **Conceptual Link**: Often involves complex tree structures and can use LRU caching for performance optimization of frequently accessed files or directories.

---

## 23.5 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 10: Elementary Data Structures (Doubly linked lists & Sentinels pp. 256–264)
  * Chapter 11: Hash Tables (pp. 272–301)
* https://leetcode.com/problems/lru-cache/
* https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU
