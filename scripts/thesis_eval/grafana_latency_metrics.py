import pandas as pd
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from datetime import datetime

# ============================================================
# Thesis Plot Settings (your global defaults)
# ============================================================
matplotlib.use("pgf")
matplotlib.rcParams.update({
    "pgf.texsystem": "pdflatex",
    "text.usetex": True,
    "pgf.rcfonts": False,
    "font.family": "serif",
    "font.serif": ["Palatino"],
    "axes.labelsize": 11,
    "font.size": 11,
    "legend.fontsize": 10,
    "xtick.labelsize": 10,
    "ytick.labelsize": 10,
})

# ============================================================
# User configuration
# ============================================================

NF='AMF'
SERVER_CSV = f'grafana_data/nov18/{NF}_SERVER.csv'
CLIENT_CSV = f"grafana_data/nov18/{NF}_CLIENT.csv"

# Endpoints to exclude from the plot (optional)
FILTER_OUT = [
    # AMF_CLIENT:
    'GET /nfconfig/access-mobility',
    'GET /nnrf-disc/v1/nf-instances',
    'PATCH /nnrf-nfm/v1/nf-instances/{var}',
    'POST /nnrf-nfm/v1/subscriptions',
    'POST /opentelemetry.proto.collector.trace.v1.TraceService/Export',
    'PUT /nnrf-nfm/v1/nf-instances/{var}'
]

# Vertical timeline separators (timestamps as strings)
VERTICAL_LINES = [
    # Example:
    # "2025-11-18 22:37:00",
    # "2025-11-18 22:40:00"
]

# ============================================================
# Helpers
# ============================================================

def convert_latency(val):
    """Convert 'XX ms' or 'YY s' → float seconds. NaN → 0."""
    if isinstance(val, str):
        val = val.strip()

        # milliseconds → convert to seconds
        if val.endswith("ms") or val.endswith(" ms"):
            ms = float(val.replace("ms", "").strip())
            return ms / 1000.0

        # seconds → already seconds
        if val.endswith("s") or val.endswith(" s"):
            seconds = float(val.replace("s", "").strip())
            return seconds

        # numeric-like (assume seconds)
        try:
            return float(val)
        except:
            return 0.0

    return 0.0 if pd.isna(val) else float(val)

# ============================================================
# Main
# ============================================================

def main():
    # Load CSV
    df = pd.read_csv(CLIENT_CSV)
    df["Time"] = pd.to_datetime(df["Time"])

    # Identify endpoint columns
    endpoint_cols = [c for c in df.columns if c != "Time"]
    endpoints = [c for c in endpoint_cols if c not in FILTER_OUT]

    # Convert latency fields
    for col in endpoints:
        df[col] = df[col].apply(convert_latency)

    # ============================================================
    # Plotting
    # ============================================================

    plt.figure(figsize=(7.5, 3.8))

    markers = ["o", "s", "^", "D", "v", "p", "X", "*"]
    marker_index = 0
    lines = []

    for col in endpoints:
        marker_index += 1
        plt.plot(df["Time"], df[col], "-", linewidth=0.8, label=col)

    for ts in VERTICAL_LINES:
        t_parsed = pd.to_datetime(ts)
        plt.axvline(t_parsed, color="red", linestyle="--", linewidth=0.8)


    plt.xlabel("Time")
    plt.ylabel("Endpoint p95 Latency (s)")
    plt.title(f'{NF} Client Span Latencies')

    plt.grid(True, linestyle='--', linewidth=0.4, alpha=0.7)
    plt.legend(framealpha=0.5, fontsize=9)
    plt.tight_layout()

    plt.savefig(f"{NF.lower()}_endpoint_latencies.png", dpi=300)
    # plt.savefig("udm_endpoint_latencies.pgf")


if __name__ == "__main__":
    main()