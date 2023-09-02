package handlers

import (
	"fmt"
	//"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(path string, method string, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	//id := request.PathParameters["id"]
	//idn, _ := strconv.Atoi(id)

	return 400, "Method Invalid"
}
