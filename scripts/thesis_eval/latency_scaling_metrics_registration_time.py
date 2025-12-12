import matplotlib.pyplot as plt
import pandas as pd

def main():
    # Data provided
    data = {
        "": [1.716030040234375e+07, 1.353243395890411e+07, 1.331434862890625e+07, 1.3592443805882353e+07, 1.272334385265226e+07,],
    }

    df = pd.DataFrame(data)
    df['Iteration'] = [1, 2, 3, 4, 5]

    # Convert microseconds to seconds
    for col in [""]:
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
    cols = [""]

    for i, col in enumerate(cols):
        plt.plot(df["Iteration"], df[col], "-", marker=markers[i], linewidth=0.8, label=col)

    plt.xlabel("Iteration")
    plt.ylabel("Total Reg. Time (s)")
    plt.title("Registration Time per Iteration")
    plt.grid(True, linestyle='--', linewidth=0.4, alpha=0.7)
    # plt.legend(framealpha=1, fontsize=10)
    plt.xticks([1, 2, 3, 4, 5])
    plt.tight_layout()
    plt.savefig("latency_scaling_registration.png", dpi=300)

if __name__ == "__main__":
    main()