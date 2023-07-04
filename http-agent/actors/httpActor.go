package actors

import (
	"log"
	"strconv"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

type HTTPActor struct{}

type (
	GetRandomWeights struct{}
	Model            struct{}
	InitWeights      struct{ Weights messages.WeightsResponse }
	Collect          struct{ Collect messages.CollectResponse }
	AllPeersDone     struct{}
	OneEpoch         struct{}
	Weigths          struct{}
	Exit             struct{}
)

func (state *HTTPActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *GetRandomWeights:
		{
			// retval from this to be later used to communicate with gossiper agent
			retval, statusCode, err := messages.GetRandomWeights()
			if err != nil {
				log.Println("Could not get random weights for message: ", msg)
				break
			}
			context.Respond(retval)
			log.Println("Got status code: " + strconv.Itoa(statusCode) + " and returned random weights to coordinator")
			break
		}
	case *Model:
		{
			retval, err := messages.GetModel()
			if err != nil {
				log.Println("Could not get model for message: ", msg)
				break
			}
			context.Respond(retval)
			log.Println("Returned model response to coordinator")
			break
		}
	case *InitWeights:
		{
			statusCode, err := messages.InitWeights(msg.Weights)
			if err != nil {
				log.Println("Could not init model weights for message: ", msg)
				break
			}
			context.Respond(statusCode)
			log.Println("Got status code: " + strconv.Itoa(statusCode) + " and responded back to coordinator")
			break
		}
	case *Collect:
		{
			statusCode, err := messages.CollectWeights(msg.Collect)
			if err != nil {
				log.Println("Could not init model weights for message: ", msg)
				break
			}
			if statusCode == 200 || statusCode == 201 {
				context.Respond(statusCode)
				log.Println("Got status code: " + strconv.Itoa(statusCode) + " and responded back to coordinator")
				break
			} else {
				log.Println("Got unexpected status code: " + strconv.Itoa(statusCode))
				break
			}
		}
	case *AllPeersDone:
		{
			statusCode, err := messages.AllPeersDone()
			if err != nil {
				log.Println("Could not send all peers done for message: ", msg)
				break
			}
			context.Respond(statusCode)
			log.Println("Got status code: " + strconv.Itoa(statusCode) + " and responded back to coordinator")
			break
		}
	case *OneEpoch:
		{
			statusCode, err := messages.DoOneEpoch()
			if err != nil {
				log.Println("Could not send one epoch for message: ", msg)
				break
			}
			context.Respond(statusCode)
			log.Println("Got status code: " + strconv.Itoa(statusCode) + " and responded back to coordinator")
			break
		}
	case *Weigths:
		{
			retval, statusCode, err := messages.GetWeights()
			if err != nil {
				log.Println("Could not get weights for message: ", msg)
				break
			}
			context.Respond(retval)
			log.Println("Got status code: " + strconv.Itoa(statusCode) + " and sent the weights to coordinator")
			break
		}
	case *Exit:
		{
			statusCode, err := messages.Exit()
			if err != nil {
				log.Println("Could not send plot loss for message: ", msg)
				break
			}
			context.Respond(statusCode)
			log.Println("Got status code: " + strconv.Itoa(statusCode) + " and responded back to coordinator")
			break
		}
	}
}

func NewHTTPActor() actor.Actor {
	return &HTTPActor{}
}
