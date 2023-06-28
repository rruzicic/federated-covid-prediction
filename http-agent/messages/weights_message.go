package messages

import (
	"io"
	"log"
	"net/http"
)

func Weights() ([]byte, error) {
	res, err := http.Get("localhost:6900/weights")
	if err != nil {
		log.Println("Could not get weights. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read form response body. Error: ", err)
		return nil, err
	}

	return body, err
}
