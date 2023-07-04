package messages

import (
	"bytes"
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

type ModelResponse struct {
	Hidden_weights [][]float32 `json:"hidden_weights"`
	Output_weights [][]float32 `json:"output_weights"`
	Output         []float32   `json:"output"`
	Loss           []float32   `json:"loss"`
}

type CollectResponse struct {
	Hidden_weights [][]float32 `json:"hidden_weights"`
	Output_weights [][]float32 `json:"output_weights"`
	Peers          int         `json:"peers"`
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

func GetRandomWeights() (WeightsResponse, int, error) {
	var weightsResponse WeightsResponse
	res, err := http.Get(pyServerAdress + "/random-weights")
	if err != nil {
		log.Println("Could not get random weights")
		return weightsResponse, res.StatusCode, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read the response body")
		return weightsResponse, res.StatusCode, err
	}

	if err := json.Unmarshal(body, &weightsResponse); err != nil {
		log.Println("Could not unmarshall response body into response structure. Error: ", err.Error())
		return weightsResponse, res.StatusCode, err
	}

	return weightsResponse, res.StatusCode, nil
}

func GetModel() (ModelResponse, error) {
	var modelResponse ModelResponse
	res, err := http.Get(pyServerAdress + "/model")
	if err != nil {
		log.Println("Could not get random weights")
		return modelResponse, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read the response body")
		return modelResponse, err
	}

	if err := json.Unmarshal(body, &modelResponse); err != nil {
		log.Println("Could not unmarshall response body into response structure. Error: ", err.Error())
		return modelResponse, err
	}

	return modelResponse, nil
}

func AllPeersDone() (int, error) {
	res, err := http.Get(pyServerAdress + "/all-peers-sent-weights")
	if err != nil {
		log.Println("Could not send all peers done signal. Error: ", err)
		return res.StatusCode, err
	}

	return res.StatusCode, nil
}

func DoOneEpoch() (int, error) {
	res, err := http.Get(pyServerAdress + "/one-epoch")
	if err != nil {
		log.Println("Could not send one epoch done signal. Error: ", err)
		return res.StatusCode, err
	}

	return res.StatusCode, nil
}

func GetWeights() (WeightsResponse, int, error) {
	var weightsResponse WeightsResponse
	res, err := http.Get(pyServerAdress + "/weights")
	if err != nil {
		log.Println("Could not get weights")
		return weightsResponse, res.StatusCode, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Could not read the response body")
		return weightsResponse, res.StatusCode, err
	}

	if err := json.Unmarshal(body, &weightsResponse); err != nil {
		log.Println("Could not unmarshall response body into response structure. Error: ", err.Error())
		return weightsResponse, res.StatusCode, err
	}

	return weightsResponse, res.StatusCode, nil
}

func Exit() (int, error) {
	res, err := http.Get(pyServerAdress + "/exit")
	if err != nil {
		log.Println("Could not send exit signal. Error: ", err)
		return res.StatusCode, err
	}

	return res.StatusCode, nil
}

func InitWeights(weights WeightsResponse) (int, error) {
	reqBody, err := json.Marshal(weights)
	if err != nil {
		log.Println("Could not marshal WeightsResponse struct")
		return 500, err
	}

	res, err := http.NewRequest("POST", pyServerAdress+"/init", bytes.NewReader(reqBody))
	if err != nil {
		log.Println("Could not send weights to py server. Error: ", err)
		return res.Response.StatusCode, err
	}

	return res.Response.StatusCode, nil
}

func CollectWeights(collectResponse CollectResponse) (int, error) {
	reqBody, err := json.Marshal(collectResponse)
	if err != nil {
		log.Println("Could not marshal CollectResponse struct")
		return 500, err
	}

	res, err := http.NewRequest("POST", pyServerAdress+"/collect", bytes.NewReader(reqBody))
	if err != nil {
		log.Println("Could not send collect message to py server. Error: ", err)
		return res.Response.StatusCode, err
	}

	return res.Response.StatusCode, nil
}
