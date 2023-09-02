package main

import (
	"context"
	"os"

	//"strings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/nbedregal/gambit/awsgo"
	"github.com/nbedregal/gambit/bd"

	//	"github.com/nbedregal/gambit/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		panic("Error en los parámetros. Debe enviar 'SecretName', 'UserPoolId', 'Region', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	//prefix := os.Getenv("UrlPrefix")
	/*path := strings.Replace(request.RawPath, prefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers*/

	bd.ReadSecret()

	//TODO confirmación

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		//StatusCode: status,
		//Body: string(message),
		Headers: headersResp,
	}

	return res, nil
}

func ValidoParametros() bool {
	_, traeParametros := os.LookupEnv("SecretName")
	if !traeParametros {
		return traeParametros
	}
	_, traeParametros = os.LookupEnv("UserPoolId")
	if !traeParametros {
		return traeParametros
	}
	_, traeParametros = os.LookupEnv("Region")
	if !traeParametros {
		return traeParametros
	}
	_, traeParametros = os.LookupEnv("UrlPrefix")
	if !traeParametros {
		return traeParametros
	}

	return traeParametros
}
