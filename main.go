package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juliofernandolepore/go-mongo-twitter/awsgo"
	"github.com/juliofernandolepore/go-mongo-twitter/sm"
	"golang.org/x/net/context"
)

func main() {
	lambda.Start(ExLambda)
}

func ExLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.IniciarAWS()

	if !ValidoParametro() { // no existen variables de entorno (false)
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. deben incluir secretname, bucket name, urlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	secretModel, err := sm.GetSecret(os.Getenv("secretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura del secret" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
}

func ValidoParametro() bool {
	_, traerParametro := os.LookupEnv("secretName")
	if !traerParametro { //return false (env non exist)
		return traerParametro
	}
	return traerParametro
}
