import pandas as pd
import seaborn as sns
import matplotlib
import matplotlib.pyplot as plt

# Example data
simulations = [
    {
        'filename': 'amf_heatmap.png',
        'title': 'p95 Latencies per Protocol per AMF Instance (256 UE)',
        'data': {
            # SecurityModeComplete, NAS AuthenticationResponse, NAS Registration Complete, NAS Registration Request
            "1": [(9.75 + 9.75 + 9.75) / 3, (1.97 + 1.96 + 1.38) / 3, (0.057 + 0.544 + 0.051) / 3, (1.39 + 1.32 + 0.985) / 3],
            "2": [(9.75 + 9.75 + 9.66) / 3, (4.84 + 1.97 + 1.96) / 3, (0.091 + 0.076 + 0.087) / 3, (2.05 + 1.65 + 1.37) / 3],
            "3": [(9.75 + 9.75 + 9.75) / 3, (4.83 + 1.96 + 4.20) / 3, (0.075 + 0.089 + 0.068) / 3, (1.84 + 1.31 + 1.90) / 3],
            "4": [(9.75 + 9.75 + 9.75), (1.97 + 1.97 + 1.98), (0.062 + 0.090 + 0.066), (1.91 + 1.89 + 1.83)],
        },
        'endpoints': ["NAS SecurityModeComplete", "NAS AuthenticationResponse", "NAS RegistrationComplete", "NAS RegistrationRequest"]
    },
    {
        'filename': 'udm_heatmap.png',
        'title': 'p95 Latencies per Endpoint per UDM Instance (256 UE)',
        'data': {
            "1": [23, 18, 35],
            "2": [30, 25, 40],
            "3": [22, 18, 40],
            "4": [28, 22, 38],
        },
        'endpoints': ["EP1", "EP2", "EP3"]
    },
    {
        'filename': 'ausf_heatmap.png',
        'title': 'p95 Latencies per Endpoint per AUSF Instance (256 UE)',
        'data': {
            "1": [23, 18, 35],
            "2": [30, 25, 40],
            "3": [22, 18, 40],
            "4": [28, 22, 38],
        },
        'endpoints': ["EP1", "EP2", "EP3"]
    },
    {
        'filename': 'lb_heatmap.png',
        'title': 'p95 Latencies per Endpoint per AMF Instance with DRSM LB (256 UE)',
        'data': {
            "1": [23, 18, 35],
            "2": [30, 25, 40],
            "3": [23, 18, 35],
            "4": [30, 25, 40],
        },
        'endpoints': ["EP1", "EP2", "EP3"]
    },
]


def main():
    for sim in simulations:
        df = pd.DataFrame(sim['data'], index=sim['endpoints'])

        matplotlib.use('pgf')
        matplotlib.rcParams.update({
            'pgf.texsystem': 'pdflatex',
            'text.usetex': True,
            'pgf.rcfonts': False,
            'font.family': 'serif',
            'font.serif': ['Palatino'],
            'axes.labelsize': 11,
            'font.size': 11,
            'legend.fontsize': 10,
            'xtick.labelsize': 10,
            'ytick.labelsize': 10,
        })

        # Create heatmap
        plt.figure(figsize=(8, 6))
        sns.heatmap(df, annot=True, fmt=".2f", cmap="YlOrRd", cbar_kws={'label': 'p95 Latency (s)'})
        plt.title(sim['title'])
        plt.ylabel("Endpoints")
        plt.xlabel("Instance Count")
        plt.tight_layout()
        plt.savefig(sim['filename'])  # Saves it as PDF for LaTeX


if __name__ == '__main__':
    main()