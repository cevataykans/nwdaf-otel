import yaml
import sys
import re

def edit_config(data, start_imsi, ue_count, cur_gnb, is_parallel):
    imsi_pattern = r'{{START_IMSI}}'
    ue_pattern = r'{{UE_COUNT}}'
    imsi_replaced = re.sub(imsi_pattern, rf'{start_imsi + 1}', data)
    ue_replaced = re.sub(ue_pattern, rf'{ue_count}', imsi_replaced)
    return ue_replaced

def get_config_name(prefix, index):
    return f'{prefix}-{index}.yaml'

def create_gnbsim_custom_configs(folder_path, template_path, prefix, start_imsi, gnb_count, ue_count, is_parallel):
    with open(template_path, 'r') as f:
        data = f.read()

    for i in range(gnb_count):
        new_config = edit_config(data, start_imsi, ue_count, i, is_parallel)

        config_name = get_config_name(prefix, i)
        with open(f'{folder_path}/{config_name}', 'w') as f:
            #print(new_config)
            f.write(new_config)

        start_imsi = start_imsi + 1 + ue_count # assume each gnb is also one imsi


def edit_gnbsim_main_config(main_config_path, gnb_count, prefix):
    with open(main_config_path, 'r') as f:
        data = yaml.safe_load(f)

    data['gnbsim']['docker']['container']['count'] = gnb_count
    servers = data['gnbsim']['servers'][0]
    servers.clear()

    config_path = 'deps/gnbsim/config'
    for i in range(gnb_count):
        servers.append(f'{config_path}/{get_config_name(prefix, i)}')

    data['gnbsim']['servers'][0] = servers

    # Write it back to the YAML file
    with open(main_config_path, 'w') as f:
        yaml.dump(data, f, default_flow_style=False, sort_keys=False)

def edit_gnbsim_config(aether_path, is_parallel, gnb_count, ue_count_per_gnb):
    imsi_start = 208930100007487
    gnbsim_main_config_path = f'{aether_path}/vars/main.yml'
    gnbsim_gnb_configs_folder = f'{aether_path}/deps/gnbsim/config'
    config_name_prefix = 'gnbsim-custom'
    gnbsim_template_config_path = f'scripts/simulation/{config_name_prefix}.yaml'

    #edit_gnbsim_main_config(gnbsim_main_config_path, gnb_count, config_name_prefix)
    create_gnbsim_custom_configs(gnbsim_gnb_configs_folder, gnbsim_template_config_path,
                                 config_name_prefix, imsi_start, gnb_count, ue_count_per_gnb, is_parallel)

# usage gnb_count, ue_count per gnb
if len(sys.argv) != 4:
    print('Usage: python3 gnbsim_configs.py <exec_parallel> <gnb_count> <ue_count_per_gnb>')
    sys.exit(1)

def parse_str(arg):
    return arg.lower() in ('yes', 'true', '1', 'True', 't')

parallel = False
try:
    parallel = parse_str(sys.argv[1])
except ValueError:
    print('Error: argument must be a bool: true, false')
    sys.exit(1)

gnb_count = -1
try:
    gnb_count = int(sys.argv[2])
except ValueError:
    print('Error: argument must be an integer')
    sys.exit(1)

ue_count = -1
try:
    ue_count = int(sys.argv[3])
except ValueError:
    print('Error: argument must be an integer')
    sys.exit(1)



print(f'Creating {gnb_count} with {ue_count} per gnb... parallel -> {parallel}')

#onramp_path = '../../Thesis/test/aether-onramp'
onramp_path = '../aether-onramp'
file_path = onramp_path
edit_gnbsim_config(file_path, parallel, gnb_count, ue_count)
