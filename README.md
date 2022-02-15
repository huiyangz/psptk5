# psptk5

- 学习研究对象：PSP 太阁立志传 5(TAIKOU RISSHIDEN V) ULJM-05525|303A88BDAC0128BC|0001|G
- 学习目的：从 DATA.BIN 文件中提取所有的文件

## 使用方法

psptk5 extract [path of DATA.BIN] [path of EBOOT.BIN] [dir to extract]。
其中 EBOOT.BIN 要先解成 ELF，可以使用 deceboot.exe (v0.3)。

## 学习目标

- [x] 提取所有的文件
- [x] 解压其中 4 个 CMPS 文件
- [ ] 修改文件后，再压缩回 CMPS
- [ ] 修改提取后的文件后，再生成 DATA.BIN 和修改 EBOOT.BIN

## 有意思的信息

- 011B5000.cmps 文件是字库
- gmo 文件是游戏中单挑场景下，3D 模型和骨骼数据
- 游戏中界面文本信息在 EBOOT.BIN，其它还包括名胜、招式、称号、官职等

## 有困惑的地方

- 暂时没有找到每一个文件的名称
