#!/bin/bash

SQLC_CONFIG="./sql/sqlc.yaml"

if ! command -v sqlc &> /dev/null
then
    echo "sqlc could not be found, please install it first."
    exit 1
fi

echo "Generating Go code from SQL queries..."
sqlc -f "$SQLC_CONFIG" generate

if [ $? -eq 0 ]; then
    echo "Go code generated successfully."
else
    echo "Error generating Go code."
    exit 1
fi
