#!/bin/bash
trap "rm server;kill 0" EXIT

# 启动缓存服务器 最后一次会启动api服务器
go build -o server
./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=1 &

sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &

# sleep 2
# echo ">>> second test"
# curl "http://localhost:9999/api?key=Tom"

wait