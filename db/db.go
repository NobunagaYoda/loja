package db

import "database/sql"

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=Tony2016 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
