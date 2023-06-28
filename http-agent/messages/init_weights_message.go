package messages

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func InitWeights(weights []byte) ([]byte, error) {
	res, err := http.NewRequest("POST", "localhost:6900/init", bytes.NewBuffer(weights))
	if err != nil {
		log.Println("Could not send weights to py server. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic("Could not read from response body. Error: ", err)
		return nil, err
	}

	return body, err
}
