import re

def filter_ueid_lines(file_path,total_requested_ues):
    try:
        with open(file_path, 'r') as file:
            lines = file.readlines()

        # Find all occurrences of "ueId-start: " or "ueId-end: " followed by an integer
        pattern = re.compile(r"(ueId-(?:start|end): )\"(\d+)\"")
        max_value = None
        max_line_index = None
        max_match = None

        # Search for the max integer in matching lines
        for index, line in enumerate(lines):
            match = pattern.search(line)
            if match:
                #if "ueId-start: " in line or "ueId-end: " in line:
                number = int(match.group(2))
                if max_value is None or number > max_value:
                    max_value = number
                    max_line_index = index
                    max_match = match

        if max_value is not None:
            new_value = max_value + total_requested_ues

            # repalce old end value
            lines[max_line_index] = pattern.sub(fr'ueId-end: "{new_value}"', lines[max_line_index], 1)

            for i_s, line_s in enumerate(lines):
                if '- "'+str(max_value)+'"' in line_s:
                    for iter in range(total_requested_ues):
                        lines[i_s]+='                - "'+str(max_value + 1 + iter) +'"\n'

            with open(file_path, 'w') as file:
                file.writelines(lines)

            print(f"Updated largest ueId ({max_value}) to {new_value} in the file.")
        else:
            print("No matching 'ueId-start:' or 'ueId-end:' entries found.")


    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")


file_path = "../aether-onramp/deps/5gc/roles/core/templates/sdcore-5g-values.yaml"
total_requested_ues=500
filter_ueid_lines(file_path, total_requested_ues)
