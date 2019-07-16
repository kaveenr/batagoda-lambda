package gateway

import (
	"encoding/json"
	"errors"
)

// Unmarshall JSON from request to intrerface
func UnmarshallGatewayRequestBody(req GatewayRequest, v interface{}) (int) {
	
	// Unmarshall Request body to struct pointer
	json.Unmarshal([]byte(req.Body), &v)

	return req.StatusCode
}

// Build JSON response for gateway from status code and interface
func BuildJSONGatewayResponse(status_code int, v interface{}) (GatewayResponse, error) {

	// Marshal JSON webhook response
	raw_msg, err := json.Marshal(v)

	if (err == nil){

		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"

		return GatewayResponse{ StatusCode : status_code, Headers: headers, Body : string(raw_msg)}, nil
	}

	return GatewayResponse{ StatusCode : 500, Body : "i don't feel so good"}, errors.New("Unable marshal interface")

}

// Build response to gateway with status code
func BuildGatewayResponse(status_code int) (GatewayResponse, error) {

	return GatewayResponse{ StatusCode : status_code }, nil
}