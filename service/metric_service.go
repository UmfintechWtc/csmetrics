package service

// 初始化全局Metrics接口
type MetricService interface {
	// 初始化TCP接口
	TcpMetric
	// 初始化Process接口
	ProcessMetric
	// 初始化Session接口
	SessionMetric
}

type TcpMetric interface {
	// 获取TCP的TIME_WAIT、ESTABLISHED
	TCP(name string, age int) error
}

type ProcessMetric interface {
	// 获取Process的总进程数、starcross用户启动的进程数
	Process(name string, age int) error
}

type SessionMetric interface {
	// 获取当前节点ssh session tty数量
	Session(name string, age int) error
}
