package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/asynkron/protoactor-go/actor"
)

func LoadCoordinatorPIDS() ([]actor.PID, error) {
	jsonPids, err := os.ReadFile("coordinator_pids.json")
	if err != nil {
		log.Println("Could not read from coordinator_pids. Error: ", err)
		return nil, err
	}

	var pids []actor.PID
	if err := json.Unmarshal(jsonPids, &pids); err != nil {
		log.Println("Could not unmarshall json with pids. Error: ", err)
		return nil, err
	}

	return pids, nil
}

func HandleCoordinatorPID(coordinatorPID actor.PID) {
	pids, err := LoadCoordinatorPIDS()
	if err != nil {
		log.Println("Could not load coordinator pids. Error: ", err)
		return
	}
	pids = append(pids, coordinatorPID)

	jsonStr, err := json.MarshalIndent(pids, "", " ")
	if err != nil {
		log.Println("Could not marshall coordinator pids slice to json")
		return
	}

	if err := os.WriteFile("coordinator_pids.json", jsonStr, 0644); err != nil {
		log.Println("Could not write to coordinator pids json file. Error: ", err)
		return
	}
}
