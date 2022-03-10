package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "seuusuario:suasenha@/suabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Mauricio")

	_, err = stmt.Exec(1, "Tiago") // chave duplicada
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

	//ele não vai adicionar ninguém na tabela, ele trata como uma transação

}
