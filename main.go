package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 00
	if podeSacar {
		c.saldo -= valorDoSaque
		return "> Saque realizado com sucesso."
	} else {
		return "> Saldo insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "> Depósito realizado com sucesso.", c.saldo
	} else {
		return "> Valor de depósito inválido.", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println("Saldo inicial da conta: ", contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400.))
	fmt.Println("Saldo após saque: ", contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Depositar(-200))
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, "Saldo:", valor)

	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}
	statusTransferencia := contaDaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println(statusTransferencia)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
