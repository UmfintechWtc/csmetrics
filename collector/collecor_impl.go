package collector

import "sync"

type CollectorValuesImpl struct {
	CVMap sync.Map
}

// NewCollectorValuesImpl 用于构造 CollectorValuesImpl 实例
func NewCollectorValuesImpl() CollectorValues {
	return &CollectorValuesImpl{}
}
