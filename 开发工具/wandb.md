# wandb是什么及作用
wandb是Weight & Bias的缩写，这是一个与Tensorboard类似的参数可视化平台。不过，相比较TensorBoard而言，Wandb更加的强大，主要体现在以下的几个方面：
### 复现模型：
Wandb更有利于复现模型。这是因为Wandb不仅记录指标，还会记录超参数和代码版本。
### 自动上传云端：
如果你把项目交给同事或者要去度假，Wandb可以让你便捷地查看你制作的所有模型，你就不必花费大量时间来重新运行旧实验。
### 快速、灵活的集成：
只需5分钟即可把Wandb加到自己的项目。下载Wandb免费的开源Python包，然后在代码中插入几行，以后你每次运行模型都会得到记录完备的指标和记录。
### 集中式指示板：
Wandb提供同样的集中式指示板。不管在哪里训练模型，不管是在本地机器、实验室集群还是在云端实例；
这样就不必花时间从别的机器上复制TensorBoard文件。
### 强大的表格：
对不同模型的结果进行搜索、筛选、分类和分组。可以轻而易举地查看成千上万个模型版本，并找到不同任务的最佳模型。
# wandb安装与注册登陆
1、在终端中安装wandb库
```
pip install wandb
```
2、在官网注册账号获取API

https://wandb.ai/site # 注册后不可更改的用户名(username)会被用于与wandb连接,谨慎取名

3、在终端与wandb账号绑定
```
wandb login # 将API粘贴上,实现绑定
```

4、测试代码
```
import wandb
 
config = dict (
  learning_rate = 0.01,
  momentum = 0.2,
  architecture = "CNN",
  dataset_id = "peds-0192",
  infra = "AWS",
)
 
wandb.init(
  project="detect-pedestrians",
  notes="tweak baseline",
  tags=["baseline", "paper1"],
  config=config,
)
```
打开运行结果里有火箭那一行的链接，浏览器中出现如下结果就说明成功了。

只需要
import wandb
```
wandb.init(project="your project name",
            name='test1')
```
上传离线文件
wandb sync ./run-20230316_161158-m7886rzk