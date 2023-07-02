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
		log.Println("Could not read peer addresses file. Error: ", err)
		return nil, err
	}

	var peerAddresses []string
	if err := yaml.Unmarshal(peersFile, peerAddresses); err != nil {
		log.Println("Could not unmarshall addresses")
		return nil, err
	}

	return peerAddresses, nil
}

func NumberOfPeers() (int, error) {
	peers, err := LoadPeerAddresses()
	if err != nil {
		log.Println("Could not load peer addresses. Error: ", err)
		return -1, err
	}

	return len(peers), nil
}
