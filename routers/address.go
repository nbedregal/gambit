package routers

import (
	"encoding/json"

	"github.com/nbedregal/gambit/bd"
	"github.com/nbedregal/gambit/models"
)

func InsertAddress(body string, user string) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}
	if t.AddAddress == "" {
		return 400, "Debe especificar el address"
	}

	if t.AddName == "" {
		return 400, "Debe especificar el name"
	}

	if t.AddTitle == "" {
		return 400, "Debe especificar el title"
	}

	if t.AddCity == "" {
		return 400, "Debe especificar el city"
	}

	if t.AddPhone == "" {
		return 400, "Debe especificar el phone"
	}

	if t.AddPostalCode == "" {
		return 400, "Debe especificar el postal code"
	}

	err = bd.InsertAddress(t, user)
	if err != nil {
		return 400, "Ocurrio un error a intentar un registro de address del usuario " + user + " > " + err.Error()
	}

	return 200, "InsertAddress OK"

}

func UpdateAddress(body string, user string, id int) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}
	t.AddId = id

	var encontrado bool
	err, encontrado = bd.AddExists(user, t.AddId)
	if err != nil {
		return 400, "Error al ejecutar consulta de existencia " + err.Error()
	}
	if !encontrado {
		return 400, "No se encontró el usuario asociado a Address" + user
	}

	err = bd.UpdateAddress(t)
	if err != nil {
		return 400, "Ocurrio un error al intentar actualizar el address para usuario " + user + " > " + err.Error()
	}

	return 200, "UpdateAddress OK"

}

func DeleteAddress(user string, id int) (int, string) {

	err, encontrado := bd.AddExists(user, id)
	if err != nil {
		return 400, "Error al ejecutar consulta de existencia " + err.Error()
	}
	if !encontrado {
		return 400, "No se encontró el usuario asociado a Address" + user
	}

	err = bd.DeleteAddress(id)

	if err != nil {
		return 400, "Error al intentar eliminar una dirección"
	}

	return 200, "Delete OK"
}
