package bd

import (
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
