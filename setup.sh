k3d cluster start
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm upgrade --install redis bitnami/redis
kubectl apply -f manifests
sleep 5
kubectl port-forward svc/redis-master 6380:6379
kubectl port-forward svc/redis-replicas 6379:6379