package main

// type LRUCache struct {
//     capacity int
//     data map[int][]int
// }

// func Constructor(capacity int) LRUCache {
//     return LRUCache{capacity: capacity, data: make(map[int][]int)}
// }

// func (this *LRUCache) Get(key int) int {
//     if v, ok := this.data[key]; ok {
//         v[1]++
//         return v[0]
//     } else {
//         return -1
//     }
// }

// func (this *LRUCache) Put(key int, value int) {
//     if v, ok := this.data[key]; ok {
//         v[0] = value
//     } else {
//         if len(this.data) + 1 > this.capacity {
//         min := 1000
//         key := -1
//         for k,v := range this.data {
//             if v[1] < min {
//                 min = v[1]
//                 key = k
//             }
//         }
//         delete(this.data, key)

//         }

//         this.data[key] = append(this.data[key], value)
//     }
// }

type LRUCache struct {
	cache    [][2]int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make([][2]int, 0),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	for i := range this.cache {
		if this.cache[i][0] == key {
			tmp := this.cache[i]
			this.cache = append(this.cache[:i], this.cache[i+1:]...)
			this.cache = append(this.cache, tmp)
			return tmp[1]
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	for i := range this.cache {
		if this.cache[i][0] == key {
			tmp := this.cache[i]
			this.cache = append(this.cache[:i], this.cache[i+1:]...)
			tmp[1] = value
			this.cache = append(this.cache, tmp)
			return
		}
	}

	if len(this.cache) == this.capacity {
		this.cache = this.cache[1:]
	}

	this.cache = append(this.cache, [2]int{key, value})
}
