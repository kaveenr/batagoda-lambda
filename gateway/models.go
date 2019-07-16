package gateway

type GatewayRequest struct {
	StatusCode int `json:"statusCode"`
	Body string `json:"body"`
}

type GatewayResponse struct {
	StatusCode int `json:"statusCode"`
	Body string `json:"body"`
	Headers map[string]string `json:"headers"`
}