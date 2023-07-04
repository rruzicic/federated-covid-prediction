package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	gossip_actors "github.com/rruzicic/federated-covid-prediction/gossiper/actors"
	grpctransformations "github.com/rruzicic/federated-covid-prediction/gossiper/grpc_transformations"
	grpc_messages "github.com/rruzicic/federated-covid-prediction/grpc"
	"github.com/rruzicic/federated-covid-prediction/peer/services"
)

type Coordinator struct {
	behavior actor.Behavior
}

type Message struct{} // default message

type ( // messages that are sent from other gossipers that end up in the coordinator
	BecomeLeader     struct{}
	GossipedWeights  struct{}
	CollectedWeights struct{}
	AllPeersDone     struct{}
	PeerExited       struct {
		Address services.AddressAndHost // GetYourAddress from peer/services/address_service.go
		PID     actor.PID               // since gossiper will be sending this. this will be available from ctx.Parent()
	}
)

func (state *Coordinator) Receive(ctx actor.Context) {
	state.behavior.Receive(ctx)
}

func (state *Coordinator) Startup(ctx actor.Context) {
	// explicitly set state from go program itself
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator in state Startup. Received &Message")

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)

		ctx.Send(gossiperPid, &gossip_actors.BroadcastCoordinatorPID{})

		state.behavior.Become(state.Init)

	case *BecomeLeader:
		log.Println("Coordinator in state Startup. Received &Message")

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)

		ctx.Send(gossiperPid, &gossip_actors.BroadcastCoordinatorPID{})

		state.behavior.Become(state.InitLeader)

	case *PeerExited:
		log.Println("Coordinator in state Startup. Received &PeerExit")

		services.RemovePeerAddress(msg.Address)
		services.RemoveCoordinatorPID(msg.PID)
	}
}

func (state *Coordinator) InitLeader(ctx actor.Context) {
	// explicitly set state from go program itself
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state InitLeader. Received &Message")

		// get random weights using http actor. future request them back here

		// init weights using http actor. possibly future request 200 back here
		// can use props and pid from http actor from above

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)
		ctx.Send(gossiperPid, &gossip_actors.GossipWeights{}) // fill with res from http agent

		state.behavior.Become(state.OneEpoch)
		ctx.Send(ctx.Self(), &Message{})

	case *PeerExited:
		log.Println("Coordinator is in state InitLeader. Received &PeerExit")

		services.RemovePeerAddress(msg.Address)
		services.RemoveCoordinatorPID(msg.PID)
	}
}

func (state *Coordinator) Init(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *grpc_messages.GRPCWeights:
		log.Println("Coordinator is in state Init")
		messageWeights := grpctransformations.GRPCWeightsToMessageWeights(msg)

		// send weights to http agent wait for 200 before changing state

		state.behavior.Become(state.OneEpoch)
		ctx.Send(ctx.Self(), &Message{})

	case *PeerExited:
		log.Println("Coordinator is in state Init. Received &PeerExit")

		services.RemovePeerAddress(msg.Address)
		services.RemoveCoordinatorPID(msg.PID)
	}
}

func (state *Coordinator) OneEpoch(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state OneEpoch. Received &Message")

		// spawn http actor send one epoch signal, request future result for 200 or 201

		// if 200
		state.behavior.Become(state.Collect)
		ctx.Send(ctx.Self(), &Message{})

		// if 201
		state.behavior.Become(state.Exit)
		ctx.Send(ctx.Self(), &Message{})

	case *PeerExited:
		log.Println("Coordinator is in state OneEpoch. Received &PeerExit")

		services.RemovePeerAddress(msg.Address)
		services.RemoveCoordinatorPID(msg.PID)
	}
}

func (state *Coordinator) Collect(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state Collect. Received &Message")

	case *AllPeersDone:
		log.Println("Coordinator is in state Collect. Received &AllPeersDone")

		state.behavior.Become(state.OneEpoch)
		ctx.Send(ctx.Self(), &Message{})

	case *PeerExited:
		log.Println("Coordinator is in state Collect. Received &PeerExit")

		services.RemovePeerAddress(msg.Address)
		services.RemoveCoordinatorPID(msg.PID)
	}
}

func (state *Coordinator) Exit(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state Exit. Received &Message")
	}
}

func NewCoordinator() actor.Actor {
	coord := &Coordinator{
		behavior: actor.NewBehavior(),
	}
	coord.behavior.Become(coord.Startup)
	return coord
}

func setupRemote() (remote.Remote, actor.RootContext) {
	yourAddress, _ := services.GetYourAddress()

	sys := actor.NewActorSystem()
	rmtCfg := remote.Configure(yourAddress.Address, yourAddress.Port)
	remoter := remote.NewRemote(sys, rmtCfg)

	return *remoter, *sys.Root
}

func SetupCoordinator() {
	// starts coordinator, agent system and supervision

	remoter, ctx := setupRemote()
	remoter.Start()

	supervision := actor.NewOneForOneStrategy(10, 1000, actor.DefaultDecider) // possibly implmenet your decider like in coordinator_mock
	props := actor.PropsFromProducer(NewCoordinator, actor.WithSupervisor(supervision))
	pid := ctx.Spawn(props)
	ctx.Send(pid, &Message{})
}

func SetupLeaderCoordinator() {
	// starts coordinator, agent system and supervision except after sending the first message explicity changes state to InitLeader

	remoter, ctx := setupRemote()
	remoter.Start()

	supervision := actor.NewOneForOneStrategy(10, 1000, actor.DefaultDecider) // possibly implmenet your decider like in coordinator_mock
	props := actor.PropsFromProducer(NewCoordinator, actor.WithSupervisor(supervision))
	pid := ctx.Spawn(props)
	ctx.Send(pid, &BecomeLeader{})
}
