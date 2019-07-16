package main
 
import (
		"context"
		"strings"
		"os"
		"github.com/aws/aws-lambda-go/lambda"
		"github.com/aws/aws-lambda-go/events"

		utils "github.com/kaveenr/batagoda-lambda/utils"
		telegram "github.com/kaveenr/batagoda-lambda/telegram"
		batz "github.com/janithl/batagoda/batagodax"
)

// Lambda event to handle gateway requests
func HandleLambdaEvent(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Parse Telegram webhook request from gateway
	req := telegram.WebHookRequest {}
	utils.UnmarshallGatewayRequestBody(event, &req)
	
	call_word := os.Getenv("CALL_WORD")
	msg_text := req.Message.Text

	// If call word is present, respond
	if (strings.Contains(strings.ToLower(msg_text), call_word)) {

		// Build Response
		msg := telegram.WebHookResponse{Method : "sendMessage", ChatID : req.Message.Chat.ID}
		msg.Text = batz.Respond(strings.ReplaceAll(msg_text, call_word, ""))

		return utils.BuildJSONGatewayResponse(200, msg)
	}

	// Respond To Gateway
	return utils.BuildGatewayResponse(200)
}

func main() {

        lambda.Start(HandleLambdaEvent)
}