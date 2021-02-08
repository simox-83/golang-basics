package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	fmt.Printf("%v, %T\n", a, a)
	a[0] = 5
	fmt.Printf("%d\n", a)
	b := [5]int{10}
	fmt.Printf("%d\n", b)

	c := [...]string{"a", "b", "c"}
	fmt.Printf("%T, %s\n", c, c)

	//b = a. No! Sono 2 diversi tipi di dato

	modifyElem(c)
	fmt.Printf("L'array c contiene %s\n", c) // Non cambia il valore perche' gli array sono passati per valore. E' una copia!

	for i := 0; i < len(a); i++ {
		a[i] *= 2
	}
	fmt.Println("L'array a ora contiene", a)

	for i, n := range b {
		n++
		b[i] = n
	}
	fmt.Println("L'array b ora contiene", b)

	t := [2][3]string{{"first", "row", "data"}, {"second", "row", "different"}}
	fmt.Printf("%s\n", t)

	for i, r := range t {
		for j, c := range r {
			c = "changeme"
			r[j] = c

		}
		t[i] = r
	}
	fmt.Printf("%s\n", t)

	s := b[:]
	fmt.Printf("%T,%v\n", s, s) // s e' una slice di interi che contiene i valori di b

	s[0] = 10
	fmt.Println("s contiene ", s, "b contiene", b)

	for i := range s {
		s[i] *= 3
	}
	fmt.Println("s contiene ", s, "b contiene", b)

	s = b[1:4]
	fmt.Printf("s contiene %d, la lunghezza di s e' %d, la capacita' di s e' %d\n", s, len(s), cap(s))

	z := make([]string, 0, len(s))
	fmt.Printf("%d, %d\n", len(z), cap(z))
	z = append(z, "sto appendendo dei valori con append", "speriamo", "funzioni", "continuo", "ad", "appendere")
	fmt.Println("z ora contiene", z)
	fmt.Printf("%d, %d\n", len(z), cap(z))
	modifyElemSlice(z)
	fmt.Println("z ora contiene", z)
}

func modifyElem(m [3]string) {
	m[0] = "newvalue"
}

func modifyElemSlice(m []string) {
	m[0] = "newvalue"
}
