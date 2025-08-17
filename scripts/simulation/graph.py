import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
from pathlib import Path

# === Configuration ===
DB_FOLDER = "../data"
DB_FILE = "series.db"
TABLE = "series"
TIME_COLUMN = "ts"
VALUE_COLUMN = "cpu_usage"
OUTPUT_FILE = "cpu_vs_time.png"

def main():
    db_path = Path(DB_FOLDER) / DB_FILE

    # Connect to SQLite
    conn = sqlite3.connect(db_path)

    # Query data
    query = f"SELECT {TIME_COLUMN}, {VALUE_COLUMN} FROM {TABLE} WHERE service='amf' ORDER BY {TIME_COLUMN}"
    df = pd.read_sql_query(query, conn)
    conn.close()

    # Convert timestamp if needed
    df[TIME_COLUMN] = pd.to_datetime(df[TIME_COLUMN], errors="coerce")

    # Plot
    plt.figure(figsize=(12,6))
    plt.plot(df[TIME_COLUMN], df[VALUE_COLUMN], marker="o", linestyle="-")
    plt.xlabel("Time")
    plt.ylabel("CPU Usage")
    plt.title("AMF CPU Usage")
    plt.grid(True)
    plt.tight_layout()

    # Save to file
    plt.savefig(OUTPUT_FILE)
    print(f"✅ Saved plot to {OUTPUT_FILE}")

if __name__ == "__main__":
    main()
