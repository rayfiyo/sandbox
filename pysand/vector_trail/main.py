import numpy as np
import matplotlib.pyplot as plt
from fractions import Fraction

# ===== 伝達関数リスト =====
transfer_functions = [
    ("G1", lambda s: 4 / (s + 1)),
    ("G2", lambda s: 4 / (s + 4)),
    ("G3", lambda s: 1 / (s**2 + 2 * s + 1)),
    ("G4", lambda s: 1 / (s**3 + 3 * s**2 + 3 * s + 1)),
]

# ωの範囲・マーク点
omega = np.linspace(0, 100, 1000)
omega_marks = [0, 0.5, 1, 5, 10, 50, 100]

for name, G in transfer_functions:
    # 軌跡データ
    s_vals = 1j * omega
    G_vals = G(s_vals)
    real_part = np.real(G_vals)
    imag_part = np.imag(G_vals)

    # プロット
    fig, ax = plt.subplots(figsize=(8, 8))
    ax.plot(real_part, imag_part, label=rf"${name}(j\omega)$", color="blue")

    # 代表点とラベル
    for w in omega_marks:
        Gw = G(1j * w)
        re, im = Gw.real, Gw.imag
        re_f = Fraction(re).limit_denominator(1000)
        im_f = Fraction(im).limit_denominator(1000)
        ax.plot(re, im, "ro")
        ax.text(
            re,
            im,
            f"ω={w}\nRe={re_f}\nIm={im_f}",
            fontsize=9,
            ha="left",
            va="bottom",
            color="darkred",
        )

    # 設定
    ax.set_xlabel("Re")
    ax.set_ylabel("Im")
    ax.set_title(f"Complex Plane: {name}")
    ax.grid(True)
    ax.axhline(0, color="gray", linewidth=0.5)
    ax.axvline(0, color="gray", linewidth=0.5)
    ax.legend()
    ax.set_aspect("equal", "box")
    fig.tight_layout()

    # ファイル保存＆クローズ
    fig.savefig(f"graph_{name}.png")
    plt.close(fig)
