package main

import (
	"github.com/xuzhuoxi/ImageAlphaTool/src/lib"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/logx"
	"os"
)

var (
	globalLogger logx.ILogger
)

func main() {
	globalLogger = logx.NewLogger()
	globalLogger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})

	flagConfig, err := lib.ParseFlag(globalLogger)
	if err != nil {
		globalLogger.Error(err)
		return
	}
	invoke(flagConfig)
}

func invoke(cfg *lib.FlagConfig) {
	var list []string
	err := filex.WalkInDir(cfg.Source.Folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		if !formatx.CheckFormatRegistered(filex.GetExtWithoutDot(fileName)) {
			return nil
		}
		_, _, extName := filex.SplitFileName(fileName)
		if !cfg.Source.CheckFormat(extName) {
			return nil
		}

		list = append(list, path)
		return nil
	})
	if nil != err {
		globalLogger.Error(err)
		return
	}
	for _, file := range list {
		invokeImage(file, cfg)
	}
}

func invokeImage(path string, cfg *lib.FlagConfig) {
	img, _, err := imagex.LoadImage(path, "")
	if nil != err {
		return
	}
	size := img.Bounds().Size()
	//var colorModel := img.ColorModel()
	for yIndex := 0; yIndex < size.Y; yIndex += 1 {
		for xIndex := 0; xIndex < size.X; xIndex += 1 {
			_, _, _, alpha := img.At(xIndex, yIndex).RGBA()
			if !cfg.Filter.CheckFilter(alpha) {
				goto end
			}
		}
	}
	handleImage(path, cfg)
end:
	return
}

func handleImage(path string, cfg *lib.FlagConfig) {
	switch cfg.Result.Mode {
	case lib.Delete:
		filex.Remove(path)
		globalLogger.Info("Delete Image At:", path)
	}
}
