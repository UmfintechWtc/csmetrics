package service

import (
	"fmt"
)

type MetricImpl struct{}

// 实例化TcpMetric
func (i *MetricImpl) TCP(name string, age int) error {
	fmt.Println(fmt.Sprintf("%v === %v", name, age))
	return nil
}

// 实例化SessionMetric
func (m *MetricImpl) Session(name string, age int) error {
	fmt.Println(fmt.Sprintf("%v === %v", name, age))
	return nil
}

// 实例化ProcessMetric
func (m *MetricImpl) Process(name string, age int) error {
	fmt.Println(fmt.Sprintf("%v === %v", name, age))
	return nil

}

func NewMetricImpl() MetricService {
	return &MetricImpl{}
}
