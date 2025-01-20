#!/bin/bash

usage() {
    echo "Usage: $0 <migration_name>"
    echo
    echo "Arguments:"
    echo "  migration_name   The name of the migration to create."
    echo
    echo "Example:"
    echo "  $0 add_users_table"
    exit 1
}

create_migration() {
    if [ -z "$1" ]; then
        echo "Error: Migration name is required."
        usage
    fi

    echo "Creating new migration: $1"
    goose create "$1" sql
    
    if [ $? -eq 0 ]; then
        echo "Migration created successfully."
    else
        echo "Error creating migration."
        exit 1
    fi
}

if [ $# -eq 0 ]; then
    usage
else
    create_migration "$1"
fi
