mkdir -p logs
go build -o main .
nohup ./main > ./logs/payment-recipt-generator.log 2>&1 &
# Print the PID of the process
echo $! > ./logs/payment-recipt-generator.pid