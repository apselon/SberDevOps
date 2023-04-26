istioctl install --set profile=minimal -y

kubectl label namespace default istio-injection=enabled
kubectl apply -f K8s/pokemon-deployment.yml

kubectl apply -f K8s/ingress-gateway.yml
kubectl apply -f K8s/ingress-service.yml

kubectl apply -f K8s/egress-gateway.yml
kubectl apply -f K8s/egress-service.yml

kubectl apply -f K8s/nginx-deployment.yml

minikube tunnel
