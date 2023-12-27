package config

type VerifyType struct {
	TCP []string
}

func NewVerifyType() *VerifyType {
	return &VerifyType{
		TCP: []string{
			"ESTABLISHED",
			"SYN_SENT",
			"SYN_RECV",
			"FIN_WAIT1",
			"FIN_WAIT2",
			"TIME_WAIT",
			"CLOSE",
			"CLOSE_WAIT",
			"LAST_ACK",
			"LISTEN",
			"CLOSING",
		},
	}
}
