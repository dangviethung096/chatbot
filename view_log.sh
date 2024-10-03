# find the newest log file
LOG_FILE=$(ls -t ./logs/*.log | head -n 1)

tail -n 100 -f $LOG_FILE | ccze -A