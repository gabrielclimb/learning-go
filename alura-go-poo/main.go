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

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {

	contaGabriel := NewContaCorrente()
	contaGabriel.titular = "Gabriel"
	contaGabriel.numeroAgencia = 546
	contaGabriel.numeroConta = 123
	contaGabriel.saldo = 5465.58

	fmt.Println(contaGabriel.saldo)

	fmt.Println(contaGabriel.Sacar(600))
	fmt.Println(contaGabriel.saldo)
}
