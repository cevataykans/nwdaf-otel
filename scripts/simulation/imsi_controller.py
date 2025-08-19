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

def edit_imsi_lines(core_values_path, total_requested_ues):
    initial_device_count = 14
    imsi_start = 208930100007487

    with open(core_values_path, 'r') as f:
        text = f.read()

    fixed_text = quote_jinja_vars(text)

    with open(core_values_path, 'w') as f:
        f.write(fixed_text)

    with open(core_values_path, 'r') as f:
        data = yaml.safe_load(f)

    config = data['omec-sub-provision']['config']['simapp']['cfgFiles']['simapp.yaml']['configuration']
    subscribers = config['subscribers']
    subscribers.clear()

    subscribers.append({
        "ueId-start": f"{imsi_start}",
        "ueId-end": f"{imsi_start + initial_device_count - 1}",
        "plmnId": "20893",
        "opc": "981d464c7c52eb6e5036234984ad0bcf",
        "op": "",
        "key": "5122250214c33e723a5dd523fc145fc0",
        "sequenceNumber": "16f3b3f70fc2"
    })
    cur_device_count = initial_device_count
    if cur_device_count < total_requested_ues:
        new_imsi_start = imsi_start + cur_device_count
        remaining_count = total_requested_ues - cur_device_count
        subscribers.append({
            "ueId-start": f"{new_imsi_start}",
            "ueId-end": f"{new_imsi_start + remaining_count}",
            "plmnId": "20893",
            "opc": "981d464c7c52eb6e5036234984ad0bcf",
            "op": "",
            "key": "5122250214c33e723a5dd523fc145fc0",
            "sequenceNumber": "16f3b3f70fc2"
        })
    config['subscribers'] = subscribers

    available_imsis = []
    for i in range(0, total_requested_ues):
        available_imsis.append(f'{imsi_start + i}')

    config['device-groups'][0]['imsis'] = available_imsis

    # Write it back to the YAML file
    with open(core_values_path, 'w') as f:
        yaml.dump(data, f, default_flow_style=False, sort_keys=False)

    with open(core_values_path, 'r') as f:
        text = f.read()

    reverted_text = unquote_jinja_vars(text)

    with open(core_values_path, 'w') as f:
        f.write(reverted_text)


if len(sys.argv) != 2:
    print('Usage: python3 imsi_controller.py <maximum_ue_count>')
    sys.exit(1)

ue_count = -1
try:
    ue_count = int(sys.argv[1])
except ValueError:
    print('Error: argument must be an integer')
    sys.exit(1)

print(f'Creating imsi for devices up to: {ue_count}')

file_path = '../aether-onramp/deps/5gc/roles/core/templates/sdcore-5g-values.yaml'
edit_imsi_lines(file_path, ue_count)
