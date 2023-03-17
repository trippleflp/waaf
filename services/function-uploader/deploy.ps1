docker build -t localhost:5001/waaf/services/uploader .
docker push localhost:5001/waaf/services/uploader
kubectl replace -f ./deployment/deployment.yaml --force