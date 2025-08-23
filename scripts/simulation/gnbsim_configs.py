import yaml
import sys
import re

def quote_jinja_vars(yaml_text: str) -> str:
    """
    Wrap unquoted {{ ... }} placeholders in double quotes
    so YAML parsers won't choke, while keeping Ansible templating intact.
    """
    # Matches {{ ... }} that are not already inside quotes
    pattern = r'(?<!["\'])({{.*?}})(?!["\'])'
    return re.sub(pattern, r'"\1"', yaml_text)

def unquote_jinja_vars(yaml_text: str) -> str:
    """
    Remove quotes around {{ ... }} placeholders
    so Ansible can correctly resolve them.
    """
    # Match double or single quotes wrapping a {{ ... }} block
    pattern = r'(["\'])({{.*?}})\1'
    first_pass = re.sub(pattern, r'\2', yaml_text)
    fixed_text = re.sub(r"''(?!$)", r"'", first_pass, flags=re.MULTILINE)
    return fixed_text

def edit_config(data, start_imsi, ue_count, cur_gnb):
    pdu_test = data['configuration']['profiles'][1]
    pdu_test['ueCount'] = ue_count
    pdu_test['startImsi'] = f'{start_imsi + 1}' # assume this gnb takes one imsi
    data['configuration']['profiles'][1] = pdu_test

    base_gnb_id = '001001'
    gnb_id = int(base_gnb_id, 16)
    gnb_id += cur_gnb
    gnb_id_hex = f'{gnb_id:0x}'
    data['configuration']['gnbs']['gnb1']['globalRanId']['gNbId']['gNBValue'] = f'00{gnb_id_hex}'

def get_config_name(prefix, index):
    return f'{prefix}-{index}.yaml'

def create_gnbsim_custom_configs(folder_path, template_path, prefix, start_imsi, gnb_count, ue_count):
    with open(template_path, 'r') as f:
        data = yaml.safe_load(f)

    for i in range(gnb_count):
        edit_config(data, start_imsi, ue_count, i)

        config_name = get_config_name(prefix, i)
        with open(f'{folder_path}/{config_name}', 'w') as f:
            yaml.dump(data, f, default_flow_style=False, sort_keys=False)

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

def edit_gnbsim_config(aether_path, gnb_count, ue_count_per_gnb):
    imsi_start = 208930100007487
    gnbsim_main_config_path = f'{aether_path}/vars/main.yml'
    gnbsim_gnb_configs_folder = f'{aether_path}/deps/gnbsim/config'
    config_name_prefix = 'gnbsim-custom'
    gnbsim_template_config_path = f'scripts/simulation/{config_name_prefix}.yaml'

    edit_gnbsim_main_config(gnbsim_main_config_path, gnb_count, config_name_prefix)
    create_gnbsim_custom_configs(gnbsim_gnb_configs_folder, gnbsim_template_config_path,
                                 config_name_prefix, imsi_start, gnb_count, ue_count_per_gnb)

# usage gnb_count, ue_count per gnb
if len(sys.argv) != 3:
    print('Usage: python3 gnbsim_configs.py <gnb_count> <ue_count_per_gnb>')
    sys.exit(1)

gnb_count = -1
try:
    gnb_count = int(sys.argv[1])
except ValueError:
    print('Error: argument must be an integer')
    sys.exit(1)

ue_count = -1
try:
    ue_count = int(sys.argv[2])
except ValueError:
    print('Error: argument must be an integer')
    sys.exit(1)

print(f'Creating {gnb_count} with {ue_count} per gnb...')

#onramp_path = '../../Thesis/test/aether-onramp'
onramp_path = '../aether-onramp'
file_path = onramp_path
edit_gnbsim_config(file_path, gnb_count, ue_count)
