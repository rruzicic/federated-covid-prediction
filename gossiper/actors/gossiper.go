package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/gossiper/messages"
	http_messages "github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

type Gossiper struct{}

type (
	BroadcastCoordinatorPID struct{}
	GossipWeights           struct {
		Weights http_messages.WeightsResponse
	}
)

func (state *Gossiper) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *BroadcastCoordinatorPID:
		log.Println("Broadcasting coordinators pid. PID: ", ctx.Parent())
		messages.BroadcastCoordinatorPIDMessage(ctx)

	case *GossipWeights:
		log.Println("Gossiping weights received from py server.")
		messages.GossipWeightsMessage(msg.Weights, ctx)
	}
}

func NewGossiper() actor.Actor {
	return &Gossiper{}
}
