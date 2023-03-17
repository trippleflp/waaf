Copy-Item -Recurse -Force ../../libs .
Copy-Item -Recurse -Force ../api-gateway .
Copy-Item -Recurse -Force ../deployer .

docker build -t localhost:5001/waaf/services/fn-group .
docker push localhost:5001/waaf/services/fn-group

Remove-Item -Recurse -Force ./libs
Remove-Item -Recurse -Force ./api-gateway
Remove-Item -Recurse -Force ./deployer

kubectl replace -f ./deployment/deployment.yaml --force