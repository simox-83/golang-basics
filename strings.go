package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, playground")
	s := "Señor"
	//s[0]="c" --> non funziona perche' le stringhe sono immutabili in Go
	fmt.Printf("%s\n", s)

	for i, c := range s {
		fmt.Printf("Indice %d corrisponde al carattere %c\n", i, c)
		fmt.Printf("Tipo di indice %T, tipo di valore %T\n", i, c)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("Indice %d corrisponde al carattere %c\n", i, s[i])
		fmt.Printf("Tipo di indice %T, tipo di valore %T\n", i, s[i])
	}

	conta := 0
	for _, c := range s {
		if c == 'ñ' {
			conta++
		}
	}
	fmt.Printf("Ci sono %d 'ñ' nella stringa\n", conta)

	conta = 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'ñ' {
			conta++
		}

	}
	fmt.Printf("Ci sono %d 'ñ' nella stringa\n", conta) // Ce ne sono zero! s[i] e' una rune, non un carattere...

	r := ""

	for i, c := range s {
		if i%2 == 0 {
			r += string(c)

		}
	}

	fmt.Printf("La stringa risultante e' %s, di tipo %T\n", r, r)

	arr := make([]rune, len(s))
	j := 0
	for i, c := range s {
		if i%2 == 0 {
			arr[j] = c
			j++
		}
	}

	fmt.Printf("La stringa risultante e' %s, di tipo %T\n", string(arr), string(arr))
}
