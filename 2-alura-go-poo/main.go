package main

import (
	"fmt"
	"poo/clientes"
	"poo/contas"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {

	clienteGabriel := clientes.Titular{
		Nome:      "Gabriel Soares",
		CPF:       "31485698533",
		Profissao: "Engenheiro de Software",
	}

	contaGabriel := contas.ContaCorrente{
		Titular:       clienteGabriel,
		NumeroAgencia: 1234,
		NumeroConta:   5647,
	}

	contaGabriel.Depositar(200)
	fmt.Println(contaGabriel)
	PagarBoleto(&contaGabriel, 60)
	fmt.Println("Saldo => ", contaGabriel.MostraSaldo())
	//contaRenata := contas.ContaCorrente{
	//	Titular: clientes.Titular{
	//		Nome:      "Renata Panseri",
	//		CPF:       "12345678900",
	//		Profissao: "Manda Chuva",
	//	},
	//	NumeroAgencia: 556,
	//	NumeroConta:   489,
	//	Saldo:         300,
	//}
	//fmt.Println(contaRenata)

	//fmt.Println(contaGabriel.Saldo)
	//fmt.Println(contaGabriel.Sacar(600))
	//fmt.Println(contaGabriel.Saldo)
	//status, valor := contaGabriel.Depositar(654.56)
	//fmt.Println(status, "\nSaldo: $", valor)

	//status := contaRenata.Transferir(300, &contaGabriel)
	//
	//fmt.Println(status)
	//fmt.Println("Saldo Gabriel:", contaGabriel.Saldo)
	//fmt.Println("Saldo Renata:", contaRenata.Saldo)

}
