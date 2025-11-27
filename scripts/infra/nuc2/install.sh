#!/bin/bash

current_dir=$(pwd)

AETHER_DIR=/home/sevinc/cores/aether-onramp-3-1-0
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

echo "****** ISTIO INSTALLATION ******"
sh scripts/infra/istio_install.sh
sleep 1m

echo "****** KEDA INSTALLATION ******"
helm repo add kedacore https://kedacore.github.io/charts
helm repo update
helm install keda kedacore/keda --version 2.10.2 --namespace keda --create-namespace
sleep 1m

echo "****** AETHER 5GC INSTALLATION ******"
cd "$AETHER_DIR"
make aether-5gc-install
cd "$current_dir"
make start-nwdaf
echo "****** AETHER AMP INSTALLATION ******"
cd "$AETHER_DIR"
make monitor-install
cd "$current_dir"
kubectl apply -f scripts/otel-service-monitor.yaml
kubectl apply -f scripts/tempo-service-monitor.yaml
cd "$AETHER_DIR"
make monitor-load
echo "****** AETHER GNBSIM INSTALLATION ******"
make aether-gnbsim-install
cd "$current_dir"
kubectl apply -f scripts/udm_scaled_object.yaml

echo "****** REMOVING ISTIO FROM MET ******"
bash scripts/k8s/filter_istio_sidecar.sh