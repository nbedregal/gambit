package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nbedregal/gambit/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza el registro de InsertOrder")

	err := DbConnect()

	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	sentencia += o.OrderUserUUID + "'," + strconv.FormatFloat(o.OrderTotal, 'e', -1, 64) + "," + strconv.Itoa(o.OrderAddId) + ")"

	var result sql.Result

	result, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	lastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	for _, od := range o.OrderDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(lastInsertId))
		sentencia += "," + strconv.Itoa(od.ODProdId) + "," + strconv.Itoa(od.ODQuantity) + "," + strconv.FormatFloat(od.ODPrice, 'e', -1, 64) + ")"

		fmt.Println(sentencia)

		_, err = Db.Exec(sentencia)

		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}

	}

	fmt.Println("Insert order > Ejecución exitosa")

	return lastInsertId, nil

}

func SelectOrders(user string, fechaDesde string, fechaHasta string, page int, orderId int) ([]models.Orders, error) {
	fmt.Println("Comienza el registro de SelectOrders")
	var orders []models.Orders

	sentencia := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders "

	if orderId > 0 {
		sentencia += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			offset = (10 * (page - 1))
		}

		if len(fechaHasta) == 10 {
			fechaHasta += " 23:59:59"
		}

		var where string
		var whereUser string = " Order_UserUUID = '" + user + "'"

		if len(fechaDesde) > 0 && len(fechaHasta) > 0 {
			where += " WHERE Order_Date BETWEEN '" + fechaDesde + "' AND '" + fechaHasta + "'"
		}
		if len(where) > 0 {
			where += " AND " + whereUser
		} else {
			where += " WHERE " + whereUser
		}

		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}

		sentencia += where + limit

		fmt.Println(sentencia)
	}

	err := DbConnect()

	if err != nil {
		return orders, err
	}

	defer Db.Close()

	var rows *sql.Rows

	rows, err = Db.Query(sentencia)
	if err != nil {
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var order models.Orders

		var orderAddId sql.NullInt32

		err := rows.Scan(&order.OrderId, &order.OrderUserUUID, &orderAddId, &order.OrderDate, &order.OrderTotal)
		if err != nil {
			return orders, err
		}
		order.OrderAddId = int(orderAddId.Int32)

		var rowsD *sql.Rows
		sentenciaD := "SELECT OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderID = " + strconv.Itoa(order.OrderId)

		rowsD, err = Db.Query(sentenciaD)

		if err != nil {
			return orders, err
		}

		for rowsD.Next() {
			var odId int64
			var odProdId int64
			var odQuantity int64
			var odPrice float64

			err = rowsD.Scan(&odId, &odProdId, &odQuantity, &odPrice)

			if err != nil {
				return orders, err
			}

			var od models.OrderDetails
			od.ODId = int(odId)
			od.ODProdId = int(odProdId)
			od.ODQuantity = int(odQuantity)
			od.ODPrice = odPrice

			order.OrderDetails = append(order.OrderDetails, od)
		}

		orders = append(orders, order)

		rowsD.Close()
	}

	fmt.Println("SelectOrders > Ejecucion exitosa")

	return orders, nil
}
