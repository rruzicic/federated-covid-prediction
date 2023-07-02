import yaml

with open("model/config.yaml", "r", encoding="utf-8") as f:
    config = yaml.safe_load(f)

DATA_PATH = config["data_path"]
LEARN_RATE = config["learn_rate"]
LAMBDA = config["lambda"]
EPOCHS_PER_REQUEST = config["epochs_per_request"]
TOTAL_EPOCHS = config["total_epochs"]
