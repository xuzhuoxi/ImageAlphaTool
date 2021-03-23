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
	// 无
	ResultNone ResultMode = 1 << iota
	// 记录
	ResultLog
	// 删除
	ResultDelete
)
