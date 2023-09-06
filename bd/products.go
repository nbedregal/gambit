package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

func UpdateProduct(p models.Product) error {
	fmt.Println("Comienza registro de UpdateProduct")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE products SET "

	//tools
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Title", "S", 0, 0, p.ProdTitle)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Description", "S", 0, 0, p.ProdDescription)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Price", "F", 0, p.ProdPrice, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_CategoryId", "N", p.ProdCategoryId, 0, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Stock", "N", p.ProdStock, 0, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Path", "S", 0, 0, p.ProdPath)

	sentencia += " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Product > Ejecución exitosa")

	return nil
}

func DeleteProduct(id int) error {
	fmt.Println("Comienza registro de DeleteProduct")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "DELETE FROM products WHERE Prod_Id =" + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Product > Ejecución exitosa")

	return nil
}

func SelectProduct(p models.Product, choice string, page int, pagesize int, orderType string, orderField string) (models.ProductResp, error) {
	fmt.Println("Comieza SelectProduct")
	var resp models.ProductResp
	var prod []models.Product

	err := DbConnect()

	if err != nil {
		return resp, err
	}

	defer Db.Close()

	var sentencia string
	var sentenciaCount string
	var where, limit string

	sentencia = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products "
	sentenciaCount = "SELECT count(*) as registros FROM products "

	switch choice {
	case "P":
		where = " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)
	case "S":
		where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%' "

	case "C":
		where = " WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategoryId)
	case "U":
		where = " WHERE UCASE(Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%' "
	case "K":
		join := " JOIN category ON Prod_CategoryId = Categ_Id AND CategPath LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%' "
		sentencia += join
		sentenciaCount += join
	}
	sentenciaCount += where

	var rows *sql.Rows
	rows, err = Db.Query(sentenciaCount)

	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}

	rows.Next()
	var regi sql.NullInt32
	err = rows.Scan(&regi)

	registros := int(regi.Int32)

	if page > 0 {
		if registros > pagesize {
			limit = " LIMIT " + strconv.Itoa(pagesize)
			if page > 1 {
				offset := pagesize * (page - 1)
				limit += " OFFSET " + strconv.Itoa(offset)
			}
		} else {
			limit = ""
		}
	}

	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			orderBy = " ORDER BY Prod_Id"

		case "T":
			orderBy = " ORDER BY Prod_Title"

		case "D":
			orderBy = " ORDER BY Prod_Description"

		case "F":
			orderBy = " ORDER BY Prod_CreatedAt"

		case "P":
			orderBy = " ORDER BY Prod_Price"

		case "S":
			orderBy = " ORDER BY Prod_Stock"

		case "C":
			orderBy = " ORDER BY Prod_CategoryId"
		}
		if orderType == "D" {
			orderBy += " DESC"
		}
	}

	sentencia += where + orderBy + limit

	fmt.Println(sentencia)

	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var p models.Product
		var prodId sql.NullInt32
		var prodTitle sql.NullString
		var prodDescription sql.NullString
		var prodCreatedAt sql.NullTime
		var prodUpdated sql.NullTime
		var prodPrice sql.NullFloat64
		var prodpath sql.NullString
		var prodCategoryId sql.NullInt32
		var prodStock sql.NullInt32

		err := rows.Scan(&prodId, &prodTitle, &prodDescription, &prodCreatedAt, &prodUpdated, &prodPrice, &prodpath, &prodCategoryId, &prodStock)

		if err != nil {
			return resp, err
		}

		p.ProdId = int(prodId.Int32)
		p.ProdTitle = prodTitle.String
		p.ProdDescription = prodDescription.String
		p.ProdCreatedAt = prodCreatedAt.Time.String()
		p.ProdUpdated = prodUpdated.Time.String()
		p.ProdPrice = prodPrice.Float64
		p.ProdPath = prodpath.String
		p.ProdCategoryId = int(prodCategoryId.Int32)
		p.ProdStock = int(prodStock.Int32)

		prod = append(prod, p)
	}

	resp.TotalItems = registros
	resp.Data = prod

	fmt.Println("Select Product > Ejecución exitosa")

	return resp, nil
}
