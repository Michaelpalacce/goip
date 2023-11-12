package clients

// Client is a general interface implemented by all clients
type Client interface {
	Auth() error
    SetIp(ip string) error
    CheckEnv() error
}
