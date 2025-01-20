#!/bin/bash

apply_migrations() {
    echo "Applying migrations..."
    goose up
    
    if [ $? -eq 0 ]; then
        echo "Migrations applied successfully."
    else
        echo "Error applying migrations."
        exit 1
    fi
}

apply_migrations
