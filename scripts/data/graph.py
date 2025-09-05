import sqlite3
import time

import pandas as pd
import matplotlib.pyplot as plt
from pathlib import Path
import sys

# === Configuration ===
DB_FOLDER = "../data"
DB_FILE = "series.db"
TABLE = "series"
TIME_COLUMN = "ts"
VALUE_COLUMN = "cpu_usage"
OUTPUT_FOLDER = 'graphs'
OUTPUT_FILE = "cpu_vs_time.png"

def parse_args():
    # usage gnb_count, ue_count per gnb
    if len(sys.argv) != 3:
        print('Usage: python3 graph.py <start_ts_utc_unix> <end_ts_utc_unix>')
        sys.exit(1)

    start_ts = -1
    try:
        start_ts = int(sys.argv[1])
    except ValueError:
        print('Error: start_ts_utc_unix must be an integer')
        sys.exit(1)

    end_ts = -1
    try:
        end_ts = int(sys.argv[2])
    except ValueError:
        print('Error: end_ts_utc_unix must be an integer')
        sys.exit(1)

    return start_ts, end_ts

def main():

    start_ts, end_ts = parse_args()
    db_path = Path(DB_FOLDER) / DB_FILE

    # Connect to SQLite
    conn = sqlite3.connect(db_path)
    services = [
        'bessd',
        'amf',
        'ausf',
        'nrf',
        'nssf',
        'pcf',
        'smf',
        'udm',
        'udr',
    ]

    # column name -> title of plot
    columns = {
        'cpu_usage': 'CPU Total Seconds',
        'memory_usage': 'Memory',
        'total_bytes_sent': 'Network Bytes Sent',
        'total_bytes_received': 'Network Bytes Received',
        'total_packets_sent': 'Network Packets Sent',
        'total_packets_received': 'Network Packets Received',
        'avg_trace_duration': 'Average Trace Duration'
    }

    cur_time = int(time.time())
    folder_path = Path.home() / OUTPUT_FOLDER / str(cur_time)
    folder_path.mkdir(parents=True, exist_ok=True)
    for service in services:
        service_folder_path = folder_path / service
        service_folder_path.mkdir(exist_ok=True)

        query = f"SELECT * FROM {TABLE} WHERE service='{service}' AND ts BETWEEN {start_ts} AND {end_ts} ORDER BY {TIME_COLUMN}"
        df = pd.read_sql_query(query, conn)
        conn.close()

        # Convert timestamp if needed
        df[TIME_COLUMN] = pd.to_datetime(df[TIME_COLUMN], unit="s", errors="coerce")
        # draw each graph at the folder
        for column, title in columns.items():
            plt.figure(figsize=(12,6))
            plt.plot(df[TIME_COLUMN], df[column], marker="o", linestyle="-")
            plt.xlabel("Time")
            plt.ylabel(column)
            plt.title(title)
            plt.grid(True)
            plt.tight_layout()

            graph_path = service_folder_path / column
            plt.savefig(graph_path)
            print(f"✅ Saved plot to {graph_path}")

if __name__ == "__main__":
    main()
