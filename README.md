# ImageClearer

ImageClearer can be used to clean up images in the catalog with pixels that match the criteria.

## Compatibility

go 1.15.8

## Getting Started

### Download Release

- Download the release [here](https://github.com/xuzhuoxi/ImageClearer/releases).

- Download the repository:

	```sh
	go get -u github.com/xuzhuoxi/ImageClearer
	```
	
	This will retrieve the library.

### Build

Execution the construction file([build.sh](/build/build.sh)) to get the releases if you have already downloaded the repository.

You can modify the construction file([build.sh](/build/build.sh)) to achieve what you want if necessary. The command line description is [here](https://github.com/laher/goxc).

## Run

### Command Line

Supportted command line parameters as follow:

| -       | -            | -                                                            |
| :------ | :----------- | ------------------------------------------------------------ |
| -src    | **required** | 来源目录路径 																									|
| -format | **required** | 来源文件格式(多个格式用","分开) 																	|
| -fm     | **required** | 配置模式(1区间2相等3小于4小于等于5大于6大于等于) 									|
| -fv     | **required** | 配置参数(fm=1时：数值,数值。其它：数值。注意数值：[0,255]) 				|
| -rm     | **required** | 结果模式(按位)	位1：记录 位2：直接删除 														|
| -rm2    | optional     | 记录格式(当rm=1时有效，支持log、json、yaml格式) 									|
| -rmv    | optional     | 文件路径(当rm=1时有效) 																				|

E.g.:

-src=./dst

-format=png,jpeg

-fm=2

-fv=0

-rm=3

-rm2=yml

-rmv=./delete.yml

## Related Library

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## Contact

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## License

ImageClearer source code is available under the MIT [License](/LICENSE).