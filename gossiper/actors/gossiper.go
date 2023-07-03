package actors

import "github.com/asynkron/protoactor-go/actor"

type Gossiper struct{}

type (
	BroadcastCoordinatorPID struct{}
	GossipWeights           struct{}
)

func (state *Gossiper) Receive(ctx actor.Context) {

}

func NewGossiper() actor.Actor {
	return &Gossiper{}
}
