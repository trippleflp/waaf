$reg_name = 'kind-registry'
$reg_port = '5001'
$temp_regestry_yaml = './local-registry.temp.yaml'
$registry_yaml = './local-registry.yaml'
$registry_config_map = './local-registry.configmap.yaml'
$temp_registry_config_map = './local-registry.configmap.temp.yaml'

Write-Output "Creating local registry"
docker run -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" registry:2
Write-Output "Local registry created"

Write-Output "Create node image"
$node_image = "kind-node-image-with-plugins:0.0.1"
docker build -t ${node_image} .
Write-Output "Created node image"

$text = Get-Content ${registry_yaml} -Raw 
$text = $text.Replace('${reg_name}', $reg_name)
$text = $text.Replace('${reg_port}', $reg_port)
$text = $text.Replace('${image}', $node_image)
$text | Out-File -FilePath $temp_regestry_yaml

Write-Output "Create kind cluster"
kind create cluster --config ${temp_regestry_yaml}

helm repo add kwasm http://kwasm.sh/kwasm-operator/
helm install -n kwasm --create-namespace kwasm-operator kwasm/kwasm-operator
kubectl annotate node --all kwasm.sh/kwasm-node=true
kubectl apply -f https://raw.githubusercontent.com/KWasm/kwasm-node-installer/main/example/test-job.yaml
kubectl logs job/wasm-test --follow

kubectl create clusterrolebinding serviceaccounts-cluster-admin --clusterrole=cluster-admin --group=system:serviceaccounts
Write-Output "Kind cluster created"

Write-Output "Connect registry to cluster"
docker network connect "kind" "${reg_name}"

$text = Get-Content ${registry_config_map} -Raw
$text = $text.Replace('${reg_port}', $reg_port)
$text | Out-File -FilePath $temp_registry_config_map
kubectl apply -f ${temp_registry_config_map}
Write-Output "Connection established"

Write-Output "Patch ingress controller"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=90s
Write-Output "Patch finished"

Write-Output "Test local registry connection"
docker pull gcr.io/google-samples/hello-app:1.0
docker tag gcr.io/google-samples/hello-app:1.0 localhost:${reg_port}/hello-app:1.0
docker push localhost:${reg_port}/hello-app:1.0
kubectl create deployment hello-server --image=localhost:${reg_port}/hello-app:1.0
kubectl rollout status deployment/hello-server
kubectl get rs
kubectl delete deploy hello-server
Write-Output "Connection to local registry successfull"


Write-Output "Test WASM behaivour"
kubectl run -it --rm --restart=Never wasi-demo --image=hydai/wasm-wasi-example:with-wasm-annotation --annotations="module.wasm.image/variant=compat" /wasi_example_main.wasm 50000000
Write-Output "Test finished"

Write-Output "Cluster setup finished"
