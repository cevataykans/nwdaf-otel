#!/bin/bash

current_dir=$(pwd)
kubectl delete -f scripts/udm_scaled_object.yaml

AETHER_DIR=/home/sevinc/cores/aether-onramp-3-1-0
# AETHER_DIR=/home/sevinc/aether-onramp/

echo "****** AETHER UNINSTALLATION ******"
cd "$AETHER_DIR"
make aether-gnbsim-uninstall
make monitor-uninstall
cd "$current_dir"
make stop-nwdaf
cd "$AETHER_DIR"
make aether-5gc-uninstall
cd "$current_dir"

helm uninstall keda -n keda

echo "****** REMOVE TEMPOOOOO ******"
kubectl delete -f scripts/tempo.yaml

echo "****** REMOVE FILTERED ELASTIC ******"
kubectl delete -f scripts/collector_filtered_elastic.yaml

echo "****** REMOVE OTEL ******"
kubectl delete -f https://github.com/open-telemetry/opentelemetry-operator/releases/download/v0.136.0/opentelemetry-operator.yaml

echo "****** REMOVE CERT MANAGER ******"
kubectl delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.18.2/cert-manager.yaml

echo "Checking all remaining pods before k8s uninstallation ..."
kubectl get pods --all-namespaces

sleep 30
cd "$AETHER_DIR"
echo "****** REMOVE K8S ******"
make aether-k8s-uninstall
cd "$current_dir"