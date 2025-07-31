kubectl patch deployment webui -n  aether-5gc -p '{"spec":{"template":{"metadata":{"annotations":{ "sidecar.istio.io/inject": "false"}}}}}'
kubectl patch deployment metricfunc -n aether-5gc -p '{"spec":{"template":{"metadata":{"annotations":{ "sidecar.istio.io/inject": "false"}}}}}'
kubectl patch deployment simapp -n aether-5gc -p '{"spec":{"template":{"metadata":{"annotations":{ "sidecar.istio.io/inject": "false"}}}}}'
kubectl patch deployment nwdaf-analytics-info -n aether-5gc -p '{"spec":{"template":{"metadata":{"annotations":{ "sidecar.istio.io/inject": "false"}}}}}'