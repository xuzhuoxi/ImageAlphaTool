package lib

type FilterMode int

const (
	// 无
	FilterNone FilterMode = iota
	// 区间
	FilterBetween
	// 等于
	FilterEqual
	// 小于
	FilterSmaller
	// 小于等于
	FilterSmallerOrEqual
	// 大于
	FilterLarger
	// 小于等于
	FilterLargerOrEqual
)

type ResultMode int

const (
	// 记录
	ResultLog ResultMode = 1 << iota
	// 删除
	ResultDelete
	// 无
	ResultNone ResultMode = 0
)

const (
	ResultFileLog  = "log"
	ResultFileJson = "json"
	ResultFileYml  = "yml"
)
