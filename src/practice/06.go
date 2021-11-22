package main

import (
	"fmt"
	"sort"
)

/**
map/
*/
func main() {
	var map1 map[string]string
	fmt.Println(map1)
	fmt.Println(map1 == nil)
	map2 := make(map[string]int)
	fmt.Println(map2)
	fmt.Println(map2 == nil)
	map3 := make(map[string]string, 2)
	fmt.Println(len(map3))
	map4 := map[string]string{"a": "a", "b": "b"}
	map4["c"] = "c"
	fmt.Println(map4)
	fmt.Println(map4["c"])
	for k, v := range map4 {
		fmt.Printf("key:%s,value:%s\n", k, v)
	}
	v, ok := map4["a"]
	if ok {
		fmt.Println("exist", v)
	}
	// 排序
	// 1、将key放入切片
	var keys []string
	for k, _ := range map4 {
		keys = append(keys, k)
	}
	// 2、对切片进行排序
	sort.Strings(keys)
	// 3、通过排序好的切片从map中取有序的数值
	for _, k := range keys {
		fmt.Printf("key:%s, value:%s\n", k, map4[k])
	}
}
