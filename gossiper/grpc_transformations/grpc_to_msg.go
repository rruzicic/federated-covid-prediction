package grpctransformations

import (
	grpc_messages "github.com/rruzicic/federated-covid-prediction/grpc"
	http_messages "github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

func makeMatrix(array []float32, rows int, cols int) [][]float32 {
	matrix := make([][]float32, rows)
	for i := 0; i < rows; i++ {
		rowArray := make([]float32, cols)
		for j := 0; j < cols; j++ {
			rowArray[j] = array[i*cols+j]
		}
		matrix[i] = rowArray
	}

	return matrix
}

func GRPCWeightsToMessageWeights(grpcWeights *grpc_messages.GRPCWeights) http_messages.WeightsResponse {
	return http_messages.WeightsResponse{
		Hidden_weights: makeMatrix(grpcWeights.HiddenWeights, 20, 10),
		Output_weights: makeMatrix(grpcWeights.OutputWeights, 10, 1),
	}
}
