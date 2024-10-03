#!/bin/bash

# Directory to store log files
LOG_DIR="./logs"

# Create the log directory if it doesn't exist
mkdir -p "$LOG_DIR"

# Get the current date and time
CURRENT_TIME=$(date "+%Y-%m-%d_%H-%M-%S")

# Log file name
LOG_FILE="$LOG_DIR/mobifone_chatbot_log_$CURRENT_TIME.log"

# Function to find and kill the running app
kill_app() {
    # Find the PID of the running app
    APP_PID=$(pgrep chatbot)
    
    if [ -n "$APP_PID" ]; then
        echo "Stopping running instance of chatbot with PID $APP_PID"
        kill "$APP_PID"
        # Wait for the process to terminate
        wait "$APP_PID" 2>/dev/null
        echo "Stopped chatbot"
    else
        echo "No running instance of chatbot found"
    fi
}

# Function to start the app
start_app() {
    echo "Starting chatbot"
    sudo ./chatbot > "$LOG_FILE" 2>&1 &
    APP_PID=$!
    echo "Started chatbot with PID $APP_PID. Log saved to $LOG_FILE"
}

# Stop the app if it's running
kill_app

# Start the app
start_app