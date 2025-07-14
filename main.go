package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println("Saldo inicial da conta: ", contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Sacar(400.))
	fmt.Println("Saldo após saque: ", contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Depositar(-200))
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, "Saldo:", valor)

	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}
	statusTransferencia := contaDaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println(statusTransferencia)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
