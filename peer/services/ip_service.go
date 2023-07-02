package services

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadPeerAddresses() ([]string, error) {
	if _, err := os.Stat("peer_addresses.yaml"); err != nil {
		return []string{}, nil
	}

	peersFile, err := os.ReadFile("peer_addresses.yaml")
	if err != nil {
		log.Println("Could not read peer addresses file. Error: ", err.Error())
		return nil, err
	}

	var peerAddresses []string
	if err := yaml.Unmarshal(peersFile, peerAddresses); err != nil {
		log.Println("Could not unmarshall addresses")
		return nil, err
	}

	return peerAddresses, nil
}

func AddPeerAddress(address string) error {
	peers, err := LoadPeerAddresses()
	if err != nil {
		log.Println("Could not load all peers. Error: ", err.Error())
		return err
	}

	peers = append(peers, address)

	yamlString, err := yaml.Marshal(peers)
	if err != nil {
		log.Println("Could not marshall new peer. Error: ", err.Error())
		return err
	}

	if err := os.WriteFile("peer_addresses.yaml", yamlString, 0644); err != nil {
		log.Println("Could not write to addresses yaml. Error: ", err.Error())
		return err
	}

	return nil
}

func RemovePeerAddress(address string) error {
	peers, err := LoadPeerAddresses()
	if err != nil {
		log.Println("Could not load all peers. Error: ", err.Error())
		return err
	}

	var newPeers []string
	for idx, peer := range peers {
		if peer == address {
			newPeers = append(peers[:idx], peers[idx+1:]...)
			break
		}
	}

	yamlString, err := yaml.Marshal(newPeers)
	if err != nil {
		log.Println("Could not marshall new peer list. Error: ", err.Error())
		return err
	}

	if err := os.WriteFile("peer_addresses.yaml", yamlString, 0644); err != nil {
		log.Println("Could not write to addresses yaml. Error: ", err.Error())
		return err
	}

	return nil
}

func NumberOfPeers() (int, error) {
	peers, err := LoadPeerAddresses()
	if err != nil {
		log.Println("Could not load peer addresses. Error: ", err)
		return -1, err
	}

	return len(peers), nil
}
