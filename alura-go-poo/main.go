package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func NewContaCorrente() *ContaCorrente {
	return &ContaCorrente{}
}

func main() {

	contaGabriel := ContaCorrente{
		titular:       "Gabriel",
		numeroAgencia: 2656,
		numeroConta:   5458,
		saldo:         25898.56,
	}

	fmt.Println(contaGabriel)
}
