package lib

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strconv"
	"strings"
)

type FlagSource struct {
	// 来源文件夹
	Folder string
	// 输入文件格式
	Formats []string
}

func (fs *FlagSource) CheckFormat(format string) bool {
	format = strings.ToLower(format)
	if len(fs.Formats) == 0 {
		return false
	}
	for _, f := range fs.Formats {
		if f == format {
			return true
		}
	}
	return false
}

func (fs *FlagSource) ParseSource(source string, format string) error {
	BasePath := osxu.GetRunningDir()
	In := filex.FormatPath(source)
	if "" == In {
		return errors.New("FlagSource:in is empty! ")
	}
	if strings.Index(In, "./") == 0 {
		In = filex.Combine(BasePath, In)
	}
	if !filex.IsExist(In) {
		return errors.New(fmt.Sprintf("FlagSource:in(%s) is not exist! ", In))
	}
	if !filex.IsFolder(In) {
		return errors.New(fmt.Sprintf("FlagSource:in(%s) is not folder! ", In))
	}
	fs.Folder = In

	if format == "" {
		return errors.New("FlagSource:format is empty! ")
	}
	formats := strings.Split(strings.ToLower(format), ",")
	fs.Formats = formats
	return nil
}

type FlagFilter struct {
	// 过滤模式
	Mode FilterMode
	// 过滤参数
	FilterValue []uint32
}

func (ff *FlagFilter) Value() uint32 {
	return ff.FilterValue[0]
}

func (ff *FlagFilter) Min() uint32 {
	return ff.FilterValue[0]
}

func (ff *FlagFilter) Max() uint32 {
	return ff.FilterValue[1]
}

func (ff *FlagFilter) CheckFilter(value uint32) bool {
	switch ff.Mode {
	case FilterEqual:
		return ff.Value() == value
	case FilterBetween:
		return value >= ff.Value() && value < ff.Value()
	case FilterSmaller:
		return value < ff.Value()
	case FilterSmallerOrEqual:
		return value <= ff.Value()
	case FilterLarger:
		return value > ff.Value()
	case FilterLargerOrEqual:
		return value >= ff.Value()
	default:
		return false
	}
}

func (ff *FlagFilter) ParseFilter(mode FilterMode, value string) error {
	if mode == FilterNone {
		return errors.New("FlagFilter:mode is error! ")
	}
	ff.Mode = mode
	if mode == FilterBetween {
		values := strings.Split(value, ",")
		if len(values) != 2 {
			return errors.New("FlagFilter:value is error! ")
		}
		val0, err0 := strconv.Atoi(values[0])
		if nil != err0 {
			return err0
		}
		val1, err1 := strconv.Atoi(values[1])
		if nil != err1 {
			return err1
		}
		ff.FilterValue = []uint32{uint32(val0), uint32(val1)}
	} else {
		val, err := strconv.Atoi(value)
		if nil != err {
			return err
		}
		ff.FilterValue = []uint32{uint32(val)}
	}
	return nil
}

type FlagResult struct {
	// 结果模式
	Mode ResultMode
	// 记录格式，支持log、json、yaml
	LogFormat string
	// 记录文件路径
	LogPath string
}

func (fr *FlagResult) CheckMode(mode ResultMode) bool {
	return (fr.Mode & mode) > 0
}

func (fr *FlagResult) ParseResult(mode ResultMode, format string, path string, logger logx.ILogger) error {
	if mode == ResultNone {
		return errors.New("FlagResult:mode is error! ")
	}
	fr.Mode = mode

	if mode == ResultDelete {
		return nil
	}

	format = strings.ToLower(format)
	isSupportFormat := "log" == format || "json" == format || "yaml" == format
	if !isSupportFormat {
		return errors.New("ParseResult:format is error! ")
	}
	fr.LogFormat = format

	path = filex.FormatPath(path)
	if "" == path {
		return errors.New("ParseResult:path is empty! ")
	}
	BasePath := osxu.GetRunningDir()
	if strings.Index(path, "./") == 0 {
		path = filex.Combine(BasePath, path)
	}
	if filex.IsFolder(path) {
		return errors.New(fmt.Sprintf("ParseResult:path(%s) is folder! ", path))
	}
	fr.LogPath = path

	if format == "log" {
		dir, filename := filex.Split(path)
		logger.SetConfig(logx.LogConfig{Type: logx.TypeRollingFile, FileDir: dir, FileName: filename, FileExtName: ""})
	}
	return nil
}
