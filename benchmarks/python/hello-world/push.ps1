docker build -t localhost:5001/waaf-benchmark/python/helloworld-alpine -f "Dockerfile.alpine" .
docker push localhost:5001/waaf-benchmark/python/helloworld-alpine

docker build -t localhost:5001/waaf-benchmark/python/helloworld-distroless -f "Dockerfile.distroless" .
docker push localhost:5001/waaf-benchmark/python/helloworld-distroless