import numpy as np


class ModelStruct:
    def __init__(self):
        self.hidden_weights: np.ndarray
        self.output_weights: np.ndarray
        self.hidden_transfer: np.ndarray
        self.hidden_activation: np.ndarray
        self.output_transfer: np.ndarray
        self.output: np.ndarray


def _tanh(tensor: np.ndarray) -> np.ndarray:
    return np.tanh(tensor)


def _d_tanh(tensor: np.ndarray) -> np.ndarray:
    return 1 - _tanh(tensor)**2


def _sigmoid(tensor: np.ndarray) -> np.ndarray:
    return np.where(tensor >= 0, 1 / (1 + np.exp(-tensor)), np.exp(tensor) / (1 + np.exp(tensor)))


def _d_sigmoid(tensor: np.ndarray) -> np.ndarray:
    return _sigmoid(tensor) * (1 - _sigmoid(tensor))
