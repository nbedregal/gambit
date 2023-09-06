package routers

import (
	"encoding/json"

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
