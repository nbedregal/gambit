package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nbedregal/gambit/models"
	"github.com/nbedregal/gambit/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

/**
* Función para leer el secreto desde Secret Manager,
* la variable de sistema SecretName está definida en
* la configuración de la función lambda.
 */
func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

/**
* Función para conectarse con la base de datos definida
* en RDS.
 */
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la base de datos")
	return nil
}

/**
* Función para formar el dataSourceName a partir
* de los parámetros provenientes de model.
 */
func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = claves.Dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza UserIsAdmin")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserIsAdmin > Ejecución exitosa - valor devuelto " + valor)
	if valor == "1" {
		return true, ""
	}

	return false, "User is not admin"
}
