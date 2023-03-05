$reg_name = 'kind-registry'
$reg_port = '5001'
$temp_regestry_yaml = './local-registry.temp.yaml'
$registry_yaml = './local-registry.yaml'
$registry_config_map = './local-registry.configmap.yaml'
$temp_registry_config_map = './local-registry.configmap.temp.yaml'

Write-Output "Creating local registry"
docker run -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" registry:2
Write-Output "Local registry created"

$text = Get-Content ${registry_yaml} -Raw 
$text = $text.Replace('${reg_name}', $reg_name)
$text = $text.Replace('${reg_port}', $reg_port)
$text | Out-File -FilePath $temp_regestry_yaml

Write-Output "Create kind cluster"

kind create cluster --config ${temp_regestry_yaml} --image ghcr.io/liquid-reply/kind-crun-wasm:v1.23.0
Write-Output "Kind cluster created"

Write-Output "Connect registry to cluster"
docker network connect "kind" "${reg_name}"

$text = Get-Content ${registry_config_map} -Raw
$text = $text.Replace('${reg_port}', $reg_port)
$text | Out-File -FilePath $temp_registry_config_map
kubectl apply -f ${temp_registry_config_map}
Write-Output "Connection established"


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
