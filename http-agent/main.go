package main

import (
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/rruzicic/federated-covid-prediction/http-agent/actors"
)

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(actors.NewHTTPActor)
	childPid := system.Root.Spawn(props)
	system.Root.Send(childPid, &actors.Model{})

	// to never exit
	_, _ = console.ReadLine()
}
