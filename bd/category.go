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

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro de InsertCategory")

	err := DbConnect()

	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

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

	fmt.Println("Insert Category > Ejecución exitosa")

	return LastInsertId, nil

}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza registro de UpdateCategory")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE category SET "

	if len(c.CategName) > 0 {
		sentencia += "Categ_Name='" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}
		sentencia += "Categ_Path='" + tools.EscapeString(c.CategPath) + "'"
	}

	sentencia += " WHERE Categ_Id=" + strconv.Itoa(c.CategID)

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución exitosa")

	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Comienza registro de DeleteCategory")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "DELETE FROM category WHERE Categ_Id =" + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecución exitosa")

	return nil
}

func SelectCategories(categId int, slug string) ([]models.Category, error) {
	fmt.Println("Comienza registro de SelectCategories")

	var Categ []models.Category
	err := DbConnect()

	if err != nil {
		return Categ, err
	}

	defer Db.Close()

	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "

	if categId > 0 {
		sentencia += "WHERE Categ_Id = " + strconv.Itoa(categId)
	} else {
		if len(slug) > 0 {
			sentencia += "WHERE Categ_Path LIKE '%" + slug + "%'"
		}
	}

	fmt.Println(sentencia)

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Ejecución exitosa")

	return Categ, nil

}
