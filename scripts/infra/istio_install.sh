helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

# Aether k8s version supports up to this version, check compatibility here https://istio.io/latest/docs/releases/supported-releases/
helm install istio-base istio/base -n istio-system --set defaultRevision=default --create-namespace --version "1.28.0"
helm install istiod istio/istiod -n istio-system --version "1.28.0" --wait

kubectl create namespace aether-5gc
kubectl label namespace aether-5gc istio-injection=enabled

# Contains the tracer settings
kubectl apply -f scripts/istio.yaml