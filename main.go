package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type Conta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	fmt.Println("Saldo após depósito:", contaDoDenis.ObterSaldo())
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println("Saldo após boleto:", contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(200)
	fmt.Println("Saldo após depósito:", contaDaLuiza.ObterSaldo())
	PagarBoleto(&contaDaLuiza, 40)
	fmt.Println("Saldo após boleto:", contaDaLuiza.ObterSaldo())

}
