import random
import time
import signal
import math
import subprocess
import os

# Flag to indicate whether the program should keep running
running = True
ueransim_executable_path = 'ueransim_onramp/build'
ueransim_config_path = 'ueransim_onramp/config'
gnb_executable = 'nr-gnb'
ue_executable = 'nr-ue'
binder_executable = 'nr-binder'
max_device_count = 300
min_device_count = 20
cur_device_count = 0
wave_period_seconds = 600
max_device_spawn_in_second = 0.1
gnb_process: subprocess.Popen = None
ue_processes = []
starting_imsi = 208930100007487
available_imsis = []

def run_process(executable_path, args=None):
    if args is None:
        args = []
    process = subprocess.Popen(
        [executable_path] + args,
        stdout=subprocess.PIPE,
        stderr=subprocess.STDOUT,
        text=True,
        bufsize=1
    )
    return process

def run_gnb():
    print('\nSpawning GNB Process')
    gnb = run_process(
        os.path.join('..', ueransim_executable_path, 'nr-gnb'),
        args=['-c', os.path.join('..', ueransim_config_path, 'custom-gnb.yaml')]
    )
    global gnb_process
    gnb_process = gnb
    start = time.monotonic()
    # Process stdout
    for line in gnb.stdout:
        line = line.strip()
        if 'NG Setup procedure is successful' in line:
            return True

        # Detect an error message
        if 'error' in line.lower():
            print(line, end='\n')
            return False

        # Apply 10 second deadline
        if time.monotonic() - start > 10:
            return False
    return False

def run_ue():
    global available_imsis
    if len(available_imsis) > 0:
        cur_imsi = available_imsis.pop()
    else:
        global starting_imsi
        cur_imsi = starting_imsi
        starting_imsi += 1

    imsi_arg = f'imsi-{cur_imsi}'
    print(f'\nStarting UE with imsi: {imsi_arg}')
    ue = run_process(
        os.path.join('..', ueransim_executable_path, 'nr-ue'),
        args=['-c', os.path.join('..', ueransim_config_path, 'custom-ue.yaml'), '-i', imsi_arg]
    )
    # Check if successful registration
    tun_interface = ''
    start = time.monotonic()
    for line in ue.stdout:
        line = line.strip()

        # Detect an error message
        if "error" in line.lower():
            print(line)
            return False

        if 'is successful' in line and 'is up' in line:
            start = line.rfind('[')
            end = line.rfind(',')
            if start == -1 or end == -1:
                print(f'ERROR Device interface cannot be found in line: {line}\n')
                return False
            tun_interface = line[start+1:end]
            break

        # Apply 5 second deadline
        if time.monotonic() - start > 5:
            print('UE timed out for starting ... ')
            if ue.poll() is None:
                ue.terminate()
            return False

    print(f'UE tun interface: {tun_interface}\n')
    # Run ping command with deadline
    min_ping_duration = 1
    max_ping_duration = 60
    counter = random.randint(min_ping_duration, max_ping_duration)
    ping_process = run_process('ping', args=['-I', tun_interface, '8.8.8.8', '-c', f'{counter}'])

    # Save everything in list
    global ue_processes
    ue_metadata = {
        'process': ue,
        'ue_command': ping_process,
        'imsi': imsi_arg,
        'interface': tun_interface,
    }
    ue_processes.append(ue_metadata)
    return True

def spawn_ue():
    global cur_device_count
    target_device_count = current_time_to_device_count()
    success = False
    if cur_device_count < target_device_count:
        # spawn another UE
        success = run_ue()

    if success:
        cur_device_count += 1

def remove_unused_ue_resources():
    global ue_processes
    processes_to_remove = []
    for ue_metadata in ue_processes:
        cmd_process: subprocess.Popen = ue_metadata['ue_command']
        if cmd_process is not None and cmd_process.poll() is not None:
            for line in cmd_process.stdout:
                print(line)
            print(f'Will remove UE {ue_metadata["imsi"]}, {ue_metadata["interface"]}')
            # poll not none indicates command exit, can safely remove device
            processes_to_remove.append(ue_metadata)

    global available_imsis, cur_device_count
    cur_device_count -= len(processes_to_remove)
    for ue_to_remove in processes_to_remove:
        ue_processes.remove(ue_to_remove)

        # send deregister signal via nr-cli
        deregister = run_process(
            os.path.join('..', ueransim_executable_path, 'nr-cli'),
            args=[ue_to_remove['imsi'], '--exec', 'deregister normal']
        )
        deregister.wait()

        ue_process: subprocess.Popen = ue_to_remove['process']
        if ue_process is not None and ue_process.poll() is None:
            ue_process.terminate()

        imsi_number = int(ue_to_remove['imsi'][5:])
        available_imsis.append(imsi_number)

def kill_all_ue():
    global ue_processes
    for ue_metadata in ue_processes:
        cmd_process: subprocess.Popen = ue_metadata['ue_command']
        if cmd_process is not None and cmd_process.poll() is None:
            cmd_process.terminate()

        deregister = run_process(
            os.path.join('..', ueransim_executable_path, 'nr-cli'),
            args=[ue_metadata['imsi'], '--exec', 'deregister normal']
        )
        deregister.wait()

        ue_process: subprocess.Popen = ue_metadata['process']
        if ue_process is not None and ue_process.poll() is None:
            ue_process.terminate()

def current_time_to_device_count():
    global max_device_count, min_device_count, wave_period_seconds
    now = time.monotonic()
    # 0 → 2π
    phase = ((now % wave_period_seconds) / wave_period_seconds) * 2 * math.pi
    # output between [-1, 1]
    output = math.sin(phase)
    # map output to device count
    target_device_count = (output + 1) / 2 * (max_device_count - min_device_count) + min_device_count
    return target_device_count

def signal_handler(sig, frame):
    global running
    print('\nExit signal received. Shutting down...\n')
    running = False

def main_simulation():
    # Spawn GNB and check they are no errors
    gnb_ok = run_gnb()
    if not gnb_ok:
        print('GNB could not start, returning main sim ...\n')
        return

    global running
    sleep_time = 1 / max_device_spawn_in_second
    while running:
        spawn_ue()

        # Check out devices that have reached their lifespan and should deregister before exiting
        remove_unused_ue_resources()

        time.sleep(sleep_time)
    print('Main Simulation Loop Exit')


def main():
    # Register signal handlers
    signal.signal(signal.SIGINT, signal_handler)   # Ctrl+C
    signal.signal(signal.SIGTERM, signal_handler)  # Termination signal

    print('Program started. Press Ctrl+C to exit.')
    main_simulation()

    # Cleanup all devices first
    kill_all_ue()

    # stop gnb lastly
    global gnb_process
    if gnb_process is not None and gnb_process.poll() is None:
        gnb_process.terminate()

if __name__ == '__main__':
    main()