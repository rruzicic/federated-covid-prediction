package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/gossiper/messages"
)

type Gossiper struct{}

type (
	BroadcastCoordinatorPID struct{}
	GossipWeights           struct{}
)

func (state *Gossiper) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *BroadcastCoordinatorPID:
		log.Println("Broadcasting coordinators pid. PID: ", ctx.Parent())
		messages.BroadcastCoordinatorPIDMessage(ctx)
	}
}

func NewGossiper() actor.Actor {
	return &Gossiper{}
}
