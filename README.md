# GO lang: Orientação a. Objetos

Documentação: https://go.dev/ref/spec

#### Declaração curta `:=`
O `:=` é chamado de declaração curta (short variable declaration), e ele:
 - Declara uma nova variável
 - Atribui um valor a ela
 - Deduz automaticamente o tipo
Você não pode usar := fora de funções, como no escopo global.
O `=` indica apenas atribuição (variável já existe): `x = 20`

#### Inicialização zero e nil

Mesmo não provendo nenhum valor, o Go garante inicializar todas as variáveis, conforme a tabela abaixo:

| Zero initialization|  |
| ------------------ |:-----:|
| bool               | false |
| int                | 0     |
| float              | 0     |
| string             | ""    |
| struct             | {}    |

Ao atribuir o valor `nil` para uma variável devemos atrobuir um tipo para ela, caso contrário receberemos o erro `use of untyped nil`, que significa uso não tipado do nil. O compilador não sabe se esta variável é um inteiro, uma string, um array ou uma structure.

```go
func main() {
    var s *string = nil
    fmt.Println(s)
}
// Dessa maneira o programa compila e retorna <nil> como esperado.
```

Ao comparar variáveis não inicializadas de tipos diferentes recebemos erro.

```go
func main() {
    var f float64
    var i int 

    fmt.Println(f==i)
}
```

Recebemos uma mensagem com um erro informando que os tipos são incompatíveis. Não podemos comparar o valor atribuído pela inicialização zero se temos tipos diferentes.

### Type

A palavra-chave type é usada para definir um novo tipo em Go. Ela serve como base para declarar:
 - Tipos personalizados (como NumeroDaConta int)
 - Structs
 - Interfaces
 - Aliases de tipos

```go
 type NumeroDaConta int
 var x NumeroDaConta = 12340 // NumeroDaConta é uma nova variável baseada no tipo int.
```

### Struct
`struct` é uma estrutura de dados composta usada para agrupar diferentes campos sob um mesmo tipo. Sintaxe geral:
```go
type Pessoa struct {
    Nome  string
    Idade int
}
```

#### Formas de inicializar uma struct:

1. Literal Com nome dos campos: mais clara e recomendada, a ordem dos campos não importa
```go
    p := Pessoa{Nome: "João", Idade: 30}
```

2. Literal Sem nome dos campos: aordem dos camampos deve seguir a definição da struct
```go
    p := Pessoa{"João", 30}
```

3. Declaração e Atribuição Campo a Campo:
```go
    var p Pessoa
    p.Nome = "João"
    p.Idade = 30
```

4. Ponteiro para struct: útil quando você precisa modificar o valor ssm cópia, p é do tipo `*Pessoa`

Ponteiro para uma struct usando a função `new` e depois inicializa seus campos.
```go
    var pessoaAna *Pessoa
    pessoaAna = new(Pessoa)
    pessoaAna.Nome = "Ana"
    pessoaAna.Idade = 28
```

Alternativa mais compacta seria:
```go
    p := &Pessoa{Nome: "Ana", Idade: 30}
```
