package grpctransformations

import (
	grpc_messages "github.com/rruzicic/federated-covid-prediction/grpc"
	http_messages "github.com/rruzicic/federated-covid-prediction/http-agent/messages"
)

func ravelMatrix(matrix [][]float32, rows int, cols int) []float32 {
	raveled := make([]float32, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			raveled[i*cols+j] = matrix[i][j]
		}
	}

	return raveled
}

func MessageWeightsToGRPCWeights(messageWeights http_messages.WeightsResponse) *grpc_messages.GRPCWeights {
	return &grpc_messages.GRPCWeights{
		HiddenWeights: ravelMatrix(messageWeights.Hidden_weights, 20, 10),
		OutputWeights: ravelMatrix(messageWeights.Output_weights, 10, 1),
	}
}
