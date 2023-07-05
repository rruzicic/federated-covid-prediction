package actors

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	gossip_actors "github.com/rruzicic/federated-covid-prediction/gossiper/actors"
	grpctransformations "github.com/rruzicic/federated-covid-prediction/gossiper/grpc_transformations"
	grpc_messages "github.com/rruzicic/federated-covid-prediction/grpc"
	http_actors "github.com/rruzicic/federated-covid-prediction/http-agent/actors"
	http_messages "github.com/rruzicic/federated-covid-prediction/http-agent/messages"
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
	ExitWait         struct{}
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
		ctx.Send(ctx.Self(), &Message{})

	case *BecomeLeader:
		log.Println("Coordinator in state Startup. Received &BecomeLeader")

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)

		ctx.Send(gossiperPid, &gossip_actors.BroadcastCoordinatorPID{})

		state.behavior.Become(state.InitLeader)
		ctx.Send(ctx.Self(), &Message{})

	case *grpc_messages.GRPCExit:
		log.Println("Coordinator in state Startup. Received &PeerExit")

		services.RemovePeerAddress(services.AddressAndHost{
			Address: msg.Address,
			Port:    int(msg.Port),
		})
		services.RemoveCoordinatorPID(*msg.CoordinatorPID)

		ctx.Send(ctx.Self(), &Message{})
	}
}

func (state *Coordinator) InitLeader(ctx actor.Context) {
	// explicitly set state from go program itself
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state InitLeader. Received &Message")

		httpProps := actor.PropsFromProducer(http_actors.NewHTTPActor)
		httpPid := ctx.Spawn(httpProps)
		messageWeights, err := ctx.RequestFuture(httpPid, &http_actors.GetRandomWeights{}, 30*time.Second).Result()
		if err != nil {
			log.Panic("Could not send and receive future from http actor for random weights. Error: ", err.Error())
			break
		}

		// init weights using http actor
		messageStatusCode, _ := ctx.RequestFuture(httpPid, &http_actors.InitWeights{
			Weights: messageWeights.(http_messages.WeightsResponse),
		}, 30*time.Second).Result()

		// gossip weights
		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)
		ctx.Send(gossiperPid, &gossip_actors.GossipWeights{
			Weights: messageWeights.(http_messages.WeightsResponse),
		})

		if messageStatusCode.(int) == 200 {
			log.Println("Coordinator got 200 from init weights")
			state.behavior.Become(state.OneEpoch)
			ctx.Send(ctx.Self(), &Message{})
			break
		}
		log.Panic("Didn't get status code 200 when initializing weights")

	case *grpc_messages.GRPCExit:
		log.Println("Coordinator is in state InitLeader. Received &PeerExit")

		services.RemovePeerAddress(services.AddressAndHost{
			Address: msg.Address,
			Port:    int(msg.Port),
		})
		services.RemoveCoordinatorPID(*msg.CoordinatorPID)

		ctx.Send(ctx.Self(), &Message{})
	}
}

func (state *Coordinator) Init(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *grpc_messages.GRPCWeights:
		log.Println("Coordinator is in state Init")
		messageWeights := grpctransformations.GRPCWeightsToMessageWeights(msg)

		// send weights to http agent wait for 200 before changing state
		httpProps := actor.PropsFromProducer(http_actors.NewHTTPActor)
		httpPid := ctx.Spawn(httpProps)
		messageStatusCode, _ := ctx.RequestFuture(httpPid, &http_actors.InitWeights{
			Weights: messageWeights,
		}, 30*time.Second).Result()

		if messageStatusCode.(int) == 200 {
			log.Println("Coordinator got 200 from init weights")
			state.behavior.Become(state.OneEpoch)
			ctx.Send(ctx.Self(), &Message{})
			break
		}
		log.Panic("Didn't get status code 200 when initializing weights")

	case *grpc_messages.GRPCExit:
		log.Println("Coordinator is in state Init. Received &PeerExit")

		services.RemovePeerAddress(services.AddressAndHost{
			Address: msg.Address,
			Port:    int(msg.Port),
		})
		services.RemoveCoordinatorPID(*msg.CoordinatorPID)

		ctx.Send(ctx.Self(), &Message{})
	}
}

func (state *Coordinator) OneEpoch(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state OneEpoch. Received &Message")

		// spawn http actor send one epoch signal, request future result for 200 or 201
		httpProps := actor.PropsFromProducer(http_actors.NewHTTPActor)
		httpPid := ctx.Spawn(httpProps)
		messageStatusCode, _ := ctx.RequestFuture(httpPid, &http_actors.OneEpoch{}, 30*time.Second).Result()

		if messageStatusCode.(int) == 200 {
			log.Println("Got 200 from one-epoch")
			state.behavior.Become(state.Collect)
			ctx.Send(ctx.Self(), &Message{})
			break
		}

		if messageStatusCode.(int) == 201 {
			log.Println("Got 201 from one-epoch")
			state.behavior.Become(state.Exit)
			ctx.Send(ctx.Self(), &Message{})
			break
		}
		log.Panic("Didn't get 200 or 201 from one-epoch singal")

	case *grpc_messages.GRPCExit:
		log.Println("Coordinator is in state OneEpoch. Received &PeerExit")

		services.RemovePeerAddress(services.AddressAndHost{
			Address: msg.Address,
			Port:    int(msg.Port),
		})
		services.RemoveCoordinatorPID(*msg.CoordinatorPID)

		ctx.Send(ctx.Self(), &Message{})
	}
}

