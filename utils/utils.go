package utils

import (
	"github.com/aws/aws-lambda-go/events"

	"encoding/json"
	"errors"
)

// Unmarshall JSON from request to intrerface
func UnmarshallGatewayRequestBody(req events.APIGatewayProxyRequest, v interface{}) {
	
	// Unmarshall Request body to struct pointer
	json.Unmarshal([]byte(req.Body), &v)
}

// Build JSON response for gateway from status code and interface
func BuildJSONGatewayResponse(status_code int, v interface{}) (events.APIGatewayProxyResponse , error) {

	// Marshal JSON webhook response
	raw_msg, err := json.Marshal(v)

	if (err == nil){

		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"

		return events.APIGatewayProxyResponse { StatusCode : status_code, Headers: headers, Body : string(raw_msg)}, nil
	}

	return events.APIGatewayProxyResponse { StatusCode : 500, Body : "i don't feel so good"}, errors.New("Unable marshal interface")

}

// Build response to gateway with status code
func BuildGatewayResponse(status_code int) (events.APIGatewayProxyResponse , error) {

	return events.APIGatewayProxyResponse { StatusCode : status_code }, nil
}