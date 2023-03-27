echo "Receive size in bytes: "

docker inspect -f "{{ .Size }}" localhost:5001/waaf-benchmark/go/helloworld

#echo "Nop test k8s: "
#$timer = [Diagnostics.Stopwatch]::StartNew()
#kubectl run -it --rm --restart=Never go-hello --image=localhost:5001/waaf-benchmark/go/helloworld
#$timer.stop()
#echo $timer.Elapsed

echo "Nop test docker: "
$timer = [Diagnostics.Stopwatch]::StartNew()
for ($i = 1; $i -lt 50; $i++) {
    docker run -it --rm localhost:5001/waaf-benchmark/go/helloworld
}
$timer.stop()
echo $timer.Elapsed