kubectl delete configmap nginx-stack-config --namespace waaf
kubectl create configmap nginx-stack-config --namespace waaf --from-file ..\lib\index.js --from-file ..\conf\nginx.conf