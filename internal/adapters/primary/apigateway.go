package primary

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/catvotes/internal/domains/votes"
)

type APIGatewayPrimaryAdapter struct {
	service votes.PrimaryPort
}

func NewAPIGatewayPrimaryAdapter(s votes.PrimaryPort) *APIGatewayPrimaryAdapter {
	return &APIGatewayPrimaryAdapter{
		s,
	}
}

// PlaceOrder receives the request, processes it and returns a Response or an error
func (a *APIGatewayPrimaryAdapter) HandleVote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Verifying the body of the request
	v := votes.Vote{}
	err := json.Unmarshal([]byte(request.Body), &v)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Processing vote
	err = a.service.Vote(v.ImageID, v.Vote)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 502}, err
	}

	return successfulResponse(), nil
}

func successfulResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      201,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}
}
