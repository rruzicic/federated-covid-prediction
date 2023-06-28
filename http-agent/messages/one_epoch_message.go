package messages

import (
	"io"
	"log"
	"net/http"
)

func OneEpoch() ([]byte, error) {
	res, err := http.Get("http://localhost:6900/one-epoch")
	if err != nil {
		log.Println("Could not send one-epoch signal. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read from response body. Error: ", err)
		return nil, err
	}

	return body, err
}
