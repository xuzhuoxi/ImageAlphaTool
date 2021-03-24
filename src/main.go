package main

import (
	"fmt"
	"github.com/xuzhuoxi/ImageAlphaTool/src/lib"
	"github.com/xuzhuoxi/infra-go/encodingx/jsonx"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/logx"
	"io/ioutil"
	"os"
)

var (
	globalLogger logx.ILogger
	config       *lib.FlagConfig
)

func main() {
	globalLogger = logx.NewLogger()
	globalLogger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})

	cfg, err := lib.ParseFlag(globalLogger)
	if err != nil {
		globalLogger.Error(err)
		return
	}
	config = cfg
	invoke()
}

func invoke() {
	handleDelete := config.Result.CheckMode(lib.ResultDelete)
	handleLog := config.Result.CheckMode(lib.ResultLog)

	globalLogger.Traceln(fmt.Sprintf("开始处理图片：\"%s\"", config.Source.Folder))
	globalLogger.Traceln(fmt.Sprintf("设置如下："))
	globalLogger.Traceln(fmt.Sprintf("处理的图片格式=%v", config.Source.Formats))
	globalLogger.Traceln(fmt.Sprintf("处理的像素模式=%v，值为=%v", config.Filter.Mode, config.Filter.FilterValue))
	globalLogger.Traceln(fmt.Sprintf("结构执行模式=%d[删除(%t),记录(%t)]", config.Result.Mode, handleDelete, handleLog))
	if handleLog {
		globalLogger.Traceln(fmt.Sprintf("记录文件格式=(%s)", config.Result.LogFormat))
		globalLogger.Traceln(fmt.Sprintf("记录文件路径=(%s)", config.Result.LogPath))
	}
	globalLogger.Traceln(fmt.Sprintf("执行开始："))

	// 统计文件列表
	globalLogger.Traceln(fmt.Sprintf("统计文件......"))
	imageFiles, err := statFiles()
	if nil != err {
		globalLogger.Error(err)
		return
	}
	globalLogger.Traceln(fmt.Sprintf("统计文件完成，共(%d)个文件。", len(imageFiles)))

	// 检查文件
	globalLogger.Traceln(fmt.Sprintf("检查文件开始......"))
	resultFiles := checkImages(imageFiles, handleDelete)
	resultLen := len(resultFiles)
	globalLogger.Traceln(fmt.Sprintf("检查文件完成，共找到(%d)个文件符合要求。", resultLen))
	if handleDelete {
		globalLogger.Traceln(fmt.Sprintf("删除文件完成，共删除(%d)个文件。", resultLen))
	}
	// 导出结果数据
	if handleLog {
		globalLogger.Traceln(fmt.Sprintf("导出结果数据......"))
		exportResult(resultFiles)
		globalLogger.Traceln(fmt.Sprintf("导出结果数据完成，数据文件位于%s", config.Result.LogPath))
	}

	globalLogger.Traceln(fmt.Sprintf("执行结束！"))
}

// 统计文件，过滤非指定扩展名的文件
func statFiles() (imageFiles []string, err error) {
	err = filex.WaldAllFiles(config.Source.Folder, func(path string, info os.FileInfo, err error) error {
		fileName := info.Name()
		if !formatx.CheckFormatRegistered(filex.GetExtWithoutDot(fileName)) {
			return nil
		}
		_, _, extName := filex.SplitFileName(fileName)
		if !config.Source.CheckFormat(extName) {
			return nil
		}

		imageFiles = append(imageFiles, path)
		return nil
	})
	return imageFiles, err
}

// 检查全部图像的alpha通道
func checkImages(imageFiles []string, delete bool) (resultFiles []string) {
	for _, file := range imageFiles {
		ok, err := checkImage(file)
		if nil != err {
			globalLogger.Errorln(err)
			continue
		}
		if !ok {
			continue
		}
		resultFiles = append(resultFiles, file)
		globalLogger.Info("Delete Target Image At:", file)
		if delete {
			filex.Remove(file)
		}
	}
	return
}

// 检查图像的alpha通道
func checkImage(path string) (b bool, err error) {
	img, _, err := imagex.LoadImage(path, "")
	if nil != err {
		return false, err
	}
	size := img.Bounds().Size()
	for yIndex := 0; yIndex < size.Y; yIndex += 1 {
		for xIndex := 0; xIndex < size.X; xIndex += 1 {
			_, _, _, alpha := img.At(xIndex, yIndex).RGBA()
			if !config.Filter.CheckFilter(alpha) {
				return false, nil
			}
		}
	}
	return true, nil
}

// 导出结果数据
func exportResult(resultFiles []string) {
	var result = &lib.Result{Dir: config.Source.Folder, Data: resultFiles}
	if lib.ResultFileJson == config.Result.LogFormat {
		data := jsonx.NewJsonCodingHandlerSync().HandleEncode(result)
		ioutil.WriteFile(config.Result.LogPath, data, os.ModePerm)
		return
	}
	if lib.ResultFileYml == config.Result.LogFormat {
		return
	}
}
