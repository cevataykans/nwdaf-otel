#!/bin/bash

current_dir=$(pwd)

# Paths to the directories for the corresponding application (Path ending in the directory)
# Manually adapt to local setup (TODO: Clone istio and Aether if necessary BUT: Values still need to be adapted manually)
ISTIO_DIR="~/istio"
AETHER_DIR="~/aether-onramp/"

# remove dangling data from disk from older metric installations
sudo rm -rf /opt/local-path-provisioner/
sudo mkdir /opt/local-path-provisioner/

cd "$AETHER_DIR"
make aether-k8s-install
cd "$current_dir"

echo "Sleeping for 30s"
sleep 30s
# copy kubeconfig to cmvm7
# TODO: ideally integrate into kubernetes setup task in ansible
# scp -i ~/.ssh/cmvm7_key ~/.kube/config jungmann@cmvm7.cit.tum.de:~/.kube/config

# Set the necessary traffic forwarding rules in the machine running the ue simulation. Otherwise, UPF cannot forward data
# ssh -i ~/.ssh/cmvm7_key  jungmann@cmvm7.cit.tum.de "sudo ip route add 192.168.252.0/24 via 131.159.25.123 dev ens192; sudo ip route add 192.168.250.0/24 via 131.159.25.123 dev ens192"

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml
sleep 1m

kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
#kubectl apply -f otel_operator_go_dec.yaml  # Annotation to support otel Autoinstrumentation
sleep 2m

kubectl apply -f collector.yaml
kubectl apply -f jaeger_config.yaml
#kubectl apply -f my_instrumentation.yaml # Only needed for otel autoinstrumentation (specifies collector as the endpoint)

cd "$ISTIO_DIR"
sh install_istio.sh
cd "$current_dir"

sleep 1m

cd "$AETHER_DIR"
make aether-5gc-install
make aether-amp-install
#make aether-ueransim-install
cd "$current_dir"

# Jaeger UI port is forwarded - runs in the background
kubectl port-forward svc/jaeger 16686 --address 0.0.0.0 -n default &