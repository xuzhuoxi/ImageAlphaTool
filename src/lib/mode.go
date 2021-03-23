package lib

type FilterMode int

const (
	// 无
	FilterNone FilterMode = iota
	// 区间
	Between
	// 等于
	Equal
	// 小于
	Smaller
	// 小于等于
	SmallEqual
	// 大于
	Larger
	// 小于等于
	LargeEqual
)

type ResultMode int

const (
	// 无
	ResultNone ResultMode = iota
	// 记录
	Log
	// 删除
	Delete
)
