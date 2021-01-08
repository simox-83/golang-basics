package main

import (
	"fmt"
)

func main() {
	a := 100
	var b *int
	b = &a
	fmt.Printf("%T, %v, %d\n", b, b, a)

	var c *int
	if c == nil {
		fmt.Println("c punta a nil")
	}

	d := new(int)
	fmt.Printf("d punta ad una variabile con valore %d\n", *d)
	*b++
	fmt.Printf("b punta ad una variabile con valore %d, mentre a ha valore %d\n", *b, a)
	c = b
	fmt.Printf("%d\n", *c)

	change(&a)
	fmt.Printf("Ora a vale %d\n", a)

	e := new(int)
	e = addressreturn()
	fmt.Printf("Il valore di e e' %d\n", *e)

	s := []string{"a", "b", "c"}
	modify(s)
	fmt.Printf("Lo slice risultante dopo modify e' %s\n", s)

	arr := [3]string{"p", "d", "m"}
	modify(arr[:])
	fmt.Printf("Lo slice risultante dopo modify e' %s\n", arr)

	str := "changeme"
	modifystr(&str)
	fmt.Printf("La stringa risultante dopo modifystr e' %s\n", str)

}

func change(p *int) {
	*p = 10
}

func addressreturn() *int {
	i := 5
	return &i
}

func modify(slc []string) {
	slc[0] = "d"
}

func modifystr(s *string) {
	*s = "I changed you"
}
