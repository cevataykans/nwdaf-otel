import pandas as pd
import matplotlib
import matplotlib.pyplot as plt

# ============================================================
# Thesis Plot Settings (global defaults)
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
# Helper conversion functions
# ============================================================

def parse_latency(x):
    """Convert 'XX ms' or 'YY s' → float seconds. NaN → 0."""
    if isinstance(x, str):
        x = x.strip()
        if x.lower() == "nan":
            return 0.0
        if x.endswith("ms"):
            return float(x.replace("ms", "").strip()) / 1000.0
        if x.endswith("s"):
            return float(x.replace("s", "").strip())
    return 0.0

def load_latency_csv(path, columns_to_filter=None):
    df = pd.read_csv(path)

    # Filter only user-selected endpoints
    if columns_to_filter is not None:
        columns = [c for c in df.columns if c not in columns_to_filter]
        df = df[columns]

    # Convert timestamps
    df["Time"] = pd.to_datetime(df["Time"])
    t_start = df["Time"].iloc[0]
    df["Time"] = (
            (df["Time"] - t_start)
            .dt.total_seconds() / 60
    )

    # Parse latency units (ms/s -> seconds)
    for col in df.columns:
        if col == "Time":
            continue
        df[col] = df[col].apply(parse_latency)

    return df

# ============================================================
# Plotting function
# ============================================================

def plot_dual_latency(
        nf_name,
        server_csv,
        client_csv,
        columns_to_filter=None,
        vertical_lines=None,
        output="latency_dual.png"
):
    df_server = load_latency_csv(server_csv, columns_to_filter)
    df_client = load_latency_csv(client_csv, columns_to_filter)

    # --------------------------------------------------------
    # Create 1×2 subplot figure
    # --------------------------------------------------------
    fig, axes = plt.subplots(2, 1, figsize=(8, 10), sharey=True)

    # --------------------------------------------------------
    # Plot both panels
    # --------------------------------------------------------
    panels = [
        (f"{nf_name} Server Span Latencies", df_server, axes[0]),
        (f"{nf_name} Client Span Latencies", df_client, axes[1]),
    ]

    for title, df, ax in panels:
        ax.set_title(title)

        for col in df.columns:
            if col == "Time":
                continue
            ax.plot(df["Time"], df[col], linewidth=0.8, label=col)

        # Vertical zone boundaries (red dotted)
        if vertical_lines:
            for ts in vertical_lines:
                ax.axvline(pd.to_datetime(ts), color="red", linestyle="--", linewidth=0.7)

        # Smaller x labels + rotate for long endpoint names
        for tick in ax.get_xticklabels():
            tick.set_fontsize(8)

        # Grid (same as previous thesis style)
        ax.grid(True, linestyle='--', linewidth=0.4, alpha=0.7)
        ax.set_xlabel("Elapsed Time (min)")
        ax.legend(loc='upper left', fontsize=10)

    axes[0].set_ylabel("Endpoint p95 Latency (s)")
    axes[1].set_ylabel("Endpoint p95 Latency (s)")

    plt.tight_layout(rect=(0, 0, 1, 0.98))
    plt.savefig(output, dpi=300)

    print(f"Generated {output}")


# ============================================================
# Main
# ============================================================

def main():
    NF='UDM'
    folder='nov18'
    SERVER_CSV = f'grafana_data/{folder}/{NF}_SERVER.csv'
    CLIENT_CSV = f"grafana_data/{folder}/{NF}_CLIENT.csv"
    plot_dual_latency(
        nf_name=NF,
        server_csv=SERVER_CSV,
        client_csv=CLIENT_CSV,

        # Filter a subset of endpoints; None = use all
        columns_to_filter=[
            #PCF
            # 'POST /nnrf-nfm/v1/subscriptions',
            # 'PUT /nnrf-nfm/v1/nf-instances/{var}',
        ],

        # Vertical dotted lines marking simulation boundaries
        vertical_lines=[
            # "2025-11-18 22:37:00",
            # "2025-11-18 22:40:00"
        ],

        output=f"{NF.lower()}_p95_latencies_2.png"
    )


if __name__ == "__main__":
    main()