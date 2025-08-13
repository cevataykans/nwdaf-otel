import time
import signal
import math
import subprocess
import os

# Flag to indicate whether the program should keep running
running = True
ueransim_executable_path = 'ueransim_onramp/build'
gnb_executable = 'nr-gnb'
ue_executable = 'nr-ue'
binder_executable = 'nr-binder'
max_device_count = 300
min_device_count = 20
cur_device_count = 0
wave_period_seconds = 600
max_device_spawn_in_second = 1
gnb_process: subprocess.Popen = None
ue_processes = []

def run_process(process_type, executable_path, args=None):
    if args is None:
        args = []
    global gnb_process, ue_processes
    process = subprocess.Popen(
        [executable_path] + args,
        stdout=subprocess.PIPE,
        stderr=subprocess.STDOUT,
        text=True,
        bufsize=1
    )
    if process_type == 'gnb':
        gnb_process = process
    else:
        ue_processes.append({'process': process, 'device_id': None})
    return process

def run_gnb():
    print('Spawning GNB Process\n')
    gnb = run_process('gnb',
                os.path.join('..', ueransim_executable_path, 'nr-gnb'),
                args=['-c', '../config/custom-gnb.yaml'])
    start = time.monotonic()
    # Process stdout
    for line in gnb.stdout:
        print(line, end='\n')
        line = line.strip()
        if 'NG Setup procedure is successful' in line:
            return True

        # Detect an error message
        if 'error' in line.lower():
            return False

        # Apply 10 second deadline
        if time.monotonic() - start > 10:
            return False
    return False

def run_ue(process):
    print('Starting UE')
    # # Process stdout
    # for line in process.stdout:
    #     line = line.strip()
    #
    #     # Extract an ID if present
    #     match = id_pattern.search(line)
    #     if match:
    #         process_id = match.group(1)
    #         print(f"Found ID: {process_id}")
    #
    #     # Detect an error message
    #     if "error" in line.lower():
    #         print(f"Error detected: {line}")
    #
    # # Optionally handle stderr separately
    # for err_line in process.stderr:
    #     print(f"[stderr] {err_line.strip()}")

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
    print('\nExit signal received. Shutting down...')
    running = False

def main_simulation():
    # Spawn GNB and check they are no errors
    gnb_ok = run_gnb()
    if not gnb_ok:
        print('GNB could not start, returning main sim ... \n')
        return

    global cur_device_count, running
    sleep_time = 1 / max_device_spawn_in_second
    while running:
        target_device_count = current_time_to_device_count()
        if cur_device_count < target_device_count:
            # spawn another UE
            print('TODO, SPAWN UE')
        time.sleep(sleep_time)
    print('Main Simulation Loop Exit\n')


def main():
    # Register signal handlers
    signal.signal(signal.SIGINT, signal_handler)   # Ctrl+C
    signal.signal(signal.SIGTERM, signal_handler)  # Termination signal

    print('Program started. Press Ctrl+C to exit.')
    main_simulation()

    # Cleanup resources, make every device disconnect, then stop gnb
    global gnb_process
    if gnb_process is not None and gnb_process.poll() is None:
        gnb_process.terminate()

if __name__ == '__main__':
    main()