package messages

import (
	"encoding/json"
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

type WeightsResponse struct {
	Hidden_weights [][]float32 `json:"hidden_weights"`
	Output_weights [][]float32 `json:"output_weights"`
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
		return err
	}

	log.Println("Response from the Python server: " + string(body))
	return nil
}

func GetRandomWeights() (WeightsResponse, error) {
	var weightsResponse WeightsResponse
	res, err := http.Get(pyServerAdress + "/random-weights")
	if err != nil {
		log.Println("Could not get random weights")
		return weightsResponse, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read the response body")
		return weightsResponse, err
	}

	if err := json.Unmarshal(body, &weightsResponse); err != nil {
		log.Println("Could not unmarshall response body into response structure. Error: ", err.Error())
	}

	return weightsResponse, nil
}
