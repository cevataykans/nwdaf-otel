#!/bin/bash

current_dir=$(pwd)

# Paths to the directories for the corresponding application (Path ending in the directory)
# Manually adapt to local setup (TODO: Clone istio and Aether if necessary BUT: Values still need to be adapted manually)
ISTIO_DIR=/home/sevinc/jungmann/istio
AETHER_DIR=/home/sevinc/aether-onramp/

echo "****** REMOVE JAEGER CONFIG ******"
kubectl delete -f jungmann/setup_scripts/jaeger_config_3.yaml       #jaeger_config.yaml

echo "****** FILTERED ELASTIC INSTALLATION ******"
kubectl delete -f jungmann/setup_scripts/collector_filtered_elastic.yaml   # collector_filtered.yaml

echo "****** REMOVE OTEL INSTALLATION ******"
#kubectl delete -f otel_operator_go_dec.yaml
kubectl delete -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
sleep 2m

echo "****** REMOVE CERT MANAGER ******"
kubectl delete -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml
sleep 1m

cd "$AETHER_DIR"
echo "****** REMOVE K8S ******"
make aether-k8s-uninstall
cd "$current_dir"

sleep 30s