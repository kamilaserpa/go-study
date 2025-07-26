package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectaComBancoDeDados() {
	stringDeConexao := "user=root password=root dbname=root host=localhost port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
