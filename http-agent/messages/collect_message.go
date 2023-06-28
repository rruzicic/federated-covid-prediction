package messages

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func Collect(collectJson []byte) (int, []byte, error) {
	res, err := http.NewRequest("POST", "localhost:6900/collect", bytes.NewBuffer(collectJson))
	if err != nil {
		log.Println("Could not post to /collect. Error: ", err)
		return -1, nil, err
	}

	statusCode := res.Response.StatusCode

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read from response body. Error: ", err)
		return -1, nil, err
	}

	return statusCode, body, nil
}
