#!/bin/bash

current_dir=$(pwd)

# Paths to the directories for the corresponding application (Path ending in the directory)
# Manually adapt to local setup (TODO: Clone istio and Aether if necessary BUT: Values still need to be adapted manually)
ISTIO_DIR=/home/sevinc/jungmann/istio
AETHER_DIR=/home/sevinc/aether-onramp/

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
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml
sleep 1m

echo "****** OTEL INSTALLATION ******"
#kubectl apply -f otel_operator_go_dec.yaml
kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
sleep 3m

echo "****** FILTERED ELASTIC INSTALLATION ******"
kubectl apply -f scripts/collector_filtered_elastic.yaml   # collector_filtered.yaml
echo "****** JAEGER CONFIG INSTALLATION ******"
kubectl apply -f scripts/jaeger_config.yaml

cd "$ISTIO_DIR"
echo "****** ISTIO INSTALLATION ******"
sh install_istio.sh
cd "$current_dir"

sleep 1m

cd "$AETHER_DIR"
echo "****** AETHER 5GC INSTALLATION ******"
make aether-5gc-install
echo "****** REMOVING ISTIO FROM MET ******"
bash scripts/k8s/remove_istio_from_met_nf.sh
echo "****** AETHER AMP INSTALLATION ******"
make aether-amp-install
echo "****** AETHER UERANSIM INSTALLATION ******"
#make aether-ueransim-install
cd "$current_dir"

echo "****** REMOVING ISTIO FROM MET ******"
bash scripts/k8s/remove_istio_from_met_nf.sh
#echo "****** PORT FORWARDING JAEGER ******"
#bash scripts/k8s/port-forward.sh &
#bash ./injection.sh