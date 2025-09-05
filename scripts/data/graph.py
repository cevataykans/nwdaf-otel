import os
import sqlite3
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
OUTPUT_FOLDER = '../graphs'
OUTPUT_FILE = "cpu_vs_time.png"

def parse_args():
    # usage gnb_count, ue_count per gnb
    if len(sys.argv) != 4:
        print('Usage: python3 graph.py <service_name> <start_ts_utc_unix> <end_ts_utc_unix>')
        sys.exit(1)

    service_name = sys.argv[1]

    start_ts = -1
    try:
        start_ts = int(sys.argv[2])
    except ValueError:
        print('Error: start_ts_utc_unix must be an integer')
        sys.exit(1)

    end_ts = -1
    try:
        end_ts = int(sys.argv[3])
    except ValueError:
        print('Error: end_ts_utc_unix must be an integer')
        sys.exit(1)
    return service_name, start_ts, end_ts

def main():

    service_name, start_ts, end_ts = parse_args()
    db_path = Path(DB_FOLDER) / DB_FILE

    # Connect to SQLite
    conn = sqlite3.connect(db_path)

    # Query data
    query = f"SELECT {TIME_COLUMN}, {VALUE_COLUMN} FROM {TABLE} WHERE service='amf' AND ts BETWEEN {start_ts} AND {end_ts} ORDER BY {TIME_COLUMN}"
    df = pd.read_sql_query(query, conn)
    conn.close()

    # Convert timestamp if needed
    df[TIME_COLUMN] = pd.to_datetime(df[TIME_COLUMN], unit="s", errors="coerce")

    # Plot
    plt.figure(figsize=(12,6))
    plt.plot(df[TIME_COLUMN], df[VALUE_COLUMN], marker="o", linestyle="-")
    plt.xlabel("Time")
    plt.ylabel("CPU Usage")
    plt.title("AMF CPU Usage")
    plt.grid(True)
    plt.tight_layout()

    folder_path = Path(OUTPUT_FOLDER)
    folder_path.mkdir(exist_ok=True)
    # Save to file
    joined_path = folder_path / OUTPUT_FILE
    plt.savefig(joined_path)
    print(f"✅ Saved plot to {joined_path}")


if __name__ == "__main__":
    main()
