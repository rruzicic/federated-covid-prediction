from flask import Flask
from flask import Response
from flask import request

from model.data import load_data
from model.neural_network import ModelStruct, _init_model_weights

DATA, LABELS = load_data()
MODEL = ModelStruct()
APP = Flask(__name__)

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
    and updates the global model
    """
    global MODEL

    weights_dict = request.get_json(force=True)
    MODEL = _init_model_weights(MODEL, weights_dict)
    return Response(status=200)


def main():
    global APP
    APP.run(host="localhost", port=6900, debug=True)

if __name__ == '__main__':
    main()
