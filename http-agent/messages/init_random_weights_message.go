package messages

import (
	"io"
	"log"
	"net/http"
)

func GetRandomWeights() ([]byte, error) {
	// retval is a byte array of the randomly generated weights
	res, err := http.Get("http://localhost:6900/random-weights")
	if err != nil {
		log.Println("Could not get random weights. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read response body. Error: ", err)
		return nil, err
	}

	return body, nil
}
