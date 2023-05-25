from flask import Flask
from flask import Response
from flask import request

from model.data import load_data
from model.neural_network import ModelStruct, _init_model_weights, personalized_weight_update

DATA, LABELS = load_data()
MODEL = ModelStruct()
APP = Flask(__name__)
OTHERS_HIDDEN_WEIGHTS = []
OTHERS_OUTPUT_WEIGHTS = []

@APP.get('/')
def hello_world():
    return "<p>I am alive</p>"


@APP.get('/model')
def get_model():
    return MODEL.to_dict()


@APP.post('/init')
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


@APP.post('/collect')
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
        peers: int
    }
    """
    global MODEL, OTHERS_HIDDEN_WEIGHTS, OTHERS_OUTPUT_WEIGHTS

    req = request.get_json(force=True)
    peers = req['peers']
    weights = {
        "hidden_weights": req["hidden_weights"],
        "output_weights": req["output_weigths"]
    }

    OTHERS_HIDDEN_WEIGHTS.append(weights["hidden_weights"])
    OTHERS_OUTPUT_WEIGHTS.append(weights["output_weights"])

    if (OTHERS_HIDDEN_WEIGHTS == peers-1) and (OTHERS_OUTPUT_WEIGHTS == peers-1):
        MODEL = personalized_weight_update(MODEL, OTHERS_HIDDEN_WEIGHTS, OTHERS_OUTPUT_WEIGHTS)
        OTHERS_HIDDEN_WEIGHTS = []
        OTHERS_OUTPUT_WEIGHTS = []

    return Response(200)


def main():
    global APP
    APP.run(host="localhost", port=6900, debug=True)

if __name__ == '__main__':
    main()
