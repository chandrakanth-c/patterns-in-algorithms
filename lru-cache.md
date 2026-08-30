# 23 - Cache Design Patterns (LRU & LFU)

## 23.1 - Overview

* **LRU (Least Recently Used) Cache:** Evicts the item that has not been accessed for the longest time when maximum capacity is reached.
  * Achieved in $\mathcal{O}(1)$ time by combining a **Hash Map** (for $\mathcal{O}(1)$ key lookup) with a **Doubly Linked List** (for $\mathcal{O}(1)$ node addition and removal).
* **LFU (Least Frequently Used) Cache:** Evicts the item with the minimum access frequency count.

---

## 23.2 - Classic Example: LRU Cache Implementation

### Java Implementation

```java
import java.util.HashMap;
import java.util.Map;

public class LRUCache {

    static class Node {
        int key, value;
        Node prev, next;
        Node(int key, int value) {
            this.key = key;
            this.value = value;
        }
    }

    private final int capacity;
    private final Map<Integer, Node> map;
    private final Node head, tail; // Dummy sentinel nodes

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
        moveToHead(node); // Mark as most recently used
        return node.value;
    }

    public void put(int key, int value) {
        if (map.containsKey(key)) {
            Node node = map.get(key);
            node.value = value;
            moveToHead(node);
        } else {
            if (map.size() >= capacity) {
                // Evict LRU node from tail
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

## 23.3 - Time & Space Complexity

* **`get(key)`:** $\mathcal{O}(1)$ time.
* **`put(key, value)`:** $\mathcal{O}(1)$ time.
* **Space Complexity:** $\mathcal{O}(\text{Capacity})$ to store elements in Hash Map and Doubly Linked List.

---

## 23.4 - Classic LeetCode Problems

* **LRU Cache** (LeetCode #146)
* **LFU Cache** (LeetCode #460)
* **Design In-Memory File System** (LeetCode #588)
* **Design Hit Counter** (LeetCode #362)

---

## 23.5 - Sources used for this file:
https://leetcode.com/problems/lru-cache/ <br>
https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU <br>
https://www.geeksforgeeks.org/lru-cache-implementation/ <br>
https://techinterviewhandbook.org/algorithms/hash-table/
