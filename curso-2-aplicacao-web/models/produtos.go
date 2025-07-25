package models

import (
	"alura/curso-2-aplicacao-web/db"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("Select * from produtos")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Error querying products:", err)
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

// CriarProdutoDb insere um novo produto no banco de dados com os parâmetros fornecidos.
// Recebe o nome, descrição, preço e quantidade do produto.
// Em caso de erro ao preparar ou executar a inserção, registra o erro no log.
func CriarProdutoDb(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	criarProdutoInstrucao, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")

	if err != nil {
		log.Println("Erro ao preparar a instrução:", err)
		panic(err.Error())
	}

	_, err = criarProdutoInstrucao.Exec(nome, descricao, preco, quantidade)

	if err != nil {
		log.Println("Erro ao inserir produto:", err)
	}
	defer db.Close()
}

func DeletarProdutoDb(id string) {
	db := db.ConectaComBancoDeDados()
	deletaProdutoInstrucao, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletaProdutoInstrucao.Exec(id)
	db.Close()
}

func BuscaProdutoDB(id string) Produto {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()
	return produto
}
