echo "Receive size in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/go/server


d