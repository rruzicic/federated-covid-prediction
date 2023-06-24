from typing import List, Dict
import numpy as np
from model.globals import LEARN_RATE, LAMBDA


class ModelStruct:
    """
    Struct in which all the relevant model data will be kept for easy access
    """

    def __init__(self):
        self.hidden_weights: np.ndarray = np.ndarray([])
        self.output_weights: np.ndarray = np.ndarray([])

        self.hidden_transfer: np.ndarray = np.ndarray([])
        self.hidden_activation: np.ndarray = np.ndarray([])

        self.output_transfer: np.ndarray = np.ndarray([])
        self.output: np.ndarray = np.ndarray([])

        self.loss: List[float] = []

    def to_dict(self):
        return {
            "hidden_weights": self.hidden_weights.tolist(),
            "output_weights": self.output_weights.tolist(),
            "output": self.output.tolist(),
            "loss": self.loss,
        }


def _tanh(tensor: np.ndarray) -> np.ndarray:
    return np.tanh(tensor)


def _d_tanh(tensor: np.ndarray) -> np.ndarray:
    return 1 - _tanh(tensor) ** 2


def _relu(tensor: np.ndarray) -> np.ndarray:
    return np.where(tensor >= 0, tensor, 0)


def _d_relu(tensor: np.ndarray) -> np.ndarray:
    return np.where(tensor >= 0, 1, 0)


def _sigmoid(tensor: np.ndarray) -> np.ndarray:
    return np.where(
        tensor >= 0, 1 / (1 + np.exp(-tensor)), np.exp(tensor) / (1 + np.exp(tensor))
    )


def _d_sigmoid(tensor: np.ndarray) -> np.ndarray:
    return _sigmoid(tensor) * (1 - _sigmoid(tensor))


def _mse_loss(model: ModelStruct, labels: np.ndarray) -> float:
    return np.sum((model.output - labels) ** 2) / (2 * labels.shape[0])


def _init_model_weights_debug(model: ModelStruct, data: np.ndarray) -> ModelStruct:
    """
    Hardcoded model weights init for debugging purposes
    """
    model.hidden_weights = np.random.uniform(-1, 1, size=(data.shape[1], 10))
    model.output_weights = np.random.uniform(-1, 1, size=(10, 1))

    return model


def _init_model_weights(
    model: ModelStruct, weights: Dict[str, np.ndarray]
) -> ModelStruct:
    """
    Given a dict with hidden_weights and output_weights as keys, initializes the model's weights
    """
    model.hidden_weights = np.array(weights["hidden_weights"])
    model.output_weights = np.array(weights["output_weights"])

    return model


def _forwardpass(model: ModelStruct, data: np.ndarray) -> ModelStruct:
    """
    Full transfer layer, updates model with data. Returns modelstruct with updated fields
    """
    model.hidden_transfer = data @ model.hidden_weights
    model.hidden_activation = _relu(model.hidden_transfer)

    model.output_transfer = model.hidden_activation @ model.output_weights
    model.output = _sigmoid(model.output_transfer)

    return model


def _calculate_loss(model: ModelStruct, labels: np.ndarray) -> ModelStruct:
    """
    Calculates the loss and appends it to the model loss
    """
    model.loss.append(_mse_loss(model, labels))
    return model


def _update_weights(
    model: ModelStruct,
    hidden_weights_grads: np.ndarray,
    output_weights_grads: np.ndarray,
    learn_rate: float,
) -> ModelStruct:
    """
    Updates both weights of the model and returns the updated model
    """
    model.hidden_weights -= learn_rate * hidden_weights_grads
    model.output_weights -= learn_rate * output_weights_grads

    return model


def _backpropagation(
    model: ModelStruct, labels: np.ndarray, data: np.ndarray, learn_rate: float
) -> ModelStruct:
    """
    Does backprop to updates model weights. Returns model with updated weights
    """
    d_output = (model.output - labels) / labels.shape[0]
    d_output_transfer = _d_sigmoid(model.output_transfer) * d_output
    d_output_weights = model.hidden_activation.T @ d_output_transfer

    d_hidden_activation = d_output_transfer @ model.output_weights.T
    d_hidden_transfer = _d_relu(model.hidden_transfer) * d_hidden_activation
    d_hidden_weights = data.T @ d_hidden_transfer

    model = _update_weights(model, d_hidden_weights, d_output_weights, learn_rate)

    return model


def one_epoch(model: ModelStruct, data: np.ndarray, labels: np.ndarray) -> ModelStruct:
    """
    Does one pass through the data
    """
    model = _forwardpass(model, data)
    model = _calculate_loss(model, labels)
    model = _backpropagation(model, labels, data, LEARN_RATE)

    return model


def personalized_weight_update(
    model: ModelStruct,
    aggregated_hidden_weights: List[np.ndarray],
    aggregated_output_weights: List[np.ndarray],
) -> ModelStruct:
    """
    Does personalized weight updates to conform to the updated ones
    """
    avg_hidden_weights = (
        1.0
        / (len(aggregated_hidden_weights) + 1)
        * np.add.reduce([*aggregated_hidden_weights, model.hidden_weights])
    )

    sq_norm_hidden_weights = np.sum(
        [
            np.linalg.norm(model.hidden_weights - other_weights) ** 2
            for other_weights in aggregated_hidden_weights
        ]
    )
    model.hidden_weights = avg_hidden_weights + LAMBDA * sq_norm_hidden_weights

    avg_output_weights = (
        1.0
        / (len(aggregated_output_weights) + 1)
        * np.add.reduce([*aggregated_output_weights, model.output_weights])
    )

    sq_norm_output_weights = np.sum(
        [
            np.linalg.norm(model.output_weights - other_weights) ** 2
            for other_weights in aggregated_output_weights
        ]
    )
    model.output_weights = avg_output_weights + LAMBDA * sq_norm_output_weights

    return model
