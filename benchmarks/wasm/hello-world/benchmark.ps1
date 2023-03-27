echo "Receive size of wasm image in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/wasm/helloworld

echo "Nop test: "
$timer = [Diagnostics.Stopwatch]::StartNew()
$myjson = '{"apiVersion":"v1","kind":"Pod","metadata":{"name":"wasm-test","annotations":{"module.wasm.image/variant":"compat-smart"}},"spec":{"containers":[{"image":"localhost:5001/waaf-benchmark/helloworld:latest","name":"wasm-test","resources":{}}],"restartPolicy":"Never","runtimeClassName":"crun"}}' | ConvertTo-Json



#kubectl run -it --rm --restart=Never wasi-demo --image=hydai/wasm-wasi-example:with-wasm-annotation --annotations="module.wasm.image/variant=compat" --overrides=$myjson /wasi_example_main.wasm 50000000
kubectl run -it --rm --restart=Never go-hello --image=localhost:5001/waaf-benchmark/helloworld:latest --overrides=$myjson --annotations="module.wasm.image/variant=compat"
#kubectl run -it --rm --restart=Never wasm-hello -f deployment.yaml
#kubectl apply -f deployment.yaml
$timer.stop()
echo $timer.Elapsed

echo "Nop test docker: "
$timer = [Diagnostics.Stopwatch]::StartNew()
for ($i = 1; $i -lt 50; $i++) {
    wasmedge.exe target/wasm32-wasi/release/hello-world.wasm
}
$timer.stop()
echo $timer.Elapsed