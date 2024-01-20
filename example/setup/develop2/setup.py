from setuptools import setup, find_packages
packages = find_packages()
print(packages)
__version__ = '1.0' # 版本号
# requirements = open('requirements.txt').readlines() # 依赖文件
setup(
    name = 'ts', # 在pip中显示的项目名称
    version = __version__,
    author = 'Dechin',
    author_email = 'dechin.phy@gmail.com',
    url = '',
    description = 'ts: Test Setup',
    packages = find_packages(exclude=['tests']), # 项目中需要拷贝到指定路径的文件夹
    # package_data={
    #     'my_project': ['ts/*'],
    # },
    python_requires = '>=3.5.0',
    # install_requires = requirements # 安装依赖
        )