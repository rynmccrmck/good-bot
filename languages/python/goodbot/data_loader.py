import json
from pathlib import Path

def load_bot_data():
    local_data_path = Path(__file__).parent.parent.parent.parent / "data/goodbots.json"
    pkg_data_path = Path(__file__).parent / "data/goodbots.json"
    if local_data_path.exists():
        data_path = local_data_path
    else:
        data_path = pkg_data_path
    with open(data_path, 'r') as file:
        data = json.load(file)
    return data


data = load_bot_data()
print(data)