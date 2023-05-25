from flask import Flask
from flask import Response
from flask import request

from model.data import load_data

DATA, LABELS = load_data()
APP = Flask(__name__)

@APP.get('/')
def hello_world():
    return "<p>I am alive</p>"


def main():
    global APP
    APP.run(host="localhost", port=6900, debug=True)

if __name__ == '__main__':
    main()