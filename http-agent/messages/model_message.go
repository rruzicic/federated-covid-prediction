package messages

import (
	"io"
	"log"
	"net/http"
)

func Model() ([]byte, error) {
	// retval is the byte array of the json of the ModelStruct from python
	res, err := http.Get("localhost:6900/model")
	if err != nil {
		log.Println("Could not send get request to python server. Error: ", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read response body. Error: ", err)
		return nil, err
	}

	return body, nil
}
