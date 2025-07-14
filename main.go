package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.456.789-00",
		Profissao: "Desenvolvedor Go"}

	contaDoBruno := contas.ContaCorrente{
		Titular:       clienteBruno,
		NumeroAgencia: 123,
		NumeroConta:   123456}

	fmt.Println(contaDoBruno)

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	fmt.Println("Saldo:", contaExemplo.ObterSaldo())

}
