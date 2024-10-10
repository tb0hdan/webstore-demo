#!/bin/bash

curl -H 'Content-Type: application/json' -X POST -d '{"discount": "10", "ProductSale":[{"id": "1","quantity": 10},{"id":"2", "quantity":23}]}' http://localhost:8080/api/v1/sales

