$baseDir = pwd

kubectl create namespace waaf

echo "Create stack config"
cd ./function-group-stack/nginx/deployment
./configmap.ps1
echo "Stack config created"

cd $baseDir

echo "Deploy api gateway"
cd ./services/api-gateway
./deploy.ps1
echo "Api gateway deployed"

cd $baseDir

echo "Deploy authentication"
cd ./services/authentication
./deploy.ps1
echo "Authentication deployed"



cd $baseDir

echo "Deploy deployer"
cd ./services/deployer
./deploy.ps1
echo "Deployer deployed"



cd $baseDir

echo "Deploy function-group"
cd ./services/function-group
./deploy.ps1
echo "Function-Group deployed"



cd $baseDir

echo "Deploy function-uploader"
cd ./services/function-uploader
./deploy.ps1
echo "Function-uploader deployed"

cd $baseDir


kubectl get rs --namespace waaf