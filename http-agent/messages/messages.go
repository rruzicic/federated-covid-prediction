package messages

import (
	"io"
	"log"
	"net/http"
)

const pyServerAdress = "http://localhost:6900"

func MakeMatrix(numRows int, numColumns int) [][]float32 {
	matrix := make([][]float32, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]float32, numColumns)
	}
	return matrix
}

func HelloWorld() error {
	res, err := http.Get(pyServerAdress + "/")
	if err != nil {
		log.Println("Could not get response from Python server")
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read the response body")
	}

	log.Println("Response from the Python server: " + string(body))
	return nil
}
