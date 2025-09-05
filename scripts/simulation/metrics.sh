#!bin/bash

ue_values=(16 32 64 128 256 512 1024 1256)
repetition_count=5
gnbsim_wait_time=5
sleep_time=60
aether_dir="$HOME/aether-onramp/"
nwdaf_dir="$HOME/nwdaf-otel/"

# save starting time for graph reference
start_ts=$(date +%s)
for cur_ue_count in "${ue_values[@]}"; do
    cd $nwdaf_dir
    python3 scripts/simulation/gnbsim_configs.py "True" "1" "$cur_ue_count"
    echo "running gnbsim with $cur_ue_count UEs for $repetition_count times"
    for ((i=1; i<=repetition_count; i++)); do
        echo "iteration $i"
        # cd $aether_dir
        # make aether-gnbsim-run
        sleep $gnbsim_wait_time
    done
done
end_time=$(date +%s)

echo "Start time:  $start_time"
echo "Finish time: $end_time"
echo "Sleeping for nwdaf..."
sleep $sleep_time

# Stop NWDAF
cd $nwdaf_dir
make stop-analytics

# Plot graphs