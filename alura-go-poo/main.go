package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "Valor de depósito inválido", c.saldo

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}
func main() {

	contaGabriel := ContaCorrente{
		titular:       "Gabriel Soares",
		numeroAgencia: 1234,
		numeroConta:   5647,
		saldo:         200,
	}

	contaRenata := ContaCorrente{
		titular:       "Renata Panseri",
		numeroAgencia: 556,
		numeroConta:   489,
		saldo:         300,
	}

	//fmt.Println(contaGabriel.saldo)
	//fmt.Println(contaGabriel.Sacar(600))
	//fmt.Println(contaGabriel.saldo)
	//status, valor := contaGabriel.Depositar(654.56)
	//fmt.Println(status, "\nSaldo: $", valor)

	status := contaRenata.Transferir(300, &contaGabriel)

	fmt.Println(status)
	fmt.Println("Saldo Gabriel:", contaGabriel.saldo)
	fmt.Println("Saldo Renata:", contaRenata.saldo)

}
