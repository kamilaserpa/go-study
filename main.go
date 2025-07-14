package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// Atribuição curta
	contaDaEugenia := ContaCorrente{
		titular:       "Eugênia",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.6}
	fmt.Println(contaDaEugenia)

	contaDaMaria := ContaCorrente{"Maria", 389, 234567, 147.5}
	fmt.Println(contaDaMaria)

	// Definindo apenas alguns parâmetros. O Go atribui um valor default inicial para as propriedades não definidas
	contaDaBruna := ContaCorrente{titular: "Bruna", saldo: 275.5}
	fmt.Println(contaDaBruna)
	contaDaBruna2 := ContaCorrente{titular: "Bruna", saldo: 275.5}
	fmt.Println(contaDaBruna2)
	fmt.Println("True, pois avalia os conteúdos, contaDaBruna == contaDaBruna2:")
	fmt.Println(contaDaBruna == contaDaBruna2)

	var contaDaPaula *ContaCorrente
	contaDaPaula = new(ContaCorrente)
	contaDaPaula.titular = "Paula"
	contaDaPaula.saldo = 500
	fmt.Println(contaDaPaula)

	var contaDaPaula2 *ContaCorrente
	contaDaPaula2 = new(ContaCorrente)
	contaDaPaula2.titular = "Paula"
	contaDaPaula2.saldo = 500
	fmt.Println(contaDaPaula2)

	fmt.Println("False, pois avalia os endereços, contaDaPaula == contaDaPaula2:")
	fmt.Println(contaDaPaula == contaDaPaula2) // False, pois avalia os valores

	fmt.Println("True, pois avalia os conteúdos, *contaDaPaula == *contaDaPaula2:")
	fmt.Println(*contaDaPaula == *contaDaPaula2) // True, pois avalia os conteúdos

}
