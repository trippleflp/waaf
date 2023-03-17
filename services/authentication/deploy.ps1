Copy-Item -Recurse -Force ../../libs .

docker build -t localhost:5001/waaf/services/auth .
docker push localhost:5001/waaf/services/auth

Remove-Item -Recurse -Force ./libs

kubectl replace -f ./deployment/deployment.yaml --force