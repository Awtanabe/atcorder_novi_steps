package main

import "fmt"

type TimeMap struct {
	data map[string][]pair
}

type pair struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{data: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	d := pair{}
	d.timestamp = timestamp
	d.value = value
	this.data[key] = append(this.data[key], d)
}

func (this *TimeMap) Get(key string, timestamp int) string {

	if v, ok := this.data[key]; ok {
		for i := 0; i < len(v); i++ {
			if v[i].timestamp == timestamp {
				return v[i].value
			}
		}

		if len(v) > 0 {
			for i := len(v) - 1; i >= 0; i-- {
				if timestamp >= v[i].timestamp {
					return v[i].value
				}
			}
		}
	}
	return ""
}

// ["TimeMap", "set", ["alice", "happy", 1], "get", ["alice", 1], "get", ["alice", 2], "set", ["alice", "sad", 3], "get", ["alice", 3]]
func main() {

	t := Constructor()
	t.Set("alice", "happy", 1)
	fmt.Println(t.Get("alice", 1))
	fmt.Println(t.Get("alice", 2))
	t.Set("alice", "sad", 3)
	fmt.Println(t.Get("alice", 3))

}
