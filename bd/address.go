package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nbedregal/gambit/models"
	"github.com/nbedregal/gambit/tools"
)

func InsertAddress(add models.Address, user string) error {
	fmt.Println("Comienza el registro de InsertAddress")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name )"
	sentencia += "VALUES ('" + user + "','" + add.AddAddress + "','" + add.AddCity + "','" + add.AddState + "','" + add.AddPostalCode + "','" + add.AddPhone + "','" + add.AddTitle + "','" + add.AddName + "')"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(sentencia)
	fmt.Println("InsertAddress > Ejecución exitosa")
	return nil
}

func AddExists(user string, id int) (error, bool) {
	fmt.Println("Comienza el registro de AddExists")

	err := DbConnect()

	if err != nil {
		return err, false
	}

	defer Db.Close()

	sentencia := "SELECT 1 FROM addresses WHERE Add_Id = " + strconv.Itoa(id) + " AND Add_UserId = '" + user + "'"

	rows, err := Db.Query(sentencia)
	if err != nil {
		return err, false
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("AddressExists > Ejecución exitosa > valor devuelto " + valor)

	if valor == "1" {
		return nil, true
	}
	return nil, false
}

func UpdateAddress(t models.Address) error {
	fmt.Println("Comienza el registro de UpdateAddress")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE addresses SET "

	sentencia = tools.ArmoSentencia(sentencia, "Add_Title", "S", 0, 0, t.AddTitle)
	sentencia = tools.ArmoSentencia(sentencia, "Add_Name", "S", 0, 0, t.AddName)
	sentencia = tools.ArmoSentencia(sentencia, "Add_Address", "S", 0, 0, t.AddAddress)
	sentencia = tools.ArmoSentencia(sentencia, "Add_City", "S", 0, 0, t.AddCity)
	sentencia = tools.ArmoSentencia(sentencia, "Add_State", "S", 0, 0, t.AddState)
	sentencia = tools.ArmoSentencia(sentencia, "Add_PostalCode", "S", 0, 0, t.AddPostalCode)
	sentencia = tools.ArmoSentencia(sentencia, "Add_Phone", "S", 0, 0, t.AddPhone)

	sentencia += " WHERE Add_Id = " + strconv.Itoa(t.AddId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(sentencia)
	fmt.Println("UpdateAddress > Ejecución exitosa")
	return nil

}

func DeleteAddress(id int) error {

	fmt.Println("Comienza el registro de UpdateAddress")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(sentencia)
	fmt.Println("DeleteAddress > Ejecución exitosa")
	return nil

}

func SelectAddress(user string) ([]models.Address, error) {

	fmt.Println("Comienza el registro de SelectAddress")

	addr := []models.Address{}
	err := DbConnect()

	if err != nil {
		return addr, err
	}

	defer Db.Close()

	sentencia := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses WHERE Add_UserId = '" + user + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return addr, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString

		err := rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)

		if err != nil {
			return addr, err
		}

		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String

		addr = append(addr, a)
	}
	fmt.Println("SelectAddress > Ejecución exitosa")

	return addr, nil

}
