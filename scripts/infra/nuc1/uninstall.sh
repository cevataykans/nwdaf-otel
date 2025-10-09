#!/bin/bash

current_dir=$(pwd)

# Paths to the directories for the corresponding application (Path ending in the directory)
# Manually adapt to local setup (TODO: Clone istio and Aether if necessary BUT: Values still need to be adapted manually)
ISTIO_DIR=/home/sevinc/nuc2-istio
AETHER_DIR=/home/sevinc/aether-onramp-3-1-0
# AETHER_DIR=/home/sevinc/aether-onramp/

echo "****** AETHER UNINSTALLATION ******"
cd "$current_dir"
kubectl delete -f scripts/tempo-service-monitor.yaml
kubectl delete -f scripts/otel-service-monitor.yaml
cd "$AETHER_DIR"
make monitor-uninstall
make aether-5gc-uninstall
cd "$current_dir"

cd "$ISTIO_DIR"
echo "****** REMOVE ISTIO ******"
kubectl delete -f istio-1.17.8/istio-telemetry.yaml -n istio-system
cd "$current_dir"

#echo "****** REMOVE JAEGER CONFIG ******"
#kubectl delete -f scripts/jaeger_config.yaml       #jaeger_config.yaml

echo "****** REMOVE TEMPOOOOO ******"
kubectl delete -f scripts/tempo.yaml

echo "****** REMOVE FILTERED ELASTIC ******"
kubectl delete -f scripts/collector_filtered_elastic.yaml

echo "****** REMOVE OTEL ******"
kubectl delete -f https://github.com/open-telemetry/opentelemetry-operator/releases/download/v0.136.0/opentelemetry-operator.yaml
sleep 2m

echo "****** REMOVE CERT MANAGER ******"
kubectl delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.18.2/cert-manager.yaml
sleep 1m

echo "Checking all remaining pods before k8s uninstallation ..."
kubectl get pods --all-namespaces

cd "$AETHER_DIR"
echo "****** REMOVE K8S ******"
make aether-k8s-uninstall
cd "$current_dir"