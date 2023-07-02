package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/asynkron/protoactor-go/actor"
)

func LoadCoordinatorPIDS() ([]actor.PID, error) {
	if _, err := os.Stat("coordinator_pids.json"); err != nil {
		return []actor.PID{}, nil
	}

	jsonPids, err := os.ReadFile("coordinator_pids.json")
	if err != nil {
		log.Println("Could not read from coordinator_pids. Error: ", err.Error())
		return nil, err
	}

	var pids []actor.PID
	if err := json.Unmarshal(jsonPids, &pids); err != nil {
		log.Println("Could not unmarshall json with pids. Error: ", err.Error())
		return nil, err
	}

	return pids, nil
}

func HandleCoordinatorPID(coordinatorPID actor.PID) {
	pids, err := LoadCoordinatorPIDS()
	if err != nil {
		log.Println("Could not load coordinator pids. Error: ", err.Error())
		return
	}
	pids = append(pids, coordinatorPID)

	jsonStr, err := json.MarshalIndent(pids, "", " ")
	if err != nil {
		log.Println("Could not marshall coordinator pids slice to json. Error: ", err.Error())
		return
	}

	if err := os.WriteFile("coordinator_pids.json", jsonStr, 0644); err != nil {
		log.Println("Could not write to coordinator pids json file. Error: ", err.Error())
		return
	}
}

func RemoveCoordinatorPID(coordinatorPID actor.PID) {
	pids, err := LoadCoordinatorPIDS()
	if err != nil {
		log.Println("Could not load coordinator pids. Error: ", err.Error())
		return
	}

	var newPids []actor.PID
	for idx, pid := range pids {
		// very scuffed way of checking this but who cares
		if (pid.Address == coordinatorPID.Address) && (pid.Id == coordinatorPID.Id) {
			newPids = append(pids[:idx], pids[idx+1:]...)
		}
	}

	jsonStr, err := json.MarshalIndent(newPids, "", " ")
	if err != nil {
		log.Println("Could not marshall coordinator pids slice to json. Error: ", err.Error())
		return
	}

	if err := os.WriteFile("coordinator_pids.json", jsonStr, 0644); err != nil {
		log.Println("Could not write to coordinator pids json file. Error: ", err.Error())
		return
	}
}
