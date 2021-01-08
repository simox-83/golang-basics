package main

import (
	"fmt"
)

func main() {
	//m:=make(map[string]int)
	m := map[string]int{"a": 1, "c": 3}
	m["b"] = 2
	fmt.Println(m)

	//var m1 map[string]string
	//m1["canI?"] = "NO!"

	key := "b"
	fmt.Printf("Il valore della chiave %s e' %d\n", key, m[key])
	key1 := "d"
	fmt.Printf("Il valore della chiave %s e' %d\n", key1, m[key1])
	v, ok := m[key1]
	if !ok {
		fmt.Printf("La chiave %s non esiste\n", key1)
	} else {
		fmt.Printf("La chiave %s esiste e il suo valore e' %d\n", key1, v)
	}

	for k, v := range m {
		fmt.Printf("k e' uguale a %s, v e' uguale a %d\n", k, v)
	}
	delete(m, "b")
	fmt.Println("Map m e' uguale a: ", m)

	n := m
	fmt.Println("Map n e' uguale a: ", n)

	if m != nil {
		fmt.Println("M non e' vuota")
	}

	m = map[string]int{"abc": 3, "b": 1}
	n = map[string]int{"b": 1, "abc": 3}

	if IsEqual(m, n) {
		fmt.Println("m e n sono uguali")

	} else {
		fmt.Println("m e n non sono uguali")
	}

	m1 := map[int][]string{1: {"A E I O U"}, 2: {"P D"}}
	fmt.Println(m1)

	for k, v := range m1 {
		for _, c := range v {
			fmt.Printf("%d, %s, %s\n", k, v, c)
		}
	}

}

func IsEqual(m, n map[string]int) bool {

	if len(m) != len(n) {
		return false
	}

	for k, v := range m {
		for k1, v1 := range n {
			if k == k1 && v != v1 {
				return false
			}

			_, ok := n[k]
			if !ok {
				return false
			}

		}
	}

	return true
}
