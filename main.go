// Main Package
package main
 
import (
		"github.com/aws/aws-lambda-go/lambda"
		"strings"
		"os"

		gateway "github.com/kaveenr/batagoda-lambda/gateway"
		telegram "github.com/kaveenr/batagoda-lambda/telegram"

		batz "github.com/janithl/batagoda/batagodax"
)

// Lambda event to handle gateway requests
func HandleLambdaEvent(event gateway.GatewayRequest) (gateway.GatewayResponse, error) {

	// Parse Telegram webhook request from gateway
	req := telegram.WebHookRequest {}
	gateway.UnmarshallGatewayRequestBody(event, &req)
	
	call_word := os.Getenv("CALL_WORD")
	msg_text := req.Message.Text

	// If call word is present, respond
	if (strings.Contains(strings.ToLower(msg_text), call_word)) {

		// Build Response
		msg := telegram.WebHookResponse{Method : "sendMessage", ChatID : req.Message.Chat.ID}
		msg.Text = batz.Respond(strings.ReplaceAll(msg_text, call_word, ""))

		return gateway.BuildJSONGatewayResponse(200, msg)
	}

	// Respond To Gateway
	return gateway.BuildGatewayResponse(200)
}

func main() {

        lambda.Start(HandleLambdaEvent)
}