package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	gossip_actors "github.com/rruzicic/federated-covid-prediction/gossiper/actors"
)

type Coordinator struct {
	behavior actor.Behavior
}

type Message struct{} // default message

type ( // messages that are sent from other gossipers that end up in the coordinator
	GossipedWeights  struct{}
	CollectedWeights struct{}
	AllActorsDone    struct{}
	PeerExited       struct{}
)

func (state *Coordinator) Receive(ctx actor.Context) {
	state.behavior.Receive(ctx)
}

func (state *Coordinator) Startup(ctx actor.Context) {
	// explicitly set state from go program itself
	switch ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator in state Startup. Received &Message")

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)

		ctx.Send(gossiperPid, &gossip_actors.BroadcastCoordinatorPID{})

		state.behavior.Become(state.Init)

	case *PeerExited:
		log.Println("Coordinator in state Startup. Received &PeerExit")

		// handle exit by calling peer.ip_service and peer.pid_service to remove them
	}
}

func (state *Coordinator) InitLeader(ctx actor.Context) {
	// explicitly set state from go program itself
	switch ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state InitLeader. Received &Message")

		// get random weights using http actor. future request them back here

		// init weights using http actor. possibly future request 200 back here
		// can use props and pid from http actor from above

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)
		ctx.Send(gossiperPid, &gossip_actors.GossipWeights{})

		state.behavior.Become(state.OneEpoch)
		ctx.Send(ctx.Self(), &Message{})

	case *PeerExited:
		log.Println("Coordinator is in state InitLeader. Received &PeerExit")
	}
}

func (state *Coordinator) Init(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *GossipedWeights:
		log.Println("Coordinator is in state Init")

		// send weights to http agent wait for 200 before changing state

		state.behavior.Become(state.OneEpoch)
		ctx.Send(ctx.Self(), &Message{})

	case *PeerExited:
		log.Println("Coordinator is in state Init. Received &PeerExit")
	}
}

func (state *Coordinator) OneEpoch(ctx actor.Context) {
	// if 200
	state.behavior.Become(state.Collect)

	// if 201
	state.behavior.Become(state.Exit)
}

func (state *Coordinator) Collect(ctx actor.Context) {
	state.behavior.Become(state.OneEpoch)
}

func (state *Coordinator) Exit(ctx actor.Context) {

}

func NewCoordinator() actor.Actor {
	coord := &Coordinator{
		behavior: actor.NewBehavior(),
	}
	coord.behavior.Become(coord.Startup)
	return coord
}

func SetupCoordinator() {
	// starts coordinator, agent system and supervision
}

func SetupLeaderCoordinator() {
	// starts coordinator, agent system and supervision except after sending the first message explicity changes state to InitLeader
}
