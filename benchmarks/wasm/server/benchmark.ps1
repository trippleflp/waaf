echo "Receive size of server image in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/rust/server
