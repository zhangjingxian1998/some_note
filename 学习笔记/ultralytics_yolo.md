# 1、yolo11训练后10代损失快速下降
模型训练默认在最后10代默认开启close_mosaic，数据变得简单，因此更易学习。

![](img\close_mosaic.png)

# 2、yolo11的多尺度输入
多尺度输入是由multi_scale超参控制，默认是False，即只使用原图输入；

如果设置为True，则会在原图输入的基础上，增加不同尺度的输入。

每一批次的输入尺寸是一样的，不同批次的输入尺寸之间会有变化。
![](img\multi_scale.png)

# 3、yolo anchor-based
yolov5是使用的anchor-base方法，使用kmeans聚类。<br>
yolov8使用的是anchor-free方法。<br>
yolov5的方法在autoanchor.py文件中实现<br>
yolov8的方法在utils/tal.py文件中实现TaskAlignedAssigner类

yolov5的autoanchor按照宽高来进行聚类，anchor数量是指定的9个。<br>
因为聚类时使用的是欧氏距离,为了防止一个方向上数值占据主导,计算时要对宽高两个方向的数据除以标准差。<br>
得到n个anchor后，对训练数据进行anchor分配，数据和anchor的长宽方向上，任意一个方向的覆盖率达到25%以上，则为该数据分配该锚框。如果有多个锚框满足条件，则选择覆盖率最高的锚框分配。计算出一个fitness值，fitness值越大，则该锚框的效果越好。<br>
再使用遗传算法对锚框进行优化，选取fitness最高的锚框组合。

yolov5会输出3种尺度的特征图，锚框平均分配给这三个尺度的特征图。模型会为特征图上每个点会为每个锚框预测一个置信度和偏移量。
<center>
    <img style="border-radius: 0.3125em;
    box-shadow: 0 2px 4px 0 rgba(34,36,38,.12),0 2px 10px 0 rgba(34,36,38,.08);" 
    src="img\散点数据.png">
    <br>
    <div style="color:orange; border-bottom: 1px solid #d9d9d9;
    display: inline-block;
    color: #999;
    padding: 2px;">未除以标准差数据聚类结果</div>
</center>

<center>
    <img style="border-radius: 0.3125em;
    box-shadow: 0 2px 4px 0 rgba(34,36,38,.12),0 2px 10px 0 rgba(34,36,38,.08);" 
    src="img\散点数据除以std.png">
    <br>
    <div style="color:orange; border-bottom: 1px solid #d9d9d9;
    display: inline-block;
    color: #999;
    padding: 2px;">除以标准差数据聚类结果</div>
</center>

# 4、yolo anchor-free
Anchor-Based的局限性<br>
①Anchor的设置需要手动去设计(长宽比，尺度大小，anchor的数量)，对不同数据集也需要不同的设计，相当麻烦。<br>
②Anchor的匹配机制使得极端尺度(特别大/小的object)被匹配到的频率，相对于大小适中的object被匹配到的频率更低，网络在学习时不容易学习这些极端样本。<br>
③Anchor的庞大数量使得存在严重的不平衡问题，涉及到采样、聚类的过程。但聚类的表达能力在复杂情况下是有限的<br>
④Anchor-Based为了兼顾多尺度下的预测能力，推理得到的预测框也相对较多，在输出处理时的nms计算也会更加耗时。


anchor-free的方法不需要从训练数据中聚类anchor

yolov11 没有使用anchor-base方法，而是特征图上每个特征点预测1个目标框相四个方向的偏移值和置信度。<br>
偏移值范围是[0,16]，通过特征图上的stride来转换为真实框大小。<br>
同样是3种尺度的特征图，在三种尺度特征图上创建mesh矩阵，然后计算每个特征点对应的框，通过特征图上对应的stride来转换为真实框大小。训练时与gt进行IoU计算，进行框的筛选。

在进行损失值计算之前，需要有一些预处理阶段<br>
1、通过预测的概率分布得到每个anchor_point上的预测框<br>
2、根据gt框，选择被真实框包围的anchor_point<br>
3、计算预测框和gt框的IoU<br>
4、根据每个预测框的置信度和IoU计算一个综合评价值<br>
5、为每个gt框保存10个topk最高综合值的预测框<br>

6、如果一个anchor_point对应多个gt框，则保留IoU最大的那个anchor_point与gt对应
7、为符合要求的anchor分配真实框
