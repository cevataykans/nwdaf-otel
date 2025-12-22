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
# CLIENT_CSV = f"grafana_data/nov18/{NF}_CLIENT.csv"
CLIENT_CSV = f"grafana_data/dec12/ScalingWithHeavyTraffic.csv"

# Endpoints to exclude from the plot (optional)
FILTER_OUT = [
    # AMF_CLIENT:
    'GET /nfconfig/access-mobility',
    'GET /nnrf-disc/v1/nf-instances',
    'PATCH /nnrf-nfm/v1/nf-instances/{var}',
    'POST /nnrf-nfm/v1/subscriptions',
    'POST /opentelemetry.proto.collector.trace.v1.TraceService/Export',
    'PUT /nnrf-nfm/v1/nf-instances/{var}',
    
    # AMF NATIVE TRACES LEAVE ONLY UDM
    "AMF NAS AuthenticationResponse",
    "AMF NAS RegistrationComplete",
    "AMF NAS RegistrationRequest",
    "AMF NAS SecurityModeComplete",
    "AMF NGAP InitialContextSetup",
    "AMF NGAP InitialUEMessage",
    "AMF NGAP NGSetup",
    "AMF NGAP UplinkNASTransport",
    "HTTP GET nrf/nf-instances",
    "HTTP POST nrf/subscriptions",
    "HTTP PUT nrf/nf-instances/{nfInstanceID}",
]

# Vertical timeline separators (timestamps as strings)
VERTICAL_LINES = [
    # Example:
    # "2025-11-18 22:37:00",
    # "2025-11-18 22:40:00"
    "2025-12-12 18:28:00",
    "2025-12-12 18:29:00",
    "2025-12-12 18:30:00",
    "2025-12-12 18:32:00",
    "2025-12-12 18:31:00",
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
    t_start = df["Time"].iloc[0]
    df["Time"] = (
            (df["Time"] - t_start)
            .dt.total_seconds() / 60
    )

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
        t_parsed = (t_parsed - t_start).total_seconds() / 60
        plt.axvline(t_parsed, color="red", linestyle="--", linewidth=0.8)


    plt.xlabel("Elapsed Time (min)")
    plt.ylabel("Endpoint p95 Latency (s)")
    plt.title(f'{NF} Client Span Latencies')

    plt.grid(True, linestyle='--', linewidth=0.4, alpha=0.7)
    # plt.legend(
    #     # loc='lower center',         # position relative to bounding box
    #     bbox_to_anchor=(0.5, -0.28), # x=center, y below the axes
    #     ncol=2,                   # optional: put entries on one row
    #     fontsize=8
    # )
    plt.legend(framealpha=0.5, fontsize=7.5, loc='upper right')
    plt.tight_layout()

    # plt.savefig(f"{NF.lower()}_endpoint_latencies.png", dpi=300)
    plt.savefig(f"latency_scaling_endpoint_latencies_2.png", dpi=300)
    # plt.savefig("udm_endpoint_latencies.pgf")


if __name__ == "__main__":
    main()