package contas

import "alura/curso-1-orientacao-objeto/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 00
	if podeSacar {
		c.saldo -= valorDoSaque
		return "> Saque realizado com sucesso."
	} else {
		return "> saldo insuficiente."
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "> Depósito realizado com sucesso.", c.saldo
	} else {
		return "> Valor de depósito inválido.", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
