import torch
import torch.nn as nn
import matplotlib.pyplot as plt

import seaborn as sns
conv_weights = {}
conv_biases = {}
plt.figure(figsize=(10, 6))
plt.xlim(-0.5, 0.5)
# 递归遍历所有子模块
def extract_conv_params(module, prefix=''):
    for name, child in module.named_children():
        full_name = f"{prefix}.{name}" if prefix else name
        if isinstance(child, nn.Conv2d):
            # 提取权重（weight）和偏置（bias）
            conv_weights[full_name] = child.weight.detach().numpy().copy()
            if child.bias is not None:
                conv_biases[full_name] = child.bias.detach().numpy().copy()
        # 递归遍历子模块
        extract_conv_params(child, full_name)

weight = torch.load('一些想法\\yolo11n.pt', map_location='cpu')
for i in weight['model'].model._modules:
    extract_conv_params(weight['model'].model._modules[i])

for name, conv_weight in conv_weights.items():
    conv_weight.flatten()
    sns.kdeplot(conv_weight.flatten(),label = name)
plt.savefig('weight_distribution.png')
print(weight.keys())