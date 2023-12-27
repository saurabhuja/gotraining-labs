package main

import (
	"fmt"
	"sort"
)

func main() {
	var m1 = make(map[string]string)
	m1 = map[string]string{"RJ": "RAJASTHAN", "KA": "KARNATAKA"}
	m2 := map[string]string{"RJ": "RAJASTHAN", "KA": "KARNATAKA"}
	fmt.Println(m1, &m2)
	var p1 = &m1
	fmt.Println(p1, &m1, &p1)
	for k, v := range *p1 {
		fmt.Println(k, v)
	}

	keys := make([]string, len(m1))
	values := make([]string, len(m1))
	i := 0
	for k := range *p1 {
		keys[i] = k
		i++
	}
	i = 0
	for _, v := range *p1 {
		values[i] = v
		i++
	}
	sort.Strings(values)

	for j := range keys {
		fmt.Println(j, keys[j], m1[keys[j]], (*p1)[keys[j]])
	}

	fmt.Println(keys)
	fmt.Println(values)

	delete(*p1, "k1")
	fmt.Println(p1, &m1, &p1)

	slice := []string{"apple", "banana", "orange", "date"}
	result := convertToMap(slice)
	fmt.Println(result)
}

func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	var price float64
	price = 100 / float64(len(items))
	for x := range items {
		result[items[x]] = price
	}
	// Your code goes here
	return result
}
