package routers

import (
	"encoding/json"
	"strconv"

	"github.com/nbedregal/gambit/bd"
	"github.com/nbedregal/gambit/models"
)

func InsertProduct(body string, user string) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el nombre (title) del Producto"
	}

	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar registrar producto " + t.ProdTitle + " > " + err2.Error()
	}
	return 200, "{ ProductId: " + strconv.Itoa(int(result)) + " }"
}

func UpdateProduct(body string, user string, id int) (int, string) {

	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el update del producto " + strconv.Itoa(id) + " > " + err2.Error()

	}

	return 200, "Update OK"

}

func DeleteProduct(body string, user string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar la id del producto a borrar"
	}

	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteProduct(id)

	if err != nil {
		return 400, "Ocurrio un error al intentar el delete " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete OK"
}
