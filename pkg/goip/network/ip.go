package network

func GetPublicIp() ([]byte, error) {
	return GetBody("https://icanhazip.com")
}
