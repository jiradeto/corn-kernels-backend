import requests
import json
import sqlite3
import random
import datetime
import time
from datetime import datetime

def generate_random_date():
    random_day = random.randint(1, 20)
    random_datetime = datetime(2023, 2, random_day, random.randint(0, 23), random.randint(0, 59), random.randint(0, 59))
    return random_datetime.strftime('%Y-%m-%dT%H:%M:%S.%f') + "+07:00"

def make_call(product_id, description, quantity, movement_type):
    url = "http://localhost:8080/api/v1/stock"
    payload = json.dumps({
        "ProductID": str(product_id),
        "Description": description,
        "MovementType": movement_type,
        "Quantity": quantity,
        "CreatedAt": generate_random_date()
    })
    headers = {
        'Content-Type': 'application/json'
    }
    response = requests.request("POST", url, headers=headers, data=payload)
    print(response.text)


def insert_fake_data(conn):
    product_ids = [1,2,3]
    descriptions = ['Product 1 movement', 'Product 2 movement', 'Product 3 movement', 'Product 4 movement', 'Product 5 movement', 'Product 6 movement']
    movement_types = ['in', 'out']
    for _ in range(100):
        product_id = random.choice(product_ids)
        description = random.choice(descriptions)
        quantity = None
        movement_type = random.choice(movement_types)
        if movement_type == 'in':
            quantity = random.randint(1000, 2000)
        else:
            quantity = random.randint(10, 500)
        make_call(product_id, description, quantity, movement_type)
        
    conn.commit()


if __name__ == '__main__':
    conn = sqlite3.connect("./local.db")
    cursor = conn.cursor()
    insert_fake_data(conn)
    print("COOL")