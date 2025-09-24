# nwdaf-otel 

As of now, only analytics-info is configured to serve requests.

## How to Install

* Log in to the node in which the project needs to be installed.
* Clone the repository.
* Edit paths under scripts/infra/install and scripts/infra/uninstall:
  * ISTIO_DIR="your istio installation directory"
  * AETHER_DIR="your aether installation directory"
* In the root of this project, run:
  * make install to setup aether cluster
  * make start to deploy an instance of this **NWDAF** project
  * make uninstall to erase the cluster
  * make stop to stop the instance of the **NWDAF** launched before

## Packages

* clients -> contains code related to communciation with other services: NRF, prometheus ...
* server -> contains server code of NWDAF
* cmd -> entry point for NWDAF
* scripts -> contains various scripts for benchmarks, infrastructure ...
* helm -> contains local helm chart used for deploying **NWDAF analytics info**
* repository -> contains code which allows NWDAF to have data persistency