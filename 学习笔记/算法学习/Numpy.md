

@：矩阵乘法

拼接：

numpy.concatenate([a,b],axis=0(default))

torch.cat([a,b],axis=0(default))

等0是增行，等1是扩列



torch.from_numpy(a)有小数默认返回DoubleTensor类型

torch.Tensor(a)保持原类型

torch.tensor(a)默认DoubleTensor



numpy按照某一行或者某一列的排序方法

```python
data[data[:,0].argsort()] # 按照第一列进行排序
data[data[0,:].argsort()] # 按照第一行进行排序
# 为了记录排序之前的索引，我会在数据后面扩充一个
np.arange(1,len(data)).reshape(-1,1)
np.arange(1,len(data)).reshape(1,-1)
```



numpy种各种axis的作用

```
np.concatenate([a,b],axis=?)
# axis = 0 扩行，需要列通道相同
# axis = 1 扩列，需要行通道相同

np.sort(a,axis=?)
# axis = 0 按每一列中的元素排序
# axis = 1 按每一行中的元素排序

np.argmax(a,axis=?)
# 没有输入axis时会把数组展平，返回最大数值的索引
# axis = 0 返回每一列最大值在该列的索引
# axis = 1 返回每一行最大值在该行的索引
# np.argmin()同，只是返回最小值
```



切片

```
通过索引对列切片时
a[:,0] # 会返回展平后的第0列
a[:,0:1] #会返回未展平的第0列
```



掩膜的作用

```
mask = a == 1
a[mask] # 就能得到1
如果配上索引np.arange(1,len(a.flatten()))就可以获得对应展平后的索引
```



torch.tensor和torch.cuda转换

[LucBourrat1/RoomNet-Pytorch: Unofficial Pytorch Implementation of "RoomNet: End-to-End Room Layout Estimation (ICCV, 2017)" (github.com)](https://github.com/LucBourrat1/RoomNet-Pytorch)
