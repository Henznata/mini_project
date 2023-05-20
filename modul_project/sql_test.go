package modul_project

import (
	"context"
	"fmt"
	"testing"
)

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	// db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name,pass,email string
		err = rows.Scan(&id, &name,&pass,&email)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ",id)
		fmt.Println("name: ",name)
    fmt.Println("password: ",pass)
    fmt.Println("email: ",email)
	}
}
