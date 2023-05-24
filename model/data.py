from typing import Tuple
import numpy as np
import pandas as pd

from model.globals import DATA_PATH

def load_data() -> Tuple[np.ndarray, np.ndarray]:
    """
    Loads the data and returns the preprocessed versions of it
    """
    df = pd.read_csv(DATA_PATH)

    labels = df["COVID-19"].to_numpy()
    labels = np.where(labels == 'Yes', 1, 0).reshape(-1, 1)

    data = df.drop("COVID-19", axis=1).to_numpy()
    data = np.where(data == 'Yes', 1, 0)

    return data, labels
