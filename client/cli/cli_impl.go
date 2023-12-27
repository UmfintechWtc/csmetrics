package cli

import (
	"collect-metrics/common"
	"fmt"
)

type CliImpl struct{}

func (c *CliImpl) Netstat(kwargs []string) (map[string]string, error) {
	cliK, err := common.ChangeInterfaceToCustomFormat(kwargs)
	if err != nil {
		return nil, err
	}
	cmd := fmt.Sprintf(`netstat -an | grep tcp | egrep -i %v | awk '{print $NF}' |sort | uniq -c`, cliK)
	fmt.Println(cmd)
	return nil, nil
}

func (c *CliImpl) Process(kwargs []string) (map[string]string, error) {
	return nil, nil
}

func (c *CliImpl) Tty(kwargs []string) (map[string]string, error) {
	return nil, nil
}

func NewCliImpl() Cli {
	return &CliImpl{}
}
