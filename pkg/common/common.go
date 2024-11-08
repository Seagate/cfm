package common

type ConnectionStatus string

const (
	ONLINE         ConnectionStatus = "online"
	FOUND          ConnectionStatus = "found"
	OFFLINE        ConnectionStatus = "offline"
	NOT_APPLICABLE ConnectionStatus = "n\\a"
)
