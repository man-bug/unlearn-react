#!/bin/bash

echo "Starting Tailwind CSS build process..."
./tailwindcss -i ./static/css/tailwind.css -o ./static/css/tailwind.output.css

echo "Running Tailwind in watch mode..."
./tailwindcss -i ./static/css/tailwind.css -o ./static/css/tailwind.output.css --watch &

# Check if air is installed
if ! command -v air &> /dev/null
then
    echo "air is not installed. Installing now..."
    go install github.com/air-verse/air@latest
fi

echo "Starting Go server with air..."
air

# Kill the Tailwind process when we're done
trap "kill 0" EXIT
