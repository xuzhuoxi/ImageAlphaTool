package lib

import (
	"flag"
	"github.com/xuzhuoxi/infra-go/logx"
)

type FlagConfig struct {
	Source *FlagSource
	Filter *FlagFilter
	Result *FlagResult
}

func (fc *FlagConfig) ParseSource(source string, format string) error {
	fc.Source = &FlagSource{}
	return fc.Source.ParseSource(source, format)
}

func (fc *FlagConfig) ParseFilter(mode FilterMode, value string) error {
	fc.Filter = &FlagFilter{}
	return fc.Filter.ParseFilter(mode, value)
}

func (fc *FlagConfig) ParseResult(mode ResultMode, format string, path string, logger logx.ILogger) error {
	fc.Result = &FlagResult{}
	return fc.Result.ParseResult(mode, format, path, logger)
}

// -src 		(必要)来源文件夹
// -format 		(必要)来源文件格式				多个格式用","分开
//
// -fm 			(必要)配置模式					1区间2相等3小于4小于等于5大于6大于等于
// -fv 			(必要)配置参数					fm=1时：数值,数值。其它：数值。注意数值：[0,255]
//
// -rm			(必要)结果模式					1记录2直接删除
// -rm2			(非必要)记录格式					当rm=1时有效，支持log、json、yaml格式
// -rmv			(非必要)文件路径					当rm=1时有效
func ParseFlag(globalLogger logx.ILogger) (cfg *FlagConfig, err error) {
	src := flag.String("src", "", "(必要)来源文件夹")
	format := flag.String("format", "", "(必要)来源文件格式，多个格式用','分开!")
	fm := flag.Int("fm", 0, "(必要)配置模式，1区间2相等3小于4小于等于5大于6大于等于!")
	fv := flag.String("fv", "", "(必要)配置参数，fm=1时：数值,数值。其它：数值。注意数值：[0,255]!")
	rm := flag.Int("rm", 0, "(必要)结果模式，1记录2直接删除!")
	rm2 := flag.String("rm2", "", "(非必要)记录格式，当rm=1时有效，支持log、json、yaml格式!")
	rmv := flag.String("rmv", "", "(非必要)文件路径，当rm=1时有效!")
	flag.Parse()

	cfg = &FlagConfig{}
	err = cfg.ParseSource(*src, *format)
	if nil != err {
		return nil, err
	}
	err = cfg.ParseFilter(FilterMode(*fm), *fv)
	if nil != err {
		return nil, err
	}
	err = cfg.ParseResult(ResultMode(*rm), *rm2, *rmv, globalLogger)
	if nil != err {
		return nil, err
	}

	return
}
