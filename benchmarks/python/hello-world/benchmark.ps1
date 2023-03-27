echo "Receive size of alpine image in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/python/helloworld-alpine

echo "\nReceive size of distroless image in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/python/helloworld-distroless