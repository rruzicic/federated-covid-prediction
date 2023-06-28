package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

type HTTPActor struct{}

type (
	GetRandomWeights struct{}
	Model            struct{}
	InitWeights      struct{ Weights []byte }
	Collect          struct{ Collect []byte }
	AllPeersDone     struct{}
	OneEpoch         struct{}
	Weigths          struct{}
	PlotLoss         struct{}
)

func (state *HTTPActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *GetRandomWeights:
		{
			// retval from this to be later used to communicate with gossiper agent
			_, err := messages.GetRandomWeights()
			if err != nil {
				log.Println("Could not get random weights for message: ", msg)
				break
			}
			break
		}
	case *Model:
		{
			_, err := messages.Model()
			if err != nil {
				log.Println("Could not get model for message: ", msg)
				break
			}
			break
		}
	case *InitWeights:
		{
			_, err := messages.InitWeights(msg.Weights)
			if err != nil {
				log.Println("Could not init model weights for message: ", msg)
				break
			}
			break
		}
	case *Collect:
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
	case *AllPeersDone:
		{
			_, err := messages.AllPeersDone()
			if err != nil {
				log.Println("Could not send all peers done for message: ", msg)
				break
			}
			break
		}
	case *OneEpoch:
		{
			_, err := messages.OneEpoch()
			if err != nil {
				log.Println("Could not send one epoch for message: ", msg)
				break
			}
			break
		}
	case *Weigths:
		{
			_, err := messages.Weights()
			if err != nil {
				log.Println("Could not get weights for message: ", msg)
				break
			}
			break
		}
	case *PlotLoss:
		{
			_, err := messages.PlotLoss()
			if err != nil {
				log.Println("Could not send plot loss for message: ", msg)
				break
			}
			break
		}
	}
}

func NewHTTPActor() actor.Actor {
	return &HTTPActor{}
}
