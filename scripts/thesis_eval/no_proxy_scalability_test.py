import matplotlib.pyplot as plt
import pandas as pd

def main():
    # Data provided
    data = {
        "1 UDM": [1.03958376e+07, 1.22742199e+07, 1.18276846e+07, 1.22896462e+07, 1.28115832e+07],
        "2 UDM": [1.02838588e+07, 1.01459327e+07, 1.12699414e+07, 9.90031053e+06, 1.17134242e+07],
        "4 UDM": [1.14432035e+07, 1.13605787e+07, 1.13603517e+07, 1.07611361e+07, 1.19594691e+07],
        "8 UDM": [1.14330645e+07, 1.05504399e+07, 1.08566307e+07, 1.15847962e+07, 1.08603787e+07]
    }

    df = pd.DataFrame(data)
    df['Iteration'] = [1, 2, 3, 4, 5]

    # Convert microseconds to seconds
    for col in ["1 UDM", "2 UDM", "4 UDM", "8 UDM"]:
        df[col] = df[col] / 1_000_000.0

    plt.rcParams.update({
        "font.family": "serif",
        "font.serif": ["Palatino", "DejaVu Serif"],
        "axes.labelsize": 11,
        "font.size": 11,
        "legend.fontsize": 10,
        "xtick.labelsize": 10,
        "ytick.labelsize": 10,
    })

    plt.figure(figsize=(7.5, 3.8))
    markers = ["o", "s", "^", "D"]
    cols = ["1 UDM", "2 UDM", "4 UDM", "8 UDM"]

    for i, col in enumerate(cols):
        plt.plot(df["Iteration"], df[col], "-", marker=markers[i], linewidth=0.8, label=col)

    plt.xlabel("Iteration")
    plt.ylabel("Total Reg Time (s)")
    plt.title("UDM Registration Time per Iteration")
    plt.grid(True, linestyle='--', linewidth=0.4, alpha=0.7)
    plt.legend(framealpha=1, fontsize=10)
    plt.xticks([1, 2, 3, 4, 5])
    plt.tight_layout()
    plt.savefig("no_proxy_scalability.png", dpi=300)

if __name__ == "__main__":
    main()