from flask import Flask
from flask import Response
from flask import request
import numpy as np
import matplotlib.pyplot as plt

from model.data import load_data
from model.neural_network import (
    ModelStruct,
    _init_model_weights,
    personalized_weight_update,
    one_epoch,
)
from globals import EPOCHS_PER_REQUEST

DATA, LABELS = load_data()
MODEL = ModelStruct()
APP = Flask(__name__)
OTHERS_HIDDEN_WEIGHTS = []
OTHERS_OUTPUT_WEIGHTS = []


@APP.get("/")
def hello_world():
    """
    Health check endpoint
    """
    return "<p>I am alive</p>"


@APP.get("/model")
def get_model():
    """
    Returns json with model weights, output and loss
    """
    return MODEL.to_dict()


@APP.post("/init")
def init_weights():
    """
    Endpoint that consumes a json with hidden_weights and output_weights
    and updates the global model.
    JSON: {
        "hidden_weights": [[...], [...]],
        "output_weights": [[...], [...]]
    }
    """
    global MODEL

    weights_dict = request.get_json(force=True)
    MODEL = _init_model_weights(MODEL, weights_dict)
    return Response(status=200)


@APP.post("/collect")
def collect_weights():
    """
    Json has someone elses hidden and output weights similar to init_weights' json
    The json also contains how many models are in the p2p network
    The endpoint then appends the weights to the global list
    If there are as many weights as there are peers-1 (since you are one of them)
    Then it runs personalized weight aggregation and clears both lists
    JSON: {
        "hidden_weights": [[...], [...]],
        "output_weights": [[...], [...]],
        "peers": int
    }
    """
    global MODEL, OTHERS_HIDDEN_WEIGHTS, OTHERS_OUTPUT_WEIGHTS

    req = request.get_json(force=True)
    peers = req["peers"]
    weights = {
        "hidden_weights": req["hidden_weights"],
        "output_weights": req["output_weigths"],
    }

    OTHERS_HIDDEN_WEIGHTS.append(np.array(weights["hidden_weights"]))
    OTHERS_OUTPUT_WEIGHTS.append(np.array(weights["output_weights"]))

    if (OTHERS_HIDDEN_WEIGHTS == peers - 1) and (OTHERS_OUTPUT_WEIGHTS == peers - 1):
        MODEL = personalized_weight_update(
            MODEL, OTHERS_HIDDEN_WEIGHTS, OTHERS_OUTPUT_WEIGHTS
        )
        OTHERS_HIDDEN_WEIGHTS = []
        OTHERS_OUTPUT_WEIGHTS = []

        # to tell the peer that sent you his weights that he was your last
        return Response(201)

    return Response(200)


@APP.get("/all-peers-sent-weights")
def all_peers_sent_weights():
    """
    When /collect is requested, return code is either 200 or 201.
    If it's 201 that means the requester was the last one to send their weights.
    So this endpoint is the way for the requester to notify the server that he was the last one
    """
    return Response(200)


@APP.get("/one-epoch")
def start_one_epoch():
    """
    Tells the current model to do one epoch of learning
    Is actually N epochs
    """
    global MODEL, DATA, LABELS

    for _ in range(EPOCHS_PER_REQUEST):
        MODEL = one_epoch(MODEL, DATA, LABELS)
    return Response(200)


@APP.get("/weights")
def get_model_weights_for_collecting():
    """
    Returns a json with the model's weights
    """
    return {
        "hidden_weights": MODEL.hidden_weights.tolist(),
        "output_weights": MODEL.output_weights.tolist(),
    }

@APP.get("/plot-loss")
def plot_model_loss():
    """
    Plots loss of current model and saves it in this dir as LossImage
    """
    global MODEL
    plt.plot(MODEL.loss)
    plt.savefig("LossImage.png")


def main():
    """
    main func
    """
    global APP
    APP.run(host="localhost", port=6900, debug=True)


if __name__ == "__main__":
    main()
