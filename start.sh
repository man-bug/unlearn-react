#!/bin/sh

echo "Current working directory: $(pwd)"
echo "Listing directory contents:"
ls -la

echo "Checking for main executable:"
if [ -f "./main" ]; then
    echo "main executable found"
else
    echo "main executable not found"
fi

echo "Checking PORT environment variable:"
if [ -z "$PORT" ]; then
    echo "PORT is not set, defaulting to 8080"
    export PORT=8080
else
    echo "PORT is set to $PORT"
fi

echo "Listing /app directory contents:"
ls -la /app

echo "Starting application..."
./main 2>&1 | tee /var/log/app.log
