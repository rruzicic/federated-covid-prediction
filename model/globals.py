import yaml

with open("model/config.yaml", "r", encoding="utf-8") as f:
    config = yaml.safe_load(f)

DATA_PATH = config["data_path"]
LEARN_RATE = config["learn_rate"]
LAMBDA = config["lambda"]
