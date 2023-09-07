package bd

import (
	"database/sql"
	"fmt"
	"strconv"

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

func SelectUser(userId string) (models.User, error) {
	fmt.Println("Comienza SelectUser")
	user := models.User{}

	err := DbConnect()

	if err != nil {
		return user, err
	}

	defer Db.Close()

	sentencia := "SELECT * FROM users WHERE User_UUID = '" + userId + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)

	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	rows.Scan(&user.UserUUID, &user.UserEmail, &firstName, &lastName, &user.UserStatus, &user.UserDateAdd, &dateUpg)

	user.UserFirstName = firstName.String
	user.UserLastName = lastName.String
	user.UserDateUpg = dateUpg.Time.String()

	fmt.Println("SelectUser > Ejecución exitosa")

	return user, nil

}

func SelectUsers(page int) (models.ListUser, error) {
	fmt.Println("Comienza SelectUsers")

	var lu models.ListUser
	user := []models.User{}

	err := DbConnect()

	if err != nil {
		return lu, err
	}

	defer Db.Close()

	var offset int = (page * 10) - 10
	var sentencia string
	var sentenciaCount string = "SELECT count(*) as registros FROM users"

	sentencia = "SELECT * FROM users LIMIT 10"
	if offset > 0 {
		sentencia += " OFFSET " + strconv.Itoa(offset)
	}

	var rowsCount *sql.Rows

	rowsCount, err = Db.Query(sentenciaCount)

	if err != nil {
		return lu, err
	}

	defer rowsCount.Close()

	rowsCount.Next()

	var registros int
	rowsCount.Scan(&registros)

	lu.TotalItems = registros

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return lu, err
	}

	for rows.Next() {
		var u models.User
		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime

		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpg = dateUpg.Time.String()

		user = append(user, u)
	}
	fmt.Println("Select Users > Ejecución exitosa")

	lu.Data = user
	return lu, nil

}
