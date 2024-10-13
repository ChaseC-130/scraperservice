import random
import requests
import json
import time

urls = ["https://google.com", "https://phaidra.ai", "https://amazon.com", "https://test.invalid"]



# Set to the pod IP of scrapersvc
target = "http://localhost:8080"



def request():
    selected_url = random.choice(urls)
    payload = {"url": selected_url}
    try:
        response = requests.post(target, data=json.dumps(payload))

        print(f"Response: {response} - Target: {payload}")
    except exception as e:
        print(f"An error occurred: {e}")

while True:
    request()
    time.sleep(5)
