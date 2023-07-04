package services

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AddressAndHost struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type YamlData struct {
	You   AddressAndHost   `yaml:"you"`
	Peers []AddressAndHost `yaml:"peers"`
}

func LoadAddresses() (*YamlData, error) {
	if _, err := os.Stat("peer_addresses.yaml"); err != nil {
		return &YamlData{}, nil
	}

	addressesFile, err := os.ReadFile("peer_addresses.yaml")
	if err != nil {
		log.Println("Could not read peer addresses file. Error: ", err.Error())
		return nil, err
	}

	var yamlAddresses YamlData
	if err := yaml.Unmarshal(addressesFile, &yamlAddresses); err != nil {
		log.Println("Could not unmarshall addresses")
		return nil, err
	}

	return &yamlAddresses, nil
}

func AddPeerAddress(address AddressAndHost) error {
	addresses, err := LoadAddresses()
	if err != nil {
		log.Println("Could not load all peers. Error: ", err.Error())
		return err
	}

	addresses.Peers = append(addresses.Peers, address)

	yamlString, err := yaml.Marshal(addresses)
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

func RemovePeerAddress(address AddressAndHost) error {
	addresses, err := LoadAddresses()
	if err != nil {
		log.Println("Could not load all peers. Error: ", err.Error())
		return err
	}

	for idx, peer := range addresses.Peers {
		if peer == address {
			addresses.Peers = append(addresses.Peers[:idx], addresses.Peers[idx+1:]...)
			break
		}
	}

	yamlString, err := yaml.Marshal(addresses)
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
	addresses, err := LoadAddresses()
	if err != nil {
		log.Println("Could not load peer addresses. Error: ", err)
		return -1, err
	}

	return len(addresses.Peers), nil
}

func GetYourAddress() (*AddressAndHost, error) {
	addresses, err := LoadAddresses()
	if err != nil {
		log.Println("Could not load all peers. Error: ", err.Error())
		return nil, err
	}

	return &addresses.You, nil
}

func GetPeerAddresses() ([]AddressAndHost, error) {
	addresses, err := LoadAddresses()
	if err != nil {
		log.Println("Could not load all peers. Error: ", err.Error())
		return nil, err
	}

	return addresses.Peers, nil
}
