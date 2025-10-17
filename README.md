# nwdaf-otel 

As of now, only analytics-info is configured to serve requests.

## How to Configure Aether Core

* Download your core and cd into.
* Configure the settings according to the [official doc](https://docs.aetherproject.org/master/onramp/start.html)
* For our nuc nodes, you will most probably use these as they are:
```yaml
data_iface: enp86s0 # can double check with "ip a"
amf:
   ip: "IP of the NUC core is being installed on"
```
* I also change these IPs to have multiple core deployments not clashing with other people working simultaneously
```yaml
# Notice how my subnets for UPF differ from the default subnet, I recommend changing your own subnets to a range unique to your deployment
  upf:
    access_subnet: "192.168.202.1/24"	# access subnet & gateway
    core_subnet: "192.168.200.1/24"	# core subnet & gateway
    mode: af_packet			# Options: af_packet or dpdk
    multihop_gnb: false			# set to true to override default N3 subnet
    default_upf:
      ip:
        access: "192.168.202.3"
        core:   "192.168.200.3"
      ue_ip_pool: "172.250.0.0/16"

#...gnbsim
  router:
    data_iface: enp86s0
    macvlan:
      iface: gnbaccess
      subnet_prefix: "172.30"
```
* Edit settings under hosts.ini
* Open deps/5gc/roles/core/templates/sdcore-5g-values.yaml and add the following telemetry to amf settings while setting SBIs to **HTTP** for **each NF**
  * Optionally deploy sctplb but setting deploy to **true** depending on your use case.
  * Without SBI being set to HTTP, we cannot trace requests from ISTIO proxy.
    * We do not yet support HTTPS in ISTIO.
* Depending on your use case, increase, decrease the number of UEs the core supports.
```yaml
# Remember the configure every SBI for each NF
    amf:
      ngapp:
        externalIp: {{ core.amf.ip }}
      cfgFiles:
        amfcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http
          
    smf:
      cfgFiles:
        smfcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http

    pcf:
      cfgFiles:
        pcfcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http
              
    ausf:
      cfgFiles:
        ausfcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http

    nssf:
      cfgFiles:
        nssfcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http

    udr:
      cfgFiles:
        udrcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http

    udm:
      cfgFiles:
        udmcfg.yaml:
          configuration:
            nrfUri: http://nrf:29510
            sbi:
              scheme: http

    nrf:
      cfgFiles:
        nrfcfg.yaml:
          configuration:
            sbi:
              scheme: http

# Add the telemetry option
  amf:
	  cfgFiles:
		  configuration:
			  telemetry:
				  enabled: true
				  otlp_endpoint: "simplest-collector.default.svc.cluster.local:4317"
				  ratio: 0.4      # use 1.0 for debugging                          # Optional; defaults to 1.0. If set to 0, AMF assumes 1.0.
...
```
* Open vars/main.yaml and configure monitoring version (notice the fix typo "-" to "_" in moniroing-crd -> monitoring_crd)
```yaml
...
  monitor:
    helm:
      chart_ref: rancher/rancher-monitoring
      chart_version: 104.1.4+up57.0.3

  monitor_crd:
    helm:
      chart_ref: rancher/rancher-monitoring-crd
      chart_version: 104.1.4+up57.0.3
```
* Copy scripts/data/grafana-observability.json to deps/amp/roles/monitor-load/templates/5g-monitoring/observability.json
* Copy scripts/data/nf-metrics.json to deps/amp/roles/monitor-load/templates/5g-monitoring/nf-metrics.json
* Open deps/amp/roles/monitor-load/templates/5g-monitoring/kustomization.yaml and add the copied dashboard:
```yaml
resources:
  - ./metricfunc-monitor.yaml
  - ./upf-monitor.yaml

configMapGenerator:
  - name: grafana-ops-dashboards
    namespace: cattle-dashboards
    files:
      - ./5g-dashboard.json
      - ./observability.json
      - ./nf-metrics.json
generatorOptions:
  labels:
    grafana_dashboard: "1"
```
* Open deps/amp/roles/monitor/templates/monitor-values.yaml and add the following:
```yaml
rancherMonitoring:
  enabled: false

grafana:
  additionalDataSources:
    - name: Tempo
      type: tempo
      uid: df14gae7gchkwf
      url: http://tempo.default.svc.cluster.local:3200
      basicAuth: false
      editable: true
      jsonData:
        serviceMap:
          datasourceUid: 'prometheus'
        nodeGraph:
          enabled: true
        search:
          hide: false
  service:
    type: NodePort
```

* Core is ready to be deployed!

## How to Install

* Log in to the node in which the project needs to be installed.
* Clone the repository.
* Edit paths under scripts/infra/install and scripts/infra/uninstall:
  * ISTIO_DIR="your istio installation directory"
  * AETHER_DIR="your aether installation directory"
* In the root of this project, run:
  * make install to setup aether cluster
  * make uninstall to erase the cluster
  * make start to deploy an instance of this **NWDAF** project
  * make stop to stop the instance of the **NWDAF** launched before

## How to Build

* Changes made to code trigger CI, which then builds up a Docker image.
* Local helm chart always deploys the latest image.

## Packages

* clients -> contains code related to communication with other services: NRF, prometheus ...
* server -> contains server code of NWDAF
* cmd -> entry point for NWDAF
* scripts -> contains various scripts for benchmarks, infrastructure ...
* helm -> contains local helm chart used for deploying **NWDAF analytics info**
* repository -> contains code which allows NWDAF to have data persistency