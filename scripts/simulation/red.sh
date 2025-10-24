#!bin/bash

ue_value=8
repetition_count=1
gnbsim_wait_time=30
aether_dir="$HOME/cores/aether-onramp-3-1-0/"
nwdaf_dir="$HOME/nwdaf-otel/"

#TODO: create folder with UE count, save every graph under there
current_day=$(date '+%d-%m-%Y')
archive_folder="${current_day}-ue-${ue_value}-$(date +%s)"
mkdir ${archive_folder}

start_ts=$(date +%s)
echo "Start time:  $start_ts"
cd $nwdaf_dir
python3 scripts/simulation/gnbsim_configs.py "True" "1" "$ue_value"

cd $aether_dir
for ((i=1; i<=repetition_count; i++)); do
    make aether-gnbsim-run
    sleep $gnbsim_wait_time
done
end_ts=$(date +%s)

kubectl port-forward service/rancher-monitoring-prometheus -n cattle-monitoring-system 9090:9090 &
prom_pf_process_id=$!
# Function that can query Prometheus
archive() {
  local query_name=$1
  local archive_name=$2
  local folder=$3
  curl -G "http://localhost:9090/api/v1/query_range" \
    --data-urlencode "query=${query_name}" \
    --data-urlencode "start=$(date -d '10 minutes ago' +%s)" \
    --data-urlencode "end=$(date +%s)" \
    --data-urlencode "step=30" > "${folder}/${archive_name}.json"
}

query="histogram_quantile(0.95, sum by(le, span_name) (rate(traces_spanmetrics_latency_bucket{service_name=~\"amf.aether-5gc\"}[30s])))"
name='latency_bucket'
archive $query $name $archive_folder

kill $prom_pf_process_id
echo "Start time:  $start_ts"
echo "Finish time: $end_ts"