#!/bin/bash

current_dir=$(pwd)

AETHER_DIR=/home/sevinc/aether-onramp-3-1-0
# AETHER_DIR=/home/sevinc/aether-onramp/

# remove dangling data from disk from older metric installations
sudo rm -rf /opt/local-path-provisioner/
sudo mkdir /opt/local-path-provisioner/

cd "$AETHER_DIR"
echo "****** K8S INSTALLATION ******"
make aether-k8s-install
cd "$current_dir"

sleep 30s

echo "****** CERT MANAGER INSTALLATION ******"
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.18.2/cert-manager.yaml
sleep 1m

echo "****** OTEL INSTALLATION ******"
kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/download/v0.136.0/opentelemetry-operator.yaml
sleep 5m

echo "****** FILTERED ELASTIC INSTALLATION ******"
kubectl apply -f scripts/collector_filtered_elastic.yaml
echo "****** TEMPOOOOOOO ******"
kubectl apply -f scripts/tempo.yaml
#echo "****** JAEGER CONFIG INSTALLATION ******"
#kubectl apply -f scripts/jaeger_config.yaml

echo "****** ISTIO INSTALLATION ******"
sh scripts/infra/istio_install.sh

sleep 1m

cd "$AETHER_DIR"
echo "****** AETHER 5GC INSTALLATION ******"
make aether-5gc-install
echo "****** AETHER MONITORING INSTALLATION ******"
make monitor-install
cd "$current_dir"
kubectl apply -f scripts/otel-service-monitor.yaml
kubectl apply -f scripts/tempo-service-monitor.yaml
cd "$AETHER_DIR"
make monitor-load
cd "$current_dir"

echo "****** REMOVING ISTIO FROM MET ******"
bash scripts/k8s/filter_istio_sidecar.sh