func (state *Coordinator) Collect(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state Collect. Received &Message")
		// get weights from http actor
		httpProps := actor.PropsFromProducer(http_actors.NewHTTPActor)
		httpPid := ctx.Spawn(httpProps)
		messageWeights, err := ctx.RequestFuture(httpPid, &http_actors.Weigths{}, 30*time.Second).Result()
		if err != nil {
			log.Println("Could not recieve future from http actor for get weights")
			break
		}

		// prepare message to for gossiper
		peers, _ := services.NumberOfPeers()
		messageCollect := gossip_actors.Collect{
			Weights: messageWeights.(http_messages.WeightsResponse),
			Peers:   peers,
		}

		// spawn gossiper with those weights and send the rpc request with rpc collect weights
		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)
		ctx.Send(gossiperPid, &messageCollect)

	case *grpc_messages.GRPCCollect:
		log.Println("Coordinator is in state Collect. Received &GRPCCollect")
		// unpack grpc message
		messageWeights := grpctransformations.GRPCWeightsToMessageWeights(msg.Weights)
		messageCollect := http_messages.CollectResponse{
			Hidden_weights: messageWeights.Hidden_weights,
			Output_weights: messageWeights.Output_weights,
			Peers:          int(msg.Peers),
		}

		// send it to http agent
		httpProps := actor.PropsFromProducer(http_actors.NewHTTPActor)
		httpPid := ctx.Spawn(httpProps)
		messageStatusCode, _ := ctx.RequestFuture(httpPid, &http_actors.Collect{
			Collect: messageCollect,
		}, 30*time.Second).Result()

		if messageStatusCode.(int) == 200 {
			log.Println("Got status code 200 from collect.")
			break
		}

		if messageStatusCode.(int) == 201 {
			log.Println("Got status code 201 from collect sending all peers done")

			state.behavior.Become(state.OneEpoch)
			ctx.Send(ctx.Self(), &Message{})
			break
		}

	case *grpc_messages.GRPCAllPeersDone:
		log.Println("Coordinator is in state Collect. Received &AllPeersDone")

		state.behavior.Become(state.OneEpoch)
		ctx.Send(ctx.Self(), &Message{})

	case *grpc_messages.GRPCExit:
		log.Println("Coordinator is in state Collect. Received &PeerExit")

		services.RemovePeerAddress(services.AddressAndHost{
			Address: msg.Address,
			Port:    int(msg.Port),
		})
		services.RemoveCoordinatorPID(*msg.CoordinatorPID)

		ctx.Send(ctx.Self(), &Message{})
	}
}

func (state *Coordinator) Exit(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *Message:
		log.Println("Coordinator is in state Exit. Received &Message")

		gossiperProps := actor.PropsFromProducer(gossip_actors.NewGossiper)
		gossiperPid := ctx.Spawn(gossiperProps)

		yourAddress, _ := services.GetYourAddress()
		ctx.Send(gossiperPid, &gossip_actors.Exit{
			CoordinatorPID:     *ctx.Self(),
			YourAddressAndHost: *yourAddress,
		})

		httpProps := actor.PropsFromProducer(http_actors.NewHTTPActor)
		httpPid := ctx.Spawn(httpProps)
		messageStatusCode, _ := ctx.RequestFuture(httpPid, &http_actors.Exit{}, 30*time.Second).Result()

		if messageStatusCode.(int) == 200 {
			log.Println("Coordinator got 200 from exit")
			ctx.Send(ctx.Self(), &ExitWait{})
			break
		}
		log.Panic("Didn't get status code 200 when exiting")

	case *ExitWait:
		log.Println("Coordinator is in state Exit. Received &ExitWait")
		ctx.Send(ctx.Self(), &ExitWait{})
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

	fmt.Println("Press [ENTER] when all peers get this message")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	supervision := actor.NewOneForOneStrategy(10, 1000, actor.DefaultDecider) // possibly implmenet your decider like in coordinator_mock
	props := actor.PropsFromProducer(NewCoordinator, actor.WithSupervisor(supervision))
	pid := ctx.Spawn(props)
	ctx.Send(pid, &Message{})
}

func SetupLeaderCoordinator() {
	// starts coordinator, agent system and supervision except after sending the first message explicity changes state to InitLeader

	remoter, ctx := setupRemote()
	remoter.Start()

	fmt.Println("Press [ENTER] when all peers get this message")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	supervision := actor.NewOneForOneStrategy(10, 1000, actor.DefaultDecider) // possibly implmenet your decider like in coordinator_mock
	props := actor.PropsFromProducer(NewCoordinator, actor.WithSupervisor(supervision))
	pid := ctx.Spawn(props)
	ctx.Send(pid, &BecomeLeader{})
}
