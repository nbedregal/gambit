package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"

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

func SelectProduct(request events.APIGatewayV2HTTPRequest) (int, string) {

	var t models.Product
	var page, pagesize int
	var orderType, orderField string

	param := request.QueryStringParameters

	page, _ = strconv.Atoi(param["page"])
	pagesize, _ = strconv.Atoi(param["pagesize"])
	orderType = param["orderType"]   // D = Desc. A o Nil = Asc.
	orderField = param["orderField"] // 'I' Id, 'T' Title, 'D' Description, 'F' Created At, 'P' Price, 'C' CategId, 'S' Stock

	if !strings.Contains("ITDFPCS", orderField) {
		orderField = ""
	}

	var choice string
	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}
	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch = param["search"]
	}
	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategoryId, _ = strconv.Atoi(param["categId"])
	}
	if len(param["slug"]) > 0 {
		choice = "U"
		t.ProdPath = param["slug"]
	}
	if len(param["slugCateg"]) > 0 {
		choice = "K"
		t.ProdCategPath = param["slugCateg"]
	}

	fmt.Println(param)

	result, err2 := bd.SelectProduct(t, choice, page, pagesize, orderType, orderField)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar los resultados de la búsqueda de tipo " + choice + " en productos > " + err2.Error()
	}

	product, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON la busqueda de productos"
	}

	return 200, string(product)

}

func UpdateStock(body string, user string, id int) (int, string) {

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
	err2 := bd.UpdateStock(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el update del stock de producto " + strconv.Itoa(id) + " > " + err2.Error()

	}

	return 200, "Update OK"

}
