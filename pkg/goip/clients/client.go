package clients

// General interface implemented by all clients
type Client interface {
	Auth() error
    SetIp(ip string) error
    CheckEnv() error
}
