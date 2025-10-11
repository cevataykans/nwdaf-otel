kubectl delete -f scripts/istio.yaml
helm delete istiod -n istio-system
helm delete istio-base -n istio-system
kubectl delete namespace istio-system