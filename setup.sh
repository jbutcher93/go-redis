#!/bin/bash

k3d cluster create redis --wait
kubectl config set-context k3d-redis
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm upgrade --install redis bitnami/redis
kubectl apply -f manifests
while true; do
    # Get the current status of the pod
    STATUS=$(kubectl get deploy/redis-client -o jsonpath='{.status.readyReplicas}')

    # Check if the status is "Running"
    if [ "$STATUS" == "1" ]; then
        echo "Pod is now running."
        break
    fi

    # If the status is not "Running", wait for a few seconds before checking again
    sleep 5
done
kubectl port-forward svc/redis-client 8080:80