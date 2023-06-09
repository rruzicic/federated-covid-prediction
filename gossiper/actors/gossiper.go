package actors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	grpctransformations "github.com/rruzicic/federated-covid-prediction/gossiper/grpc_transformations"
	grpc_message "github.com/rruzicic/federated-covid-prediction/grpc"
	http_messages "github.com/rruzicic/federated-covid-prediction/http-agent/messages"
	"github.com/rruzicic/federated-covid-prediction/peer/services"
)

type Gossiper struct{}

type (
	BroadcastCoordinatorPID struct{}
	GossipWeights           struct {
		Weights http_messages.WeightsResponse
	}
	Exit struct {
		CoordinatorPID     actor.PID
		YourAddressAndHost services.AddressAndHost
	}
	Collect struct {
		Weights http_messages.WeightsResponse
		Peers   int
	}
)

func (state *Gossiper) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *BroadcastCoordinatorPID:
		log.Println("Broadcasting coordinators pid. PID: ", ctx.Parent())

		// get all addresses where you need to broadcast
		peerAddresses, err := services.GetPeerAddresses()
		if err != nil {
			log.Panic("Could not get peer addresses. Error: ", err.Error())
			break
		}

		// make request body
		reqBody, err := json.Marshal(ctx.Parent())
		if err != nil {
			log.Panic("Could not marshall coordinator pid. Error: ", err.Error())
			break
		}

		// make and send post requests to each peer
		for _, address := range peerAddresses {
			url := fmt.Sprintf("http://%s:%d/coordinator-pid", address.Address, address.Port+1000)

			client := &http.Client{}
			req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
			if err != nil {
				log.Println("Could not make broadcast pid request for", ctx.Parent())
				log.Println(err.Error())
				break
			}
			_, err = client.Do(req)
			if err != nil {
				log.Println("Could not send request to peer at ", address)
				log.Println(err.Error())
				break
			}
		}

	case *GossipWeights:
		log.Println("Gossiping weights received from py server.")

		coordinators, err := services.LoadCoordinatorPIDS()
		if err != nil {
			log.Panic("Could not get coordinator pids. Error: ", err.Error())
			break
		}

		grpcWeights := grpctransformations.MessageWeightsToGRPCWeights(msg.Weights)

		for _, coordinatorPid := range coordinators {
			ctx.Send(&coordinatorPid, grpcWeights)
		}

	case *Collect:
		// make grpc message from this one
		grpcCollect := grpc_message.GRPCCollect{
			Weights: grpctransformations.MessageWeightsToGRPCWeights(msg.Weights),
			Peers:   int32(msg.Peers),
		}

		// send to all peers' coordinators
		coordinators, err := services.LoadCoordinatorPIDS()
		if err != nil {
			log.Panic("Could not get coordinator pids. Error: ", err.Error())
			break
		}

		for _, coordinatorPid := range coordinators {
			ctx.Send(&coordinatorPid, &grpcCollect)
		}

	case *Exit:
		log.Println("Gossiping that this peer is exiting the network.")
		grpcExit := grpc_message.GRPCExit{
			CoordinatorPID: &msg.CoordinatorPID,
			Address:        msg.YourAddressAndHost.Address,
			Port:           int32(msg.YourAddressAndHost.Port),
		}

		coordinators, err := services.LoadCoordinatorPIDS()
		if err != nil {
			log.Panic("Could not get coordinator pids. Error: ", err.Error())
			break
		}

		for _, coordinatorPid := range coordinators {
			ctx.Send(&coordinatorPid, &grpcExit)
		}
	}
}

func NewGossiper() actor.Actor {
	return &Gossiper{}
}
