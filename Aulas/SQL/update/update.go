package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "seuusuario:suasenha@/suabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Bianca", 2)
	stmt.Exec("Carlos", 1)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(3)

}
