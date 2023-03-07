docker build -t localhost:5001/waaf/services/deployer .
docker push localhost:5001/waaf/services/deployer
kubectl apply -f ./deployment/deployment.yaml