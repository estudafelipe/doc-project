package main

import (
	"github.com/estudafelipe/doc-project/cachorro"
	"fmt"
)

func main() {
	idadeDaChaira := cachorro.Idade(8)
	fmt.Printf("Idade da Chaira %d\n", idadeDaChaira)
}