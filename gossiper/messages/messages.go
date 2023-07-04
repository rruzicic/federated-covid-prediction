package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/peer/services"

	grpctransformations "github.com/rruzicic/federated-covid-prediction/gossiper/grpc_transformations"
	http_messages "github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

func BroadcastCoordinatorPIDMessage(ctx actor.Context) error {
	// get all addresses where you need to broadcast
	peerAddresses, err := services.GetPeerAddresses()
	if err != nil {
		log.Panic("Could not get peer addresses. Error: ", err.Error())
		return err
	}

	// make request body
	reqBody, err := json.Marshal(ctx.Parent())
	if err != nil {
		log.Panic("Could not marshall coordinator pid. Error: ", err.Error())
		return err
	}

	// make and send post requests to each peer
	for _, address := range peerAddresses {
		url := fmt.Sprintf("%s:%d/coordinator-pid", address.Address, address.Port)

		_, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
		if err != nil {
			log.Println("Could not send request to peer at ", address)
			return err
		}
	}

	return nil
}

func GossipWeightsMessage(weights http_messages.WeightsResponse, ctx actor.Context) error {
	coordinators, err := services.LoadCoordinatorPIDS()
	if err != nil {
		log.Println("Could not get coordinator pids. Error: ", err.Error())
		return err
	}

	grpcWeights := grpctransformations.MessageWeightsToGRPCWeights(weights)

	for _, coordinatorPid := range coordinators {
		ctx.Send(&coordinatorPid, grpcWeights)
	}

	return nil
}
