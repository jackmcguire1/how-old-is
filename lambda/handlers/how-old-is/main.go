package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackmcguire1/how-old-is/internal/dom/age"
	"github.com/jackmcguire1/how-old-is/internal/pkg/alexa"
	"github.com/jackmcguire1/how-old-is/internal/utils"
)

func randomResponse() string {
	responses := []string{
		"did not compute, repeat did not compute, for futher assistance please ask for 'help'",
		"Alert! Our control point is being captured",
		"For The Horde!",
		"For The Alliance!",
		"Snake?! Snakeeeeee!",
	}
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(responses))
	if i == len(responses) {
		i = i - 1
	}

	return responses[i]
}

func DispatchIntents(req alexa.Request) (res alexa.Response) {
	switch req.Body.Intent.Name {
	case age.AgeIntent:
		name := req.Body.Intent.Slots["name"].Value
		log.Println("found name", name)

		res = alexa.NewResponse("Age", age.GetTotalAgeFromName(name), false)
	case alexa.HelpIntent:
		res = alexa.NewResponse(
			"Help for How Old Is X",
			"To see how old someone or something is, simply ask how old is, followed by the person's name!",
			false,
		)
	case alexa.CancelIntent, alexa.NoIntent, alexa.StopIntent:
		res = alexa.NewResponse(
			"Thankyou for using how old is x",
			"Next time you're back, your favourite someone will probably be nanoseconds older.",
			true,
		)
	default:
		res = alexa.NewResponse(
			"Random response",
			randomResponse(),
			true,
		)
	}

	return
}

func handler(ctx context.Context, req alexa.Request) (resp alexa.Response, err error) {
	log.Println(utils.ToJSON(req))

	switch req.Body.Type {
	case alexa.LaunchRequestType:
		resp = alexa.NewResponse("how old is",
			"welcome. please ask me how old your precious someone or something is!",
			false,
		)
	case alexa.IntentRequestType:
		resp = DispatchIntents(req)
	default:
	}
	log.Println(utils.ToJSON(resp))
	return
}

func main() {
	lambda.Start(handler)
}
