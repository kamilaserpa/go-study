# GO lang: Orientação a. Objetos

Documentação: https://go.dev/ref/spec </br>
Go playground: https://go.dev/play/

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

### Funções

Uma [declaração de função](https://go.dev/ref/spec#Function_declarations) utiliza a palavra "func", vincula um [identificador](https://go.dev/ref/spec#Identifiers) (o nome da função) parâmetros de tipo, assinatura, e o corpo da função.

```go
    func soma(a int, b int) int {
    return a + b
}
// "func" → palavra-chave obrigatória
// soma → nome da função
// Assinatura → (a int, b int) int
// CorpoDaFunção → { return a + b }
```

A função pode declarar parâmetors de resultado (retorno), nesse caso a fn deve terminar em uma [instrução de terminação](https://go.dev/ref/spec#Terminating_statements).
```go 
func dobrar(x int) int {
    return x * 2
}
```

Uma função pode ter múltiplos retornos.
Em Go, é muito comum que funções retornem dois valores: o resultado da operação e um erro (error). Esse padrão permite que o código seja seguro e explicite possíveis falhas sem usar exceções
```go
import (
    "fmt"
    "os"
)

func lerArquivo(nome string) (string, error) {
    dados, err := os.ReadFile(nome)
    if err != nil {
        return "", err
    }
    return string(dados), nil
}
func main() {
    conteudo, err := lerArquivo("exemplo.txt")
    if err != nil {
        fmt.Println("Erro ao ler o arquivo:", err)
        return
    }
    fmt.Println("Conteúdo do arquivo:")
    fmt.Println(conteudo)
}

```

Uma função pode ter um, muitos ou nenhum parâmetro:
```go
    func SemParametro() string {
        return "Exemplo de função sem parâmetro!"
    }

    func UmParametro(texto string) string {
        return texto
    }

    func DoisParametros(texto string, numero int) (string, int) {
        return texto, numero
    }

    func main() {
        fmt.Println(SemParametro())
        fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
        fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))
    }
```

#### Função anônima
Útil para funções inline ou passar como argumento.
```go
soma := func(a, b int) int {
    return a + b
}
fmt.Println(soma(2, 3))
```

#### Função variádica (número variável de argumentos)
Uma variadic function tem um tipo prefixado com `...`e pode ser invocada com zero ou mais argumentos para esse parâmetro.
```go
func Somando(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
}

func main() {
    fmt.Println(Somando(1))
    fmt.Println(Somando(1,1))
    fmt.Println(Somando(1,1,1))
    fmt.Println(Somando(1,1,2,4))
}
```

#### Função como valor (funções de ordem superior)
Recebe outra função como parâmetro:
```go
func aplicar(f func(int) int, valor int) int {
    return f(valor)
}

func dobrar(x int) int {
    return x * 2
}

resultado := aplicar(dobrar, 5) // resultado = 10
```

#### Métodos
Um método é uma função especial que tem um receptor (receiver), e esse receptor é o que o vincula a um tipo definido (geralmente uma struct).

 - Receiver
O receptor é o tipo ao qual o método está vinculado.
Funciona como o `this` em outras linguagens (embora em Go ele seja explícito).
Pode ser:
 - por valor → `func (p Pessoa)`, cópia da struct, mudanças não afetam o original
 - por ponteiro → `func (p *Pessoa)`, permite modificar os campos da struct diretamente, ou seja, mudanças afetam o objeto original

```go
type Pessoa struct {
    Nome string
    Idade int
}

// Método com receptor por valor
func (p Pessoa) Saudacao() string {
    return "Olá, meu nome é " + p.Nome
}
```

