package modul_project

import (
	"database/sql"
	"testing"


	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "admin:@tcp(localhost:3306)/mini_project")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
