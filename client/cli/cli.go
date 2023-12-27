package cli

// 初始化Shell CLI接口
type Cli interface {
	// 初始化netstat方法
	Netstat(kwargs []string) (map[string]string, error)
	// 初始化process方法
	Process(kwargs []string) (map[string]string, error)
	// 初始化tty方法
	Tty(kwargs []string) (map[string]string, error)
}
