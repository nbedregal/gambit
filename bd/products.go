package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nbedregal/gambit/models"
	"github.com/nbedregal/gambit/tools"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comienza registro de InsertProduct")

	err := DbConnect()

	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO products (Prod_Title"
	values := "'" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		sentencia += ", Prod_Description"
		values += ",'" + tools.EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		sentencia += ", Prod_Price"
		values += "," + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	if p.ProdCategoryId > 0 {
		sentencia += ", Prod_CategoryId"
		values += "," + strconv.Itoa(p.ProdCategoryId)
	}

	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
		values += "," + strconv.Itoa(p.ProdStock)
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", Prod_Path"
		values += ",'" + tools.EscapeString(p.ProdPath) + "'"
	}

	sentencia += ") VALUES (" + values + ")"

	var result sql.Result

	result, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Products > Ejecución exitosa")

	return LastInsertId, nil

}
