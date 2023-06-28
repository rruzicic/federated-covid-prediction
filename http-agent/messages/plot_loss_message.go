package messages

import (
	"io"
	"log"
	"net/http"
)

func PlotLoss() ([]byte, error) {
	res, err := http.Get("localhost:6900/plot-loss")
	if err != nil {
		log.Println("Could not send plot loss singal. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read from response body. Error: ", err)
		return nil, err
	}

	return body, nil
}
