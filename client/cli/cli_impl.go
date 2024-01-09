package cli

type CliImpl struct{}

func (c *CliImpl) GaugeValues(cmd string) (*GaugeValues, error) {
	n, err := executeCommandWithFloat64(cmd)
	if err != nil {
		return nil, err
	}
	gauge.CmdRes = n
	return gauge, nil

}

func NewCliImpl() ShellCli {
	return &CliImpl{}
}
