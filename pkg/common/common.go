package common

type ConnectionStatus string

const (
	ONLINE         ConnectionStatus = "online"
	OFFLINE        ConnectionStatus = "offline"
	NOT_APPLICABLE ConnectionStatus = "n\\a"
)
