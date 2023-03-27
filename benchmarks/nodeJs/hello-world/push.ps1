docker build -t localhost:5001/waaf-benchmark/nodejs/helloworld-alpine -f "Dockerfile.alpine" .
docker push localhost:5001/waaf-benchmark/nodejs/helloworld-alpine

docker build -t localhost:5001/waaf-benchmark/nodejs/helloworld-distroless -f "Dockerfile.distroless" .
docker push localhost:5001/waaf-benchmark/nodejs/helloworld-distroless