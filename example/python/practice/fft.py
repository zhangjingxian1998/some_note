import numpy as np
import matplotlib.pyplot as plt

# 参数设置
sampling_rate = 1000  # 采样率 (Hz)
duration = 1.0        # 信号时长 (秒)
N = int(sampling_rate * duration)  # 总采样点数

# 生成时间轴 t
t = np.linspace(0, duration, N, endpoint=False)

# 生成两个不同频率的简谐波
freq1 = 50   # 频率1 (Hz)
amp1 = 0.7   # 幅度1
signal1 = amp1 * np.sin(2 * np.pi * freq1 * t)
# signal1 = np.array([np.exp(item) for item in np.arange(1000)])

freq2 = 120  # 频率2 (Hz)
amp2 = 1.0   # 幅度2
signal2 = amp2 * np.sin(2 * np.pi * freq2 * t)

# 叠加信号
combined_signal = signal1 + signal2

# 傅里叶变换
fft_result = np.fft.fft(combined_signal)
freq_axis = np.fft.fftfreq(N, 1/sampling_rate)

# 取频谱的绝对值（忽略相位）
amplitude_spectrum = np.abs(fft_result) / N  # 归一化

# 绘制结果
plt.figure(figsize=(12, 8))

# 绘制原始信号
plt.subplot(2, 2, 1)
plt.plot(t, signal1, label=f"{freq1} Hz (Amp={amp1})")
plt.xlim(0, 0.1)  # 仅显示前0.1秒的波形
plt.xlabel("Time (s)")
plt.ylabel("Amplitude")
plt.title("Signal 1")
plt.legend()

plt.subplot(2, 2, 2)
plt.plot(t, signal2, color="orange", label=f"{freq2} Hz (Amp={amp2})")
plt.xlim(0, 0.1)
plt.xlabel("Time (s)")
plt.title("Signal 2")
plt.legend()

# 绘制叠加后的信号
plt.subplot(2, 2, 3)
plt.plot(t, combined_signal, color="green")
plt.xlim(0, 0.1)
plt.xlabel("Time (s)")
plt.ylabel("Amplitude")
plt.title("Combined Signal")

# 绘制傅里叶变换结果（仅显示正频率部分）
plt.subplot(2, 2, 4)
positive_freq_mask = freq_axis >= 0
plt.plot(freq_axis[positive_freq_mask], amplitude_spectrum[positive_freq_mask], color="red")
plt.xlabel("Frequency (Hz)")
plt.ylabel("Amplitude")
plt.title("Frequency Spectrum")
plt.xlim(0, 200)  # 显示0-200 Hz范围

plt.tight_layout()
plt.show()