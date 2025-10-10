#!/bin/bash

current_dir=$(pwd)

# Paths to the directories for the corresponding application (Path ending in the directory)
# Manually adapt to local setup (TODO: Clone istio and Aether if necessary BUT: Values still need to be adapted manually)
ISTIO_DIR=/home/sevinc/jungmann/istio
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

# TODO add if machine should be integrated in the cluster
# copy kubeconfig to cmvm7
#scp -i ~/.ssh/cmvm7_key ~/.kube/config jungmann@cmvm7.cit.tum.de:~/.kube/config
#ssh -i ~/.ssh/cmvm7_key  jungmann@cmvm7.cit.tum.de "sudo rm -rf /opt/local-path-provisioner/; sudo mkdir /opt/local-path-provisioner/"

# Set the necessary traffic forwarding rules in the machine running the ue simulation. Otherwise, UPF cannot forward data
#ssh -i ~/.ssh/cmvm7_key  jungmann@cmvm7.cit.tum.de "sudo ip route add 192.168.252.0/24 via 131.159.25.123 dev ens192; sudo ip route add 192.168.250.0/24 via 131.159.25.123 dev ens192"
#ssh -i ~/.ssh/cmvm7_key  jungmann@cmvm8.cit.tum.de "sudo ip route add 192.168.252.0/24 via 131.159.25.123 dev ens192; sudo ip route add 192.168.250.0/24 via 131.159.25.123 dev ens192"

echo "****** CERT MANAGER INSTALLATION ******"
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.18.2/cert-manager.yaml
sleep 1m

echo "****** OTEL INSTALLATION ******"
#kubectl apply -f otel_operator_go_dec.yaml
kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/download/v0.136.0/opentelemetry-operator.yaml
sleep 5m

echo "****** FILTERED ELASTIC INSTALLATION ******"
kubectl apply -f scripts/collector_filtered_elastic.yaml   # collector_filtered.yaml
echo "****** TEMPOOOOOOO ******"
kubectl apply -f scripts/tempo.yaml
#echo "****** JAEGER CONFIG INSTALLATION ******"
#kubectl apply -f scripts/jaeger_config.yaml

cd "$ISTIO_DIR"
echo "****** ISTIO INSTALLATION ******"
sh install_istio.sh
cd "$current_dir"

sleep 1m

cd "$AETHER_DIR"
echo "****** AETHER 5GC INSTALLATION ******"
make aether-5gc-install
echo "****** AETHER AMP INSTALLATION ******"
make monitor-install
cd "$current_dir"
kubectl apply -f scripts/otel-service-monitor.yaml
kubectl apply -f scripts/tempo-service-monitor.yaml
cd "$AETHER_DIR"
make monitor-load
echo "****** AETHER GNBSIM INSTALLATION ******"
make aether-gnbsim-install
cd "$current_dir"

echo "****** REMOVING ISTIO FROM MET ******"
bash scripts/k8s/filter_istio_sidecar.sh
#echo "****** PORT FORWARDING JAEGER ******"
#bash scripts/k8s/port-forward_jaeger.sh &
#bash ./injection.sh