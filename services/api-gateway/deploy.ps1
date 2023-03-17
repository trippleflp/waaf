Copy-Item -Recurse -Force ../../libs .

docker build -t localhost:5001/waaf/services/api .
docker push localhost:5001/waaf/services/api

Remove-Item -Recurse -Force ./libs

kubectl replace -f ./deployment/deployment.yaml --force