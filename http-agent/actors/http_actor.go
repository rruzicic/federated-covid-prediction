package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

type HTTPActor struct{}

type (
	getRandomWeights struct{}
	model            struct{}
	initWeights      struct{ Weights []byte }
	collect          struct{ Collect []byte }
	allPeersDone     struct{}
	oneEpoch         struct{}
	weigths          struct{}
	plotLoss         struct{}
)

func (state *HTTPActor) Recieve(context actor.Context) {
	switch msg := context.Message().(type) {
	case *getRandomWeights:
		{
			// retval from this to be later used to communicate with gossiper agent
			_, err := messages.GetRandomWeights()
			if err != nil {
				log.Println("Could not get random weights for message: ", msg)
				break
			}
			break
		}
	case *model:
		{
			_, err := messages.Model()
			if err != nil {
				log.Println("Could not get model for message: ", msg)
				break
			}
			break
		}
	case *initWeights:
		{
			_, err := messages.InitWeights(msg.Weights)
			if err != nil {
				log.Println("Could not init model weights for message: ", msg)
				break
			}
			break
		}
	case *collect:
		{
			statusCode, _, err := messages.Collect(msg.Collect)
			if err != nil {
				log.Println("Could not init model weights for message: ", msg)
				break
			}
			if statusCode == 200 {
				log.Println("Got status code 200")
				break
			}
			if statusCode == 201 {
				log.Println("Got status code 201")
				break
			}
			break
		}
	case *allPeersDone:
		{
			_, err := messages.AllPeersDone()
			if err != nil {
				log.Println("Could not send all peers done for message: ", msg)
				break
			}
			break
		}
	case *oneEpoch:
		{
			_, err := messages.OneEpoch()
			if err != nil {
				log.Println("Could not send one epoch for message: ", msg)
				break
			}
			break
		}
	case *weigths:
		{
			_, err := messages.Weights()
			if err != nil {
				log.Println("Could not get weights for message: ", msg)
				break
			}
			break
		}
	case *plotLoss:
		{
			_, err := messages.PlotLoss()
			if err != nil {
				log.Println("Could not send plot loss for message: ", msg)
				break
			}
			break
		}
	default:
		log.Panic("Got message that can't be parsed. Message: ", msg)
	}
}
