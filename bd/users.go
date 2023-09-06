package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nbedregal/gambit/models"
	"github.com/nbedregal/gambit/tools"
)

func UpdateUser(u models.User, user string) error {

	fmt.Println("Comienza registro de UpdateUser")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE users SET "

	sentencia = tools.ArmoSentencia(sentencia, "User_FirstName", "S", 0, 0, u.UserFirstName)
	sentencia = tools.ArmoSentencia(sentencia, "User_LastName", "S", 0, 0, u.UserLastName)

	sentencia += ", User_DateUpg = '" + tools.FechaMySQL() + "' WHERE User_UUID = '" + user + "'"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecución exitosa")

	return nil

}
