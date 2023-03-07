kubectl delete configmap nginx-stack-config
kubectl create configmap nginx-stack-config --from-file ..\lib\index.js --from-file ..\conf\nginx.conf