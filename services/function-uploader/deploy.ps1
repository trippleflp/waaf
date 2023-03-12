docker build -t localhost:5001/waaf/services/uploader .
docker push localhost:5001/waaf/services/uploader
kubectl delete deployment waaf-services-uploader
kubectl apply -f ./deployment/deployment.yaml