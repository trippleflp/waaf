docker build -t localhost:5001/waaf/services/deployer .
docker push localhost:5001/waaf/services/deployer
kubectl replace -f ./deployment/deployment.yaml --force