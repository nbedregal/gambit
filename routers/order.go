package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nbedregal/gambit/bd"
	"github.com/nbedregal/gambit/models"
)

func InsertOrder(body string, user string) (int, string) {
	var o models.Orders

	err := json.Unmarshal([]byte(body), &o)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	o.OrderUserUUID = user

	ok, msg := validOrder(o)

	if !ok {
		return 400, msg
	}

	result, err2 := bd.InsertOrder(o)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la orden " + err2.Error()
	}

	return 200, "{ OrderID:" + strconv.Itoa(int(result)) + "}"

}

func validOrder(o models.Orders) (bool, string) {

	if o.OrderTotal == 0 {
		return false, "Debe indicar el total de la orden"
	}

	count := 0

	for _, od := range o.OrderDetails {
		if od.ODProdId == 0 {
			return false, "Debe indicar el id de producto en el detalle de la orden"
		}
		if od.ODQuantity == 0 {
			return false, "Debe indicar la cantidad del producto en el detalle de la orden"
		}
		count++
	}

	if count == 0 {
		return false, "Debe indicar items en la orden"
	}

	return true, ""
}

func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {

	var fechaDesde, fechaHasta string
	var orderId int
	var page int

	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	if len(request.QueryStringParameters["fechaDesde"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechaDesde"]
	}
	if len(request.QueryStringParameters["fechaHasta"]) > 0 {
		fechaHasta = request.QueryStringParameters["fechaHasta"]
	}
	if len(request.QueryStringParameters["page"]) > 0 {
		page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}
	if len(request.QueryStringParameters["orderId"]) > 0 {
		orderId, _ = strconv.Atoi(request.QueryStringParameters["orderId"])
	}

	result, err2 := bd.SelectOrders(user, fechaDesde, fechaHasta, page, orderId)
	if err2 != nil {
		return 400, "Error la intentar recuperar las ordenes del " + fechaDesde + " al " + fechaHasta + " > " + err2.Error()
	}

	orders, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Ocurrio un error al convertir a Json las ordenes"
	}

	return 200, string(orders)

}
