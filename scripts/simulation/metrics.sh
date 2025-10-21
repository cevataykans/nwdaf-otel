#!bin/bash

ue_values=(8)#16 32 64 128 256 512)
repetition_count=1
gnbsim_wait_time=5
wait_nwdaf=1
aether_dir="$HOME/cores/aether-onramp-3-1-0/"
nwdaf_dir="$HOME/nwdaf-otel/"

#cd $nwdaf_dir
#make start-analytics
#sleep $wait_nwdaf

# save starting time for graph reference
start_ts=$(date +%s)
echo "Start time:  $start_ts"
for cur_ue_count in "${ue_values[@]}"; do
    cd $nwdaf_dir
    python3 scripts/simulation/gnbsim_configs.py "True" "1" "$cur_ue_count"
    echo "running gnbsim with $cur_ue_count UEs"
    cd $aether_dir
    for ((i=1; i<=repetition_count; i++)); do
#        echo "iteration $i"
        make aether-gnbsim-run
        sleep $gnbsim_wait_time
    done
done
end_ts=$(date +%s)

echo "Start time:  $start_ts"
echo "Finish time: $end_ts"
#echo "Sleeping for nwdaf..."
#sleep $wait_nwdaf

# Stop NWDAF
#cd $nwdaf_dir
#make stop-analytics
#sleep $gnbsim_wait_time

# Plot graphs
#python3 scripts/data/graph.py "$start_ts" "$end_ts"
exit