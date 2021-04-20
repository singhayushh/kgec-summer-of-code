PID=$(lsof -t -i:7000)
if [[ ! -z "$PID" ]]
then
    kill $PID
fi