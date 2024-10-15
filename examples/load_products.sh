#!/bin/bash

curl -X POST -d '{"id": "1", "name": "Kettle", "price": 49.99}' http://localhost:8080/api/v1/products

curl -H 'Content-Type: application/json' -X POST -d '{"id": "2", "name": "Booop", "price": 10.99}' http://localhost:8080/api/v1/products

