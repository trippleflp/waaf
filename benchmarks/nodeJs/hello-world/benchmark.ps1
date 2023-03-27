echo "Receive size of alpine image in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/nodejs/helloworld-alpine

echo "\nReceive size of distroless image in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/nodejs/helloworld-distroless




echo "Nop test: "
$timer = [Diagnostics.Stopwatch]::StartNew()
kubectl run -it --rm --restart=Never node-hello --image=localhost:5001/waaf-benchmark/nodejs/helloworld-distroless
$timer.stop()
echo $timer.Elapsed

#echo "Nop test docker: "
#$timer = [Diagnostics.Stopwatch]::StartNew()
#for ($i = 1; $i -lt 50; $i++) {
#    docker run -it --rm localhost:5001/waaf-benchmark/nodejs/helloworld-distroless
#}
#$timer.stop()
#echo $timer.Elapsed

echo "Nop test node: "
$timer = [Diagnostics.Stopwatch]::StartNew()
for ($i = 1; $i -lt 50; $i++) {
    node index.js
}
$timer.stop()
echo $timer.Elapsed