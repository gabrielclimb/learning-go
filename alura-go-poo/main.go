package main

import (
	"fmt"
	"poo/contas"
)

func main() {

	contaGabriel := contas.ContaCorrente{
		Titular:       "Gabriel Soares",
		NumeroAgencia: 1234,
		NumeroConta:   5647,
		Saldo:         200,
	}

	contaRenata := contas.ContaCorrente{
		Titular:       "Renata Panseri",
		NumeroAgencia: 556,
		NumeroConta:   489,
		Saldo:         300,
	}

	//fmt.Println(contaGabriel.Saldo)
	//fmt.Println(contaGabriel.Sacar(600))
	//fmt.Println(contaGabriel.Saldo)
	//status, valor := contaGabriel.Depositar(654.56)
	//fmt.Println(status, "\nSaldo: $", valor)

	status := contaRenata.Transferir(300, &contaGabriel)

	fmt.Println(status)
	fmt.Println("Saldo Gabriel:", contaGabriel.Saldo)
	fmt.Println("Saldo Renata:", contaRenata.Saldo)

}
