#!bin/bash

#ue_values=(8 16 32 64 128 256 512)
aether_dir="$HOME/cores/aether-onramp-3-1-0/"
nwdaf_dir="$HOME/nwdaf-otel/"

for ((cur_ue_count=10; cur_ue_count<=510; cur_ue_count+=10)); do
    cd $nwdaf_dir
    python3 scripts/simulation/gnbsim_configs.py "True" "1" "$cur_ue_count"
    echo "running gnbsim with $cur_ue_count UEs"
    cd $aether_dir
    make aether-gnbsim-run
    sleep 45
done
end_ts=$(date +%s)
exit