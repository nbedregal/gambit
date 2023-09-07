package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nbedregal/gambit/bd"
	"github.com/nbedregal/gambit/models"
)

func UpdateUser(body string, user string) (int, string) {

	var u models.User

	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(u.UserFirstName) == 0 && len(u.UserLastName) == 0 {
		return 400, "Debe especificar el nombre o apellido del usuario"
	}

	_, encontrado := bd.UserExists(user)

	if !encontrado {
		return 400, "No existe un usuario con ese userUUID"
	}

	err2 := bd.UpdateUser(u, user)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el update de usuario " + user + " > " + err2.Error()

	}

	return 200, "Update OK"

}

func SelectUser(body string, user string) (int, string) {
	_, encontrado := bd.UserExists(user)

	if !encontrado {
		return 400, "No existe un usuario con ese userUUID"
	}

	row, err := bd.SelectUser(user)
	fmt.Println(row)
	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el select del usuario " + user + " > " + err.Error()
	}

	respJson, err2 := json.Marshal(row)
	if err2 != nil {
		return 500, "Error al formatear los datos del usuario como JSON"
	}

	return 200, string(respJson)
}

func SelectUsers(body string, user string, request events.APIGatewayV2HTTPRequest) (int, string) {

	var page int

	if len(request.QueryStringParameters["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	userData, err := bd.SelectUsers(page)

	if err != nil {
		return 400, "Ocurrio un error al intentar obtener la lista de usuarios " + err.Error()
	}

	respJson, err := json.Marshal(userData)
	if err != nil {
		return 500, "Error al formatear los datos de usuarios a JSON"
	}

	return 200, string(respJson)
}
