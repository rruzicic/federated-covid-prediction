package messages

import (
	"io"
	"log"
	"net/http"
)

func AllPeersDone() ([]byte, error) {
	res, err := http.Get("http://localhost:6900/all-peers-sent-weights")
	if err != nil {
		log.Println("Could not send all peers done signal. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read from response body. Erro: ", err)
		return nil, err
	}

	return body, nil
}
