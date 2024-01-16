import os
from typing import Any
from typing import Dict

import requests
from dotenv import load_dotenv

load_dotenv()


def send_message(message: str) -> Dict[str, Any]:
    headers = {
        "Authorization": "Bearer " + os.getenv("LINE_TOKEN"),
    }

    data = {
        "message": message,
    }

    response = requests.post(
        "https://notify-api.line.me/api/notify", headers=headers, data=data
    )

    return response.json()